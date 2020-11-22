package main

import (
	"image"
	"image/color"
	"image/gif"
	"os"

	"github.com/keftcha/floodfill/cell"
	"github.com/keftcha/floodfill/field"
)

func main() {
	// The field look like this
	// # # # # # # # # # #
	// #           ¤     #
	// #             ¤ ¤ #
	// #   ¤ ¤ ¤         #
	// #   ¤   ¤         #
	// #   ¤   ¤   S     #
	// #   ¤ ¤           #
	// #                 #
	// # # # # # # # # # #
	// # → border
	// ¤ → cell that have Changed to true (wall)
	// S → start cell
	cells := [][]cell.Cell{
		[]cell.Cell{
			cell.New(0, 0, false),
			cell.New(1, 0, false),
			cell.New(2, 0, false),
			cell.New(3, 0, false),
			cell.New(4, 0, false),
			cell.New(5, 0, true),
			cell.New(6, 0, false),
			cell.New(7, 0, false),
		},
		[]cell.Cell{
			cell.New(0, 1, false),
			cell.New(1, 1, false),
			cell.New(2, 1, false),
			cell.New(3, 1, false),
			cell.New(4, 1, false),
			cell.New(5, 1, false),
			cell.New(6, 1, true),
			cell.New(7, 1, true),
		},
		[]cell.Cell{
			cell.New(0, 2, false),
			cell.New(1, 2, true),
			cell.New(2, 2, true),
			cell.New(3, 2, true),
			cell.New(4, 2, false),
			cell.New(5, 2, false),
			cell.New(6, 2, false),
			cell.New(7, 2, false),
		},
		[]cell.Cell{
			cell.New(0, 3, false),
			cell.New(1, 3, true),
			cell.New(2, 3, false),
			cell.New(3, 3, true),
			cell.New(4, 3, false),
			cell.New(5, 3, false),
			cell.New(6, 3, false),
			cell.New(7, 3, false),
		},
		[]cell.Cell{
			cell.New(0, 4, false),
			cell.New(1, 4, true),
			cell.New(2, 4, false),
			cell.New(3, 4, true),
			cell.New(4, 4, false),
			cell.New(5, 4, false),
			cell.New(6, 4, false),
			cell.New(7, 4, false),
		},
		[]cell.Cell{
			cell.New(0, 5, false),
			cell.New(1, 5, true),
			cell.New(2, 5, true),
			cell.New(3, 5, false),
			cell.New(4, 5, false),
			cell.New(5, 5, false),
			cell.New(6, 5, false),
			cell.New(7, 5, false),
		},
		[]cell.Cell{
			cell.New(0, 6, false),
			cell.New(1, 6, false),
			cell.New(2, 6, false),
			cell.New(3, 6, false),
			cell.New(4, 6, false),
			cell.New(5, 6, false),
			cell.New(6, 6, false),
			cell.New(7, 6, false),
		},
	}

	f, _ := field.New(cells, cells[4][5])

	// Define our color palette
	palette := []color.Color{
		color.Transparent,
		color.White,
		color.Black,
		color.RGBA{0x00, 0x00, 0xFF, 0xFF},
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
			colorIdx := uint8(1)
			if f.Cells[y][x].Changed {
				colorIdx = 2
			}
			img.SetColorIndex(x, y, colorIdx)
		}
	}

	// Slice of our gif frame (also add the initial one)
	paletted := []*image.Paletted{img}
	delays := []int{10}

	// Loop while the files isn't filled
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
		delays = append(delays, 10)
	}

	// Create the gif
	anim := gif.GIF{
		Image:           paletted,
		Delay:           delays,
		BackgroundIndex: 0,
	}

	// Create the gif file
	fle, _ := os.Create("floodfill.gif")
	// Save the gif on disk
	gif.EncodeAll(fle, &anim)
}
