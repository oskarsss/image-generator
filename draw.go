package main

import (
	"image"
	"image/color"
	"image/draw"
	"math"
)

func drawRect(img *image.RGBA, col color.RGBA) {
	draw.Draw(img, img.Bounds(), &image.Uniform{col}, image.Point{}, draw.Src)
}

func drawCircle(img *image.RGBA, col color.RGBA, width, height int) {
	centerX := width / 2
	centerY := height / 2
	radiusX := width / 2
	radiusY := height / 2

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			dx := float64(x - centerX)
			dy := float64(y - centerY)
			if (dx*dx)/(float64(radiusX)*float64(radiusX))+(dy*dy)/(float64(radiusY)*float64(radiusY)) <= 1 {
				img.Set(x, y, col)
			}
		}
	}
}

func drawTriangle(img *image.RGBA, col color.RGBA, width, height int) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Check if the pixel is inside the triangle
			if x > width-(y*width/height) && x < width+(y*width/height) {
				img.Set(x, y, col)
			}
		}
	}
}

func drawOctagram(img *image.RGBA, col color.RGBA, width, height int) {
	radius := float64(min(width, height)) / 2
	centerX, centerY := float64(width)/2, float64(height)/2
	sideLength := radius / math.Sqrt(2)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			dx := math.Abs(float64(x) - centerX)
			dy := math.Abs(float64(y) - centerY)
			if (dx+dy < radius) || (dx < sideLength && dy < sideLength) {
				img.Set(x, y, col)
			}
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
