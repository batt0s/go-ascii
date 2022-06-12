package asciicon

import (
	"fmt"
	"image"
	"reflect"
)

func rgbToGray(img image.Image) *image.Gray {
	var (
		bounds = img.Bounds()
		gray   = image.NewGray(bounds)
	)
	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			var rgba = img.At(x, y)
			gray.Set(x, y, rgba)
		}
	}
	return gray
}

func asciiInvert(gray uint8) string {
	if gray < 51 {
		return "\u2001"
	}
	if gray >= 51 && gray < 102 {
		return "\u2591"
	}
	if gray >= 102 && gray < 153 {
		return "\u2592"
	}
	if gray >= 153 && gray < 204 {
		return "\u2593"
	}
	if gray >= 204 {
		return "\u2591"
	}
	return ""
}

func Convert(img image.Image) {
	gray := rgbToGray(img)
	for x := 0; x < img.Bounds().Dx(); x++ {
		for y := 0; y < img.Bounds().Dy(); y++ {
			//pixel := gray.At(x, y)
			pixel := reflect.ValueOf(gray.At(x, y)).FieldByName("Y").Uint()
			char := asciiInvert(uint8(pixel))
			fmt.Printf("%s%s", char, char)
		}
		fmt.Print("\n")
	}
}
