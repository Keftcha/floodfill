package main

import (
	"image"
	"image/color"
	"image/gif"

	"github.com/keftcha/floodfill/field"
)

func generateFloodfillGif(f field.Field, c color.RGBA, delay int) gif.GIF {
	// Define our color palette
	palette := []color.Color{
		color.Transparent,
		color.White,
		color.Black,
		c,
	}

	// Create the initial frame
	img := image.NewPaletted(
		image.Rectangle{
			image.Point{0, 0},
			image.Point{f.Width, f.Height},
		},
		palette,
	)

	// Put the black and white on the initial frame
	for y := 0; y < f.Height; y++ {
		for x := 0; x < f.Width; x++ {
			coloridx := uint8(1)
			if f.Cells[y][x].Changed {
				coloridx = 2
			}
			img.SetColorIndex(x, y, coloridx)
		}
	}

	// Slice of our gif frame (also add the initial one)
	paletted := []*image.Paletted{img}
	delays := []int{delay}

	// Loop while the filed isn't filled
	for idx := 0; !f.Filled; idx++ {
		cls := f.Step()

		// Create the new frame
		// Create the initial frame
		frm := image.NewPaletted(
			image.Rectangle{
				image.Point{0, 0},
				image.Point{f.Width, f.Height},
			},
			palette,
		)

		for _, c := range cls {
			frm.SetColorIndex(c.X, c.Y, 3)
		}

		// Add the frame to our list
		paletted = append(paletted, frm)
		delays = append(delays, delay)
	}

	// Create the gif
	anim := gif.GIF{
		Image:           paletted,
		Delay:           delays,
		BackgroundIndex: 0,
	}

	return anim
}
