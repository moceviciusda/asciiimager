package asciiimager

import (
	"fmt"
	"image"
	"image/color"
)

// const asciiTexture = " .;coPO?$@"
const asciiTexture = " .`^\",:;Il!i<>~+_-.?][}{1)(|/tfjrxnucvzxYUJCLQ0OZmwqpkdbhao*#MW&8%B@$"

func ImageToAsciiShader(img image.Image, width, height int) []byte {
	img = RescaleImage(img, width, height)
	bounds := img.Bounds()
	ascii := make([]byte, 0, bounds.Dx()*bounds.Dy())
	for y := 0; y < bounds.Dy(); y++ {
		for x := 0; x < bounds.Dx(); x++ {
			ascii = append(ascii, PixelToAscii(img.At(x, y)))
		}
		ascii = append(ascii, '\n')
	}
	return ascii
}

func ImageToAnsiShader(img image.Image, width, height int) []byte {
	img = RescaleImage(img, width, height)
	bounds := img.Bounds()
	ascii := make([]byte, 0, bounds.Dx()*bounds.Dy())
	for y := 0; y < bounds.Dy(); y++ {
		for x := 0; x < bounds.Dx(); x++ {
			ascii = append(ascii, PixelToAnsi(img.At(x, y))...)
		}
		ascii = append(ascii, '\n')
	}
	return ascii
}

func PixelToAscii(c color.Color) byte {
	r, g, b, a := c.RGBA()
	gray := 0.2126*float64(r) + 0.7152*float64(g) + 0.0722*float64(b)

	if a == 0 {
		return ' '
	}

	index := int(gray) * (len(asciiTexture) - 1) / 0xffff

	return asciiTexture[index]
}

func PixelToAnsi(c color.Color) string {
	r, g, b, a := c.RGBA()
	if a == 0 {
		return " "
	}

	gray := 0.2126*float64(r) + 0.7152*float64(g) + 0.0722*float64(b)
	index := int(gray) * (len(asciiTexture) - 1) / 0xffff

	ansiColor := fmt.Sprintf("%d;%d;%d", r, g, b)

	return "\x1b[38;2;" + ansiColor + "m" + string(asciiTexture[index]) + "\x1b[0m"
}
