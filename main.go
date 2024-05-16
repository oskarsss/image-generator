package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", formHandler)
	http.HandleFunc("/generate", generateHandler)

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}

func generateHandler(w http.ResponseWriter, r *http.Request) {
	widthStr := r.URL.Query().Get("width")
	heightStr := r.URL.Query().Get("height")
	redStr := r.URL.Query().Get("red")
	greenStr := r.URL.Query().Get("green")
	blueStr := r.URL.Query().Get("blue")
	shape := r.URL.Query().Get("shape")

	width, err := strconv.Atoi(widthStr)
	if err != nil || width <= 0 {
		http.Error(w, "Invalid width", http.StatusBadRequest)
		return
	}

	height, err := strconv.Atoi(heightStr)
	if err != nil || height <= 0 {
		http.Error(w, "Invalid height", http.StatusBadRequest)
		return
	}

	red, err := strconv.Atoi(redStr)
	if err != nil || red < 0 || red > 255 {
		http.Error(w, "Invalid red value", http.StatusBadRequest)
		return
	}

	green, err := strconv.Atoi(greenStr)
	if err != nil || green < 0 || green > 255 {
		http.Error(w, "Invalid green value", http.StatusBadRequest)
		return
	}

	blue, err := strconv.Atoi(blueStr)
	if err != nil || blue < 0 || blue > 255 {
		http.Error(w, "Invalid blue value", http.StatusBadRequest)
		return
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	transparent := color.RGBA{0, 0, 0, 0}
	col := color.RGBA{uint8(red), uint8(green), uint8(blue), 255}

	draw.Draw(img, img.Bounds(), &image.Uniform{transparent}, image.Point{}, draw.Src)

	switch shape {
	case "circle":
		drawCircle(img, col, width, height)
	case "triangle":
		drawTriangle(img, col, width, height)
	case "octagram":
		drawOctagram(img, col, width, height)
	default:
		drawRect(img, col)
	}

	w.Header().Set("Content-Type", "image/png")
	png.Encode(w, img)
}
