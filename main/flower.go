package main

import (
	"fmt"
	"github.com/hermanschaaf/allrgb"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
)

func main() {
	rand.Seed(1)

	toimg, _ := os.Create("flower.png")
	defer toimg.Close()

	width := 512
	height := 256

	pallete := allrgb.GeneratePallete(width, height)
	m := image.NewRGBA(image.Rect(0, 0, width, height))

	availableArray := []image.Point{image.Point{width / 2, height / 2}}

	available := make(map[image.Point]bool)
	taken := make(map[image.Point]color.RGBA)

	for i := 0; i < width*height; i++ {
		if i%500 == 0 {
			fmt.Println(i * 100 / (width * height))
		}
		c := pallete[i]
		chosenIndex := allrgb.ChoosePoint(c, availableArray, taken)
		p := availableArray[chosenIndex]
		// set chosen point as unavailable
		allrgb.SetTaken(width, height, chosenIndex, p, c, &availableArray, available, taken)
		allrgb.DrawPoint(m, p, c)

		availableArray = append(availableArray[:chosenIndex], availableArray[chosenIndex+1:]...)
		if len(availableArray) == 0 {
			break
		}
	}

	png.Encode(toimg, m)
}
