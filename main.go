package main

import (
	"flag"
	"image/gif"
	"log"
	"os"

	"github.com/keftcha/floodfill/field"
)

var src string
var out string
var delay int = 10

func init() {
	flag.IntVar(&delay, "delay", 10, "Delay between each frame of the gif")

	flag.Parse()

	src = flag.Arg(0)
	out = flag.Arg(1)
	if src == "" || out == "" {
		log.Fatal(
			"Floodfill usage:\n\n" +
				"floodfill [--delay int] SOURCE DEST\n\n" +
				"SOURCE: png image source\n" +
				"DEST: gif image created",
		)
	}
}

func main() {

	c, f, err := field.ConvertPNGToField(src)
	if err != nil {
		log.Fatal(err)
	}

	// Generate the floodfill gif
	anim := generateFloodfillGif(f, c, delay)

	// Create the gif file
	fle, _ := os.Create(out)
	// Save the gif on disk
	gif.EncodeAll(fle, &anim)
}
