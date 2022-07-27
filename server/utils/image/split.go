package image

import (
	"fmt"
	"image"
	"image/png"
	"io"
)

// Pixel struct example
type Pixel struct {
	R int
	G int
	B int
	A int
}

func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
	return Pixel{int(r / 257), int(g / 257), int(b / 257), int(a / 257)}
}

func Split(r io.Reader) (image.Image, error) {
	//c, s, err := image.DecodeConfig(r)
	//if err != nil {
	//	return nil, err
	//}
	//fmt.Printf("w %d, h %d, color %+v, s %s,  \n", c.Width, c.Height, c.ColorModel, s)
	img, err := png.Decode(r)
	if err != nil {
		return nil, err
	}
	fmt.Println(rgbaToPixel(img.At(0, 1).RGBA()))
	return nil, nil
}
