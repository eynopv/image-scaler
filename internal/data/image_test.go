package data

import (
	"image"
	"image/color"
	"testing"
)

func TestImageToBuffer(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 1, 1))
	img.Set(0, 0, color.RGBA{255, 0, 0, 255})

	testCases := []struct {
		name         string
		format       string
		data         image.Image
		expectError  bool
		expectBuffer bool
	}{
		{"png", "png", img, false, true},
		{"jpg", "jpeg", img, false, true},
		{"unsupported", "gif", img, true, false},
		{"nil", "jpeg", nil, true, false},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			buf, err := ImageToBuffer(testCase.data, testCase.format)

			if !testCase.expectError && err != nil {
				t.Errorf("expected nil got %v; err", err)
			}

			if testCase.expectError && err == nil {
				t.Errorf("expected non empty; err")
			}

			if testCase.expectBuffer && (buf == nil || buf.Len() == 0) {
				t.Errorf("expected non empty; buf")
			}
		})
	}
}
