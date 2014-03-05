Go AllRGB Painting Generator
============================

![colors1](https://f.cloud.github.com/assets/1121616/2332908/cb5524c6-a460-11e3-8ef8-851ac3015c01.png)

 > Look mom, I made art!

This Go program generates images that use every color in the RGB color palette exactly once, and tries to do so in as easthetically-pleasing way as possible.

It supports different modes. 

#### Tie-dye

```
go run tiedye.go
```

![tiedye](https://f.cloud.github.com/assets/1121616/2333097/4357032e-a464-11e3-98ac-08247aba1cf3.png)

#### Flower

```
go run flower.go
```

![flower](https://f.cloud.github.com/assets/1121616/2333178/4cc6d2b6-a466-11e3-950e-165768ebd2b4.png)

#### Pointillist

```
go run pointillist.go
```

![point](https://f.cloud.github.com/assets/1121616/2332998/69fe6f6e-a462-11e3-98c4-b59d00ea00e1.png)

### The Algorithm

The algorithm is inspired by [Joco's blog post](http://joco.name/2014/03/02/all-rgb-colors-in-one-image/). The idea is the following:
 - create a palette of all the available colors
 - start at one point, assign a random color to the point, and add the neighboring points to an available array
 - choose the next color randomly, and find the available spot whose neighborhood best matches the chosen color. Assign the color to this spot, add its neighbors to the available array, and repeat.

### The Juicy Details

Creating images pixel-by-pixel is pretty straight-forward in Go:

```go

import (
  "image"
  "image/draw"
  "os"
)

func main() {
  // create the new image file on disk and defer closing it
  toimg, _ := os.Create("colors.png")
  defer toimg.Close()
  
  // create a new drawing
  m := image.NewRGBA(image.Rect(0, 0, width, height))
  
  // choose the point to draw at
  p := image.Point{0, 0}
  
  // create the rectangle we want to draw
  r := image.Rectangle{
  	Min: image.Point{p.X, p.Y},
  	Max: image.Point{p.X + 1, p.Y + 1},
  }
  
  // add one pixel to the drawing
  draw.Draw(m, r, &image.Uniform{pallete[n]}, image.Point{0, 0}, draw.Src)
  
  // save the pixel to png (the rest will be white)
  png.Encode(toimg, m)
}
```

If we repeat this process in a loop for every color, we have a program that fills a png image with every color :) 

#### Final word

Change the parameters. Try different constraints. Expirement, make your own, and if you get stuck, ask questions! You can find me at [@ironzeb](https://twitter.com/ironzeb), or read my blog at (IronZebra)[http://ironzebra.com].

Personally, I'm going to print one of these [onto a canvas](http://canvaspop.com) very soon!
