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
	flag.StringVar(&filename, "file", "pointillist", "name of the file (.png will be added)")
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

	availableArray := []allrgb.Neighborhood{}

	available := make(map[image.Point]bool)
	taken := make(map[image.Point]color.RGBA)

	points := 16 * 16
	for i := 0; i < points; i++ {
		x, y := i%16, i/16
		p := image.Point{x * (width / 16), y * (height / 16)}
		c := pallete[i]
		allrgb.SetTaken(width, height, 0, p, c, &availableArray, available, taken)
		allrgb.DrawPoint(m, p, c)
	}

	prev := 0
	for i := points; i < width*height; i++ {
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
