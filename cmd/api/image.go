package main

import (
	"fmt"
	"image"
	"net/http"

	"github.com/eynopv/image-scaler/internal/data"
	"github.com/eynopv/image-scaler/internal/validator"
	"github.com/eynopv/image-scaler/pkg/transforms"
	"github.com/julienschmidt/httprouter"
)

func (app *application) uploadImageHandler(w http.ResponseWriter, r *http.Request) {
	image, err := app.readImage(r)
	if err != nil {
		app.badRequestErrorResponse(w, r, err)
		return
	}
	defer image.Close()

	v := validator.NewValidator()
	data.ValidateImageUploadFormat(image, v)

	if !v.IsValid() {
		app.badRequestErrorResponse(w, r, fmt.Errorf(v.Message))
		return
	}

	imageKey, err := app.storage.Save("", image)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	response := map[string]any{
		"imageUrl": imageKey,
	}

	err = app.writeJson(w, http.StatusOK, response)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) downloadImageHandler(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	originalFileName := params.ByName("id")
	fileName := originalFileName

	v := validator.NewValidator()

	width := app.readInt(r.URL.Query(), "width", 0, v)
	height := app.readInt(r.URL.Query(), "height", 0, v)

	v.Check(width >= 0, "expected positive width")
	v.Check(height >= 0, "expected positive height")

	if !v.IsValid() {
		app.badRequestErrorResponse(w, r, fmt.Errorf(v.Message))
		return
	}

	requestedSize := data.Size{
		Width:  width,
		Height: height,
	}

	if requestedSize.IsNonNull() {
		fileName = fmt.Sprintf("%s_%d_%d", fileName, requestedSize.Width, requestedSize.Height)
	}

	file, _ := app.storage.Get(fileName)
	if file != nil {
		app.serveFile(w, r, fileName)
		return
	}

	if fileName == originalFileName {
		app.notFoundErrorResponse(w, r)
		return
	}

	file, err := app.storage.Get(originalFileName)
	if err != nil {
		app.notFoundErrorResponse(w, r)
		return
	}
	defer file.Close()

	img, format, err := image.Decode(file)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	scaledSize := transforms.CalculateScaledSize(
		data.Size{Width: img.Bounds().Dx(), Height: img.Bounds().Dy()},
		requestedSize,
	)
	scaledImage := transforms.ScaleImage(img, scaledSize)

	buffer, err := data.ImageToBuffer(scaledImage, format)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	scaledImageKey, err := app.storage.Save(fileName, buffer)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.serveFile(w, r, scaledImageKey)
}
