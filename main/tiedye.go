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

func setTaken(width, height, chosenIndex int, p image.Point, c color.RGBA, availableArray *[]image.Point, available map[image.Point]bool, taken map[image.Point]color.RGBA) {
	taken[p] = c
	available[p] = false

	// set next x point
	newPx := image.Point{p.X + 1, p.Y}
	_, haveTaken := taken[newPx]
	if newPx.X < width && available[newPx] != true && !haveTaken {
		*availableArray = append(*availableArray, newPx)
		available[newPx] = true
	}
	// set next y point
	newPy := image.Point{p.X, p.Y + 1}
	_, haveTaken = taken[newPy]
	if newPy.Y < height && available[newPy] != true && !haveTaken {
		*availableArray = append(*availableArray, newPy)
		available[newPy] = true
	}
}

var filename string
var width int
var height int
var seed int64

func main() {
	flag.StringVar(&filename, "file", "tiedye", "name of the file (.png will be added)")
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

	availableArray := []allrgb.Neighborhood{allrgb.Neighborhood{image.Point{0, 0}, color.RGBA{0, 0, 0, 0}}}

	available := make(map[image.Point]bool)
	taken := make(map[image.Point]color.RGBA)

	prev := 0
	for i := 0; i < width*height; i++ {
		perc := i * 100 / (width * height)
		if perc > prev {
			fmt.Print(perc, "..")
			prev = perc
		}
		c := pallete[i]
		chosenIndex := allrgb.ChoosePoint(c, availableArray, taken)
		p := availableArray[chosenIndex]
		// set chosen point as unavailable
		allrgb.SetTaken(width, height, chosenIndex, p.Point, c, &availableArray, available, taken)
		allrgb.DrawPoint(m, p.Point, c)

		availableArray = append(availableArray[:chosenIndex], availableArray[chosenIndex+1:]...)
		if len(availableArray) == 0 {
			break
		}
	}

	png.Encode(toimg, m)
}
