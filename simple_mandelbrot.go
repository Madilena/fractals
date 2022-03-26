package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

func main() {
	const (
		//define the height and width of the image
		width  = 2000
		height = 2000
		//define the max and min for complex plane
		xmin = -2
		ymin = -2
		xmax = +2
		ymax = +2
	)

	//image
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	//iterate over each pixel in image
	//then map each pixel to a coordinate on complex plane
	for px := 0; px < width; px++ {

		//transform px to x
		x := float64(px)*(xmax-xmin)/width + xmin

		for py := 0; py < height; py++ {

			//transform py to y
			y := float64(py)*(ymax-ymin)/height + ymin

			c := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(c))
		}

	}

	// Encode as PNG.
	f, _ := os.Create("mandelbrot.png")
	png.Encode(f, img)

}

func mandelbrot(c complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var z complex128
	for n := uint8(0); n < iterations; n++ {
		z = z*z + c
		if cmplx.Abs(z) > 2 {
			switch {
			case n > 50:
				//blue martine
				return color.RGBA{20, 205, 200, 0xff}
			default:
				// logarithmic blurple gradient to show small differences on the
				// periphery of the fractal.
				logScale := math.Log(float64(n)) / math.Log(float64(iterations))
				return color.RGBA{85, 57, 204 - uint8(logScale*250), 0xff}
			}
		}
	}
	//diverges to infinity
	return color.Black
}
