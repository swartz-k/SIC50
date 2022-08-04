package image

import (
	"fmt"
	"github.com/BioChemML/SIC50/server/utils/log"
	"github.com/pkg/errors"
	_ "golang.org/x/image/webp"
	"image"
	"image/jpeg"
	"image/png"

	"os"
	"strings"
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

var magicTable = map[string]string{
	"\xff\xd8\xff":      "image/jpeg",
	"\x89PNG\r\n\x1a\n": "image/png",
}

func generateName(f, suffix string) string {
	ar := strings.Split(f, ".")
	arL := len(ar)
	if arL == 1 {
		return fmt.Sprintf("%s_%s", ar[0], suffix)
	} else if len(ar) == 2 {
		return fmt.Sprintf("%s_%s.%s", ar[0], suffix, ar[1])
	} else {
		n := strings.Join(ar[:len(ar)-1], ".")
		return fmt.Sprintf("%s_%s.%s", n, suffix, ar[len(ar)-1])
	}
}

func Split(f string, w, h int) (int, error) {
	img, err := os.Open(f)
	if err != nil {
		return 0, errors.Wrap(err, "open image")
	}
	defer img.Close()

	c, s, err := image.DecodeConfig(img)
	if err != nil {
		return 0, err
	}
	if c.Width < w || c.Height < h {
		tFName := generateName(f, fmt.Sprintf("%d-%d", 0, 0))
		os.Rename(f, tFName)
		return 1, nil
	}
	//log.Info("imagee %s w %d, h %d, color %+v, s %s,  \n", f, c.Width, c.Height, c.ColorModel, s)
	var imgSrc image.Image

	img.Seek(0, 0)
	if s == "png" {
		imgSrc, err = png.Decode(img)
	} else {
		imgSrc, err = jpeg.Decode(img)
	}
	if err != nil {
		return 0, errors.Wrapf(err, "decode with format %s", s)
	}

	sImg, ok := imgSrc.(interface {
		SubImage(r image.Rectangle) image.Image
	})
	if !ok {
		return 0, fmt.Errorf("image %s cannot get sub image", f)
	}

	hi := c.Height / h
	wi := c.Width / w
	for i := 0; i < wi; i++ {
		for j := 0; j < hi; j++ {
			xi := i
			yj := j
			tImg := sImg.SubImage(image.Rect(xi*w, yj*h, (xi+1)*w, (yj+1)*h))
			tFName := generateName(f, fmt.Sprintf("%d-%d", xi, yj))
			tF, err := os.Create(tFName)
			if err != nil {
				log.Info("split w:%d, h:%d failed %+v", i, j, err)
				break
			}
			defer tF.Close()
			if s == "png" {
				err = png.Encode(tF, tImg)
			} else {
				err = jpeg.Encode(tF, tImg, nil)
			}
			if err != nil {
				log.Info("encode splited image to f %s failed %+v", tFName, err)
			}
		}
	}

	return wi * hi, nil
}
