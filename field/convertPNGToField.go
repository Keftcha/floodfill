package field

import (
	"image/color"
	"image/png"
	"io/ioutil"
	"strings"

	"github.com/keftcha/floodfill/field/cell"
)

func ConvertPNGToField(pngName string) (color.RGBA, Field, error) {
	// Read the PNG file
	fle, err := ioutil.ReadFile(pngName)
	if err != nil {
		return color.RGBA{}, Field{}, err
	}

	img, err := png.Decode(strings.NewReader(string(fle)))
	if err != nil {
		return color.RGBA{}, Field{}, err
	}

	cells := make([][]cell.Cell, img.Bounds().Max.Y)
	initCell := cell.New(0, 0, false)
	var initColor color.RGBA // init color we return
	// Travel image pixels
	for lineIdx := img.Bounds().Min.Y; lineIdx < img.Bounds().Max.Y; lineIdx++ {
		line := make([]cell.Cell, img.Bounds().Max.X)
		for columnIdx := img.Bounds().Min.X; columnIdx < img.Bounds().Max.X; columnIdx++ {
			// Get the color of the pixel
			c := img.At(columnIdx, lineIdx)
			changed := c == color.RGBA{0, 0, 0, 255} // Check if it's a border

			// Create the pixel in our cell field
			line[columnIdx] = cell.New(columnIdx, lineIdx, changed)

			// Check if the cell is different of black or white
			// That is a start cell of his color
			if (c != color.RGBA{255, 255, 255, 255} && c != color.RGBA{0, 0, 0, 255}) {
				initCell.X = columnIdx
				initCell.Y = lineIdx

				// Initial color
				r, g, b, a := c.RGBA()
				initColor = color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
			}
		}
		cells[lineIdx] = line
	}

	field, err := New(cells, cells[initCell.Y][initCell.X])
	return initColor, field, err
}
