package main

import (
	"image/gif"
	"log"
	"os"
)

func main() {
	c, f, err := convertPNGToField("src.png")
	if err != nil {
		log.Fatal(err)
	}

	// Generate the floodfill gif
	anim := generateFloodfillGif(f, c, 10)

	// Create the gif file
	fle, _ := os.Create("floodfill.gif")
	// Save the gif on disk
	gif.EncodeAll(fle, &anim)
}
