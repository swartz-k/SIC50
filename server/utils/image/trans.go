package image

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"image"
	"image/color"
	"log"
	"math"
	"os"
)

func Transfer(filename string) error {
	infile, err := os.Open(filename)

	if err != nil {
		log.Printf("failed opening %s: %s", filename, err)
		panic(err.Error())
	}
	defer infile.Close()

	imgSrc, _, err := image.Decode(infile)
	if err != nil {
		panic(err.Error())
	}

	// Create a new grayscale image
	bounds := imgSrc.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	grayScale := image.NewGray(image.Rectangle{Min: image.Point{}, Max: image.Point{X: w, Y: h}})
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			imageColor := imgSrc.At(x, y)
			rr, gg, bb, _ := imageColor.RGBA()
			r := math.Pow(float64(rr), 2.2)
			g := math.Pow(float64(gg), 2.2)
			b := math.Pow(float64(bb), 2.2)
			m := math.Pow(0.2125*r+0.7154*g+0.0721*b, 1/2.2)
			Y := uint16(m + 0.5)
			grayColor := color.Gray{uint8(Y >> 8)}
			grayScale.Set(x, y, grayColor)
		}
	}

	// edge
	//verticalFilter := [[-1,-2,-1], [0,0,0], [1,2,1]]
	//horizontalFilter := [[-1,0,1], [-2,0,2], [-1,0,1]]
	vf := mat.NewDense(3, 3, []float64{-1, -2, -1, 0, 0, 0, 1, 2, 1})
	hf := mat.NewDense(3, 3, []float64{-1, 0, 1, -2, 0, 2, -1, 0, 1})

	final := image.NewGray(image.Rectangle{Min: image.Point{}, Max: image.Point{X: w, Y: h}})
	fmt.Printf("%+v", final)
	for r := 3; r < w-2; r++ {
		for c := 3; c < h-2; c++ {
			local_pixels := grayScale.SubImage(image.Rectangle{
				Min: image.Point{X: r - 1, Y: c - 1},
				Max: image.Point{X: r + 2, Y: c + 2},
			})

			vf.Mul(local_pixels)
			//		local_pixels := grayScale.SubImage(image.Rectangle{Min: })
			//			img[row-1:row+2, col-1:col+2]
			//
			//		vertical_transformed_pixels = vf * local_pixels
			//		vertical_score = vertical_transformed_pixels.sum()
			//
			//		horizontal_transformed_pixels = horizontal_filter*local_pixels
			//		horizontal_score = horizontal_transformed_pixels.sum()
			//
			//		edge_score = (vertical_score**2 + horizontal_score**2)**.5
			//
			//		edge_score = (edge_score)**0.8
			//
			//		if edge_score >= 0.2:
			//		edge_score=edge_score**0.6
			//
			//		edges_img[row,col]= edge_score
		}
	}
	return nil
}
