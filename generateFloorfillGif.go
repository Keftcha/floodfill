package main

import (
	"image"
	"image/color"
	"image/gif"
)

func contain(colors color.Palette, c color.RGBA) (bool, uint8) {
	for idx, clr := range colors {
		if clr == c {
			return true, uint8(idx)
		}
	}
	return false, 0
}

func extractPalette(img image.Image) color.Palette {
	palette := color.Palette{color.Transparent}
	for y := 0; y < img.Bounds().Max.Y; y++ {
		for x := 0; x < img.Bounds().Max.X; x++ {
			// Find the pixel color
			r, g, b, a := img.At(x, y).RGBA()
			c := color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
			// Add the pixel color if needed in our palette color
			if ok, _ := contain(palette, c); len(palette) == 0 || !ok {
				palette = append(palette, c)
			}
		}
	}
	return palette
}

func copyImageIntoPaletted(img image.Image, frame *image.Paletted) {
	model := frame.ColorModel()
	for y := 0; y < img.Bounds().Max.Y; y++ {
		for x := 0; x < img.Bounds().Max.X; x++ {
			idxColor := uint8(frame.Palette.Index(model.Convert(img.At(x, y))))
			frame.SetColorIndex(x, y, idxColor)
		}
	}
}

func initCells(frame *image.Paletted) []image.Point {
	whiteIdx := uint8(frame.Palette.Index(color.RGBA{255, 255, 255, 255}))
	blackIdx := uint8(frame.Palette.Index(color.RGBA{0, 0, 0, 255}))
	cells := make([]image.Point, 0)

	for y := 0; y < frame.Bounds().Max.Y; y++ {
		for x := 0; x < frame.Bounds().Max.X; x++ {
			pixel := frame.ColorIndexAt(x, y)
			if pixel != whiteIdx && pixel != blackIdx {
				cells = append(cells, image.Pt(x, y))
			}
		}
	}

	return cells
}

func addFrame(g *gif.GIF, f *image.Paletted, d int) {
	g.Image = append(g.Image, f)
	g.Delay = append(g.Delay, d)
}

func generateFloodfillGif(img image.Image, delay int) *gif.GIF {
	// Find image colors to make our color palette
	palette := extractPalette(img)

	// Initialize the first frame of the gif
	frame := image.NewPaletted(img.Bounds(), palette)
	copyImageIntoPaletted(img, frame)

	// Initialize pixels we need to proceed at first step
	pixels := initCells(frame)

	// Initialize our output gif
	out := &gif.GIF{
		Image: make([]*image.Paletted, 0, 1),
		Delay: make([]int, 0, 1),
	}
	addFrame(out, frame, delay)

	// Floodfill loop
	for len(pixels) != 0 {
		// Create the next frame
		nextFrame := image.NewPaletted(frame.Bounds(), palette)
		copy(nextFrame.Pix, frame.Pix)
		// Slice of next pixel position we will need to proceed next step
		nextPixels := make([]image.Point, 0)

		for _, pix := range pixels {
			// Color idx in the palette of the pixel we proceed
			colorIdx := uint8(nextFrame.Palette.Index(nextFrame.At(pix.X, pix.Y)))

			// Pixel above
			if 0 < pix.Y {
				above := nextFrame.At(pix.X, pix.Y-1)
				// If the above pixel is white, we color it
				if (above == color.RGBA{255, 255, 255, 255}) {
					// Change the color
					nextFrame.SetColorIndex(pix.X, pix.Y-1, colorIdx)
					// Add to the proceed list for next step
					nextPixels = append(nextPixels, image.Point{pix.X, pix.Y - 1})
				}
			}
			// Pixel on right
			if pix.X < nextFrame.Bounds().Max.X-1 {
				right := nextFrame.At(pix.X+1, pix.Y)
				// If the right pixel is white, we color it
				if (right == color.RGBA{255, 255, 255, 255}) {
					// Change the color
					nextFrame.SetColorIndex(pix.X+1, pix.Y, colorIdx)
					// Add to the proceed list for next step
					nextPixels = append(nextPixels, image.Point{pix.X + 1, pix.Y})
				}

			}
			// Pixel bellow
			if pix.Y < nextFrame.Bounds().Max.Y-1 {
				bellow := nextFrame.At(pix.X, pix.Y+1)
				// If the pixel bellow is white, we colore it
				if (bellow == color.RGBA{255, 255, 255, 255}) {
					// Change the color
					nextFrame.SetColorIndex(pix.X, pix.Y+1, colorIdx)
					// Add to the proceed list for next step
					nextPixels = append(nextPixels, image.Point{pix.X, pix.Y + 1})
				}
			}
			// Pixel on left
			if 0 < pix.X {
				left := nextFrame.At(pix.X-1, pix.Y)
				// If the pixel on left is white, we color it
				if (left == color.RGBA{255, 255, 255, 255}) {
					// Change the color
					nextFrame.SetColorIndex(pix.X-1, pix.Y, colorIdx)
					// Add to the proceed list for next step
					nextPixels = append(nextPixels, image.Point{pix.X - 1, pix.Y})
				}
			}
		}

		addFrame(out, frame, delay)

		// Update the new proceed pixels
		pixels = nextPixels
		// Now the next frame is the current
		frame = nextFrame
	}

	return out
}
