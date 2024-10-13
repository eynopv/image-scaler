package transforms

import (
	"image"
	"image/color"
	"image/draw"
	"testing"

	"github.com/eynopv/image-scaler/internal/data"
)

func TestCalculateScaledSize(t *testing.T) {
	portraitSize := data.Size{Width: 500, Height: 1000}
	landscapeSize := data.Size{Width: 1000, Height: 500}
	squareSize := data.Size{Width: 1000, Height: 1000}

	testCases := []struct {
		name     string
		src      data.Size
		newMax   data.Size
		expected data.Size
	}{
		{
			"portrait",
			portraitSize,
			data.Size{Width: 100, Height: 100},
			data.Size{Width: 50, Height: 100},
		},
		{
			"landspace",
			landscapeSize,
			data.Size{Width: 100, Height: 100},
			data.Size{Width: 100, Height: 50},
		},
		{
			"square",
			squareSize,
			data.Size{Width: 100, Height: 100},
			data.Size{Width: 100, Height: 100},
		},
		{
			"upscale",
			squareSize,
			data.Size{Width: 1000, Height: 1000},
			data.Size{Width: 1000, Height: 1000},
		},
		{
			"landscape scale based only on width",
			landscapeSize,
			data.Size{Width: 100, Height: 0},
			data.Size{Width: 100, Height: 50},
		},
		{
			"landspace scale based only on height",
			landscapeSize,
			data.Size{Width: 0, Height: 100},
			data.Size{Width: 200, Height: 100},
		},
		{
			"portrait scale based only on width",
			portraitSize,
			data.Size{Width: 100, Height: 0},
			data.Size{Width: 100, Height: 200},
		},
		{
			"portrait scale based only on height",
			portraitSize,
			data.Size{Width: 0, Height: 100},
			data.Size{Width: 50, Height: 100},
		},
		{
			"scale square based only on height",
			squareSize,
			data.Size{Width: 0, Height: 100},
			data.Size{Width: 100, Height: 100},
		},
		{
			"scale square based only on width",
			squareSize,
			data.Size{Width: 100, Height: 0},
			data.Size{Width: 100, Height: 100},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			newSize := CalculateScaledSize(
				testCase.src,
				testCase.newMax,
			)
			if newSize != testCase.expected {
				t.Errorf(
					"expceted %v got %v; CalculateNewSize(%v, %v)",
					testCase.expected,
					newSize,
					testCase.src,
					testCase.newMax,
				)
			}
		})
	}
}

func TestScaleImage(t *testing.T) {
	src := image.NewRGBA(image.Rect(0, 0, 2, 2))
	draw.Draw(src, src.Bounds(), &image.Uniform{color.RGBA{255, 0, 0, 255}}, image.Point{}, draw.Src)

	newSize := data.Size{Width: 4, Height: 4}
	result := ScaleImage(src, newSize)

	if result.Bounds().Dx() != newSize.Width || result.Bounds().Dy() != newSize.Height {
		t.Errorf(
			"expected %dx%d got %dx%d; result.Bound()",
			newSize.Width,
			newSize.Height,
			result.Bounds().Dx(),
			result.Bounds().Dy(),
		)
	}
}
