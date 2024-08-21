package asciiimager

import (
	"image"

	"golang.org/x/image/draw"
)

// RescaleImage rescales the input image to the specified dimensions.
func RescaleImage(img image.Image, width, height int) image.Image {
	bounds := img.Bounds()
	if bounds.Max.X == width && bounds.Max.Y == height {
		return img
	}

	resizedImg := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.NearestNeighbor.Scale(resizedImg, resizedImg.Bounds(), img, img.Bounds(), draw.Src, nil)
	return resizedImg
}

// ResizeToContent resizes the image to the smallest size that contains all non-transparent pixels.
func ResizeToContent(img image.Image) image.Image {
	bounds := img.Bounds()
	minX, minY, maxX, maxY := bounds.Max.X, bounds.Max.Y, 0, 0

	for y := 0; y < bounds.Dy(); y++ {
		for x := 0; x < bounds.Dx(); x++ {
			_, _, _, alpha := img.At(x, y).RGBA()
			if alpha != 0 {
				minX = min(minX, x)
				minY = min(minY, y)
				maxX = max(maxX, x)
				maxY = max(maxY, y)
			}
		}
	}

	if minX == bounds.Max.X && minY == bounds.Max.Y {
		return img
	}

	return img.(interface {
		SubImage(r image.Rectangle) image.Image
	}).SubImage(image.Rect(minX, minY, maxX+1, maxY+1))
}
