package transforms

import (
	"image"

	"github.com/eynopv/image-scaler/internal/data"
)

func ScaleImage(img image.Image, newSize data.Size) image.Image {
	scaled := image.NewRGBA(image.Rect(0, 0, newSize.Width, newSize.Height))
	scaleX := float64(img.Bounds().Dx()) / float64(newSize.Width)
	scaleY := float64(img.Bounds().Dy()) / float64(newSize.Height)

	for y := 0; y < newSize.Height; y++ {
		for x := 0; x < newSize.Width; x++ {
			imgX := int(float64(x) * scaleX)
			imgY := int(float64(y) * scaleY)
			scaled.Set(x, y, img.At(imgX, imgY))
		}
	}

	return scaled
}

func CalculateScaledSize(src data.Size, newMax data.Size) data.Size {
	ratio := float64(src.Width) / float64(src.Height)
	newWidth, newHeight := src.Width, src.Height

	if newMax.Width > 0 {
		newHeight = int(float64(newMax.Width) / ratio)
		if newHeight > newMax.Height && newMax.Height > 0 {
			newHeight = newMax.Height
			newWidth = int(float64(newMax.Height) * ratio)
		} else {
			newWidth = newMax.Width
		}
	}

	if newMax.Height > 0 {
		newWidth = int(float64(newMax.Height) * ratio)
		if newWidth > newMax.Width && newMax.Width > 0 {
			newWidth = newMax.Width
			newHeight = int(float64(newMax.Width) / ratio)
		} else {
			newHeight = newMax.Height
		}
	}

	newSize := data.Size{
		Width:  newWidth,
		Height: newHeight,
	}

	return newSize
}
