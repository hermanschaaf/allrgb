package allrgb

import (
	"image"
	"image/color"
	"image/draw"
	"math/rand"
	"sort"
)

const MAXDIFF int = 255*255 + 255*255 + 255*255

type RandomColor struct {
	Rand float64
	C    color.RGBA
}

type ByRand []RandomColor

func (a ByRand) Len() int           { return len(a) }
func (a ByRand) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByRand) Less(i, j int) bool { return a[i].Rand < a[j].Rand }

// generatePallete generates an array of all the colors
// we can use exactly once
func GeneratePallete(width, height int) (p []color.RGBA) {
	rands := []RandomColor{}
	for i := 0; i < width*height; i++ {
		r := i << 3 & 0xF8
		g := i >> 2 & 0xF8
		b := i >> 7 & 0xF8
		rands = append(rands, RandomColor{rand.Float64(), color.RGBA{uint8(r), uint8(g), uint8(b), 255}})
	}

	// shuffle pallete
	sort.Sort(ByRand(rands))

	for i := 0; i < len(rands); i++ {
		p = append(p, rands[i].C)
	}
	return p
}

func ColorDiff(c1, c2 color.RGBA) int {
	r := int(c1.R) - int(c2.R)
	g := int(c1.G) - int(c2.G)
	b := int(c1.B) - int(c2.B)
	return r*r + g*g + b*b
}

func ChoosePoint(c color.RGBA, available []image.Point, taken map[image.Point]color.RGBA) (i int) {
	best := 0
	bestDiff := MAXDIFF
	for a, point := range available {
		R, G, B := 0, 0, 0
		found := 0
		for x := -1; x <= 1; x++ {
			for y := -1; y <= 1; y++ {
				neighbor, t := taken[image.Point{point.X + x, point.Y + y}]
				if t {
					R += int(neighbor.R)
					G += int(neighbor.G)
					B += int(neighbor.B)
					found += 1
				}
			}
		}
		// average it out
		if found <= 0 {
			continue
		}
		R /= found
		G /= found
		B /= found

		diff := ColorDiff(c, color.RGBA{uint8(R), uint8(G), uint8(B), 255})
		if diff < bestDiff || (diff == bestDiff && rand.Intn(2) == 0) {
			best = a
			bestDiff = diff
		}
	}
	return best
}

func SetTaken(width, height, chosenIndex int, p image.Point, c color.RGBA, availableArray *[]image.Point, available map[image.Point]bool, taken map[image.Point]color.RGBA) {
	taken[p] = c
	available[p] = false

	// set next x point
	newPx := image.Point{p.X + 1, p.Y}
	_, haveTaken := taken[newPx]
	if newPx.X < width && available[newPx] != true && !haveTaken {
		*availableArray = append(*availableArray, newPx)
		available[newPx] = true
	}
	// set next -x point
	newPx = image.Point{p.X - 1, p.Y}
	_, haveTaken = taken[newPx]
	if newPx.X >= 0 && available[newPx] != true && !haveTaken {
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
	// set next -y point
	newPy = image.Point{p.X, p.Y - 1}
	_, haveTaken = taken[newPy]
	if newPy.Y >= 0 && available[newPy] != true && !haveTaken {
		*availableArray = append(*availableArray, newPy)
		available[newPy] = true
	}
}

func DrawPoint(m draw.Image, p image.Point, c color.RGBA) {
	r := image.Rectangle{
		Min: image.Point{p.X, p.Y},
		Max: image.Point{p.X + 1, p.Y + 1},
	}
	draw.Draw(m, r, &image.Uniform{c}, image.Point{0, 0}, draw.Src)
}
