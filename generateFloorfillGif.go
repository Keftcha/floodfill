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

func generateFloodfillGif(img image.Image, delay int) gif.GIF {
	// Find image colors to make our color palette
	palette := color.Palette{}
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

	// The first and Second frame of the gif
	first := image.NewPaletted(
		image.Rectangle{
			image.Point{0, 0},
			image.Point{
				img.Bounds().Max.X,
				img.Bounds().Max.Y,
			},
		},
		palette,
	)
	second := image.NewPaletted(
		image.Rectangle{
			image.Point{0, 0},
			image.Point{
				img.Bounds().Max.X,
				img.Bounds().Max.Y,
			},
		},
		palette,
	)

	// White index of the palette color (used for building the first frame)
	_, whiteIdx := contain(palette, color.RGBA{255, 255, 255, 255})
	// Slice of pixel position we have filled and need to proceed
	toProceed := make([]image.Point, 0)
	for y := 0; y < img.Bounds().Max.Y; y++ {
		for x := 0; x < img.Bounds().Max.X; x++ {
			// Find the pixel color
			r, g, b, a := img.At(x, y).RGBA()
			c := color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}

			// Finde the idx of the color in our palette
			_, idx := contain(palette, c)
			// Check if the color is different of black and white
			if (c != color.RGBA{255, 255, 255, 255} && c != color.RGBA{0, 0, 0, 255}) {
				first.SetColorIndex(x, y, whiteIdx)
				second.SetColorIndex(x, y, idx)
				toProceed = append(toProceed, image.Point{x, y})
			} else {
				first.SetColorIndex(x, y, idx)
				second.SetColorIndex(x, y, idx)
			}
		}
	}

	// Slice of our gif frame
	frames := []*image.Paletted{first, second}
	// Delays of each frame of our gif
	delays := []int{delay, delay}

	// Floodfill loop
	current := second
	for len(toProceed) != 0 {
		// Create the next frame
		next := image.NewPaletted(
			image.Rectangle{
				image.Point{0, 0},
				image.Point{
					img.Bounds().Max.X,
					img.Bounds().Max.Y,
				},
			},
			palette,
		)
		// Copy the content of the current frame in the next one
		copy(next.Pix, current.Pix)

		// Slice of next pixel position we will need to proceed next step
		toProceedNext := make([]image.Point, 0)

		for _, pixCoor := range toProceed {
			// Color idx in the palette of the pixel we proceed
			r, g, b, a := next.At(pixCoor.X, pixCoor.Y).RGBA()
			_, colorIdx := contain(palette, color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)})

			// Pixel above
			if 0 < pixCoor.Y {
				above := next.At(pixCoor.X, pixCoor.Y-1)
				// If the above pixel is white, we color it
				if (above == color.RGBA{255, 255, 255, 255}) {
					// Change the color
					next.SetColorIndex(pixCoor.X, pixCoor.Y-1, colorIdx)
					// Add to the proceed list for next step
					toProceedNext = append(toProceedNext, image.Point{pixCoor.X, pixCoor.Y - 1})
				}
			}
			// Pixel on right
			if pixCoor.X < next.Bounds().Max.X-1 {
				right := next.At(pixCoor.X+1, pixCoor.Y)
				// If the right pixel is white, we color it
				if (right == color.RGBA{255, 255, 255, 255}) {
					// Change the color
					next.SetColorIndex(pixCoor.X+1, pixCoor.Y, colorIdx)
					// Add to the proceed list for next step
					toProceedNext = append(toProceedNext, image.Point{pixCoor.X + 1, pixCoor.Y})
				}

			}
			// Pixel bellow
			if pixCoor.Y < next.Bounds().Max.Y-1 {
				bellow := next.At(pixCoor.X, pixCoor.Y+1)
				// If the pixel bellow is white, we colore it
				if (bellow == color.RGBA{255, 255, 255, 255}) {
					// Change the color
					next.SetColorIndex(pixCoor.X, pixCoor.Y+1, colorIdx)
					// Add to the proceed list for next step
					toProceedNext = append(toProceedNext, image.Point{pixCoor.X, pixCoor.Y + 1})
				}
			}
			// Pixel on left
			if 0 < pixCoor.X {
				left := next.At(pixCoor.X-1, pixCoor.Y)
				// If the pixel on left is white, we color it
				if (left == color.RGBA{255, 255, 255, 255}) {
					// Change the color
					next.SetColorIndex(pixCoor.X-1, pixCoor.Y, colorIdx)
					// Add to the proceed list for next step
					toProceedNext = append(toProceedNext, image.Point{pixCoor.X - 1, pixCoor.Y})
				}
			}
		}

		// Add the new builded frame to our frame list
		frames = append(frames, next)
		// Add the delay for the frame
		delays = append(delays, delay)

		// Update the new proceed pixels
		toProceed = toProceedNext
		// Now the next frame is the current
		current = next
	}

	return gif.GIF{
		Image: frames,
		Delay: delays,
	}
}
