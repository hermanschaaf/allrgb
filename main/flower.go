package main

import (
	"flag"
	"fmt"
	"github.com/hermanschaaf/allrgb"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
)

var filename string
var width int
var height int
var seed int64

func main() {
	flag.StringVar(&filename, "file", "flower", "name of the file (.png will be added)")
	flag.IntVar(&width, "width", 512, "width of the image")
	flag.IntVar(&height, "height", 256, "height of the image")
	flag.Int64Var(&seed, "seed", 1, "random seed")

	rand.Seed(seed)

	flag.Parse()

	fmt.Println("File will be:", filename+".png")

	toimg, _ := os.Create(filename + ".png")
	defer toimg.Close()

	pallete := allrgb.GeneratePallete(width, height)
	m := image.NewRGBA(image.Rect(0, 0, width, height))

	availableArray := []image.Point{image.Point{width / 2, height / 2}}

	available := make(map[image.Point]bool)
	taken := make(map[image.Point]color.RGBA)

	for i := 0; i < width*height; i++ {
		if i%500 == 0 {
			fmt.Print(i*100/(width*height), "..")
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
