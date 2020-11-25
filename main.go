package main

import (
	"flag"
	"image/gif"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"strings"
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
	// Read the PNG file
	fle, err := ioutil.ReadFile(src)
	if err != nil {
		log.Fatal(err)
	}
	img, err := png.Decode(strings.NewReader(string(fle)))
	if err != nil {
		log.Fatal(err)
	}

	// Generate the floodfill gif
	anim := generateFloodfillGif(img, delay)

	// Create the gif file
	outFle, _ := os.Create(out)
	// Save the gif on disk
	gif.EncodeAll(outFle, &anim)
}
