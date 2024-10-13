package main

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/eynopv/image-scaler/internal/validator"
)

func (app *application) writeJson(
	w http.ResponseWriter,
	status int,
	data map[string]any,
) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonData)

	return nil
}

func (app *application) serveFile(
	w http.ResponseWriter,
	r *http.Request,
	fileName string,
) {
	filePath := app.storage.FilePath(fileName)
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileInfo.Name()))
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", fileInfo.Size()))

	http.ServeFile(w, r, filePath)
}

func (app *application) readImage(r *http.Request) (multipart.File, error) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		return nil, err
	}

	image, _, err := r.FormFile("image")
	if err != nil {
		return nil, err
	}

	return image, nil
}

func (app *application) readInt(
	query url.Values,
	key string,
	defaultValue int,
	validator *validator.Validator,
) int {
	value := query.Get(key)

	if value == "" {
		return defaultValue
	}

	parsedValue, err := strconv.Atoi(value)
	if err != nil {
		validator.Message = err.Error()
		return defaultValue
	}

	return parsedValue
}
