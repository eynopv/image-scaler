package data

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"mime/multipart"
	"net/http"

	"github.com/eynopv/image-scaler/internal/validator"
)

func ValidateImageUploadFormat(file multipart.File, v *validator.Validator) bool {
	buffer := make([]byte, 512)

	_, err := file.Read(buffer)
	if err != nil {
		v.Message = err.Error()
	}

	file.Seek(0, 0)

	contentType := http.DetectContentType(buffer)
	if contentType != "image/jpeg" && contentType != "image/png" {
		v.Message = "only jpeg and png formats supported"
	}

	return true
}

func ImageToBuffer(img image.Image, format string) (*bytes.Buffer, error) {
	if img == nil {
		return nil, fmt.Errorf("empty image")
	}

	var buf bytes.Buffer

	var err error
	switch format {
	case "jpeg", "jpg":
		err = jpeg.Encode(&buf, img, nil)
	case "png":
		err = png.Encode(&buf, img)
	default:
		return nil, fmt.Errorf("unsupported format")
	}

	if err != nil {
		return nil, err
	}

	return &buf, nil
}
