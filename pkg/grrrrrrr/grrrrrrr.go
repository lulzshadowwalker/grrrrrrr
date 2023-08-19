package grrrrrrr

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"log"
	"math"
	"strings"

	"github.com/lulzshadowwalker/grrrrrrr/pkg/enum"
	gm "github.com/lulzshadowwalker/grrrrrrr/pkg/math"
)

type ColorChannel int

const (
	Red ColorChannel = iota
	Green
	Blue
)

func (c ColorChannel) Validate() error {	
	switch c {
	case Red, Green, Blue: return nil

	default: return &enum.EnumError{
			Enum: c, 
			Err: enum.ErrInvalidEnum,
		} 
	}
}

func StringToColorChannel(s string) (ColorChannel, error) {
	switch strings.ToLower(s) {
	case "red": return Red, nil
	case "green": return Green, nil
	case "blue": return Blue, nil
	
	default: return -1, fmt.Errorf("invalid color channel\nvalid color channels are red, green, and blue %w", enum.ErrInvalidEnum) 
	}	
}

// *quick and dirty*
//
// the human eye doesn't conceive all color "channels" with the same
// sensitivty
//
// this method doesn't take this into consideration so the luminosity
// turns out poor
func Average(m image.Image) (image.Image, error) {
	return process(m, func(c color.Color) color.Color {
		r, g, b, _ := decodeColor(c)
		p := uint8((int(r) + int(g) + int(b)) / 3)

		return color.NRGBA{p, p, p, 255}
	})
}

// takes into consideration the sensitivty the human eye perceives each color channel
func Luma(m image.Image) (image.Image, error) {
	const (
		rf float64 = 0.2126
		gf = 0.7152
		bf = 0.0722
	)	

	return process(m, func(c color.Color) color.Color {
		r, g, b, _ := decodeColor(c)
		p := uint8(float64(r) * rf + float64(g) * gf + float64(b) * bf)

		return color.NRGBA{p, p, p, 255}
	})
}

// RGB -> HSL -> Saturation = 0
//
// flatter, softer, least contrast
func Desaturate(m image.Image) (image.Image, error) {
	return process(m, func(c color.Color) color.Color {
		r, g, b, _ := decodeColor(c)

		min := gm.Min(gm.Min(r, g), b) 
		max := gm.Max(gm.Max(r, g), b)
		p := uint8((int(max) + int(min)) / 2)
		return color.NRGBA{p, p, p, 255}
	})
}

// maximum channel value 
//
// results in a brighter output than [decomposeMin]
func DecomposeMax(m image.Image) (image.Image, error) {
		return process(m, func(c color.Color) color.Color {
		r, g, b, _ := decodeColor(c) 

		p := uint8(gm.Max(gm.Max(r, g), b))
		return color.NRGBA{p, p, p, 255}
	})
}

// minimum channel value
// 
// results in a darker output than [decomposeMax]
func DecomposeMin(m image.Image) (image.Image, error) {
		return process(m, func(c color.Color) color.Color {
		r, g, b, _ := decodeColor(c)

		p := uint8(gm.Min(gm.Min(r, g), b))
		return color.NRGBA{p, p, p, 255}
	})
}

func SingleChannel(m image.Image, ch ColorChannel) (image.Image, error) {
		return process(m, func(c color.Color) color.Color {
		r, g, b, _ := decodeColor(c)

		err := ch.Validate()
		if errors.Is(err, enum.ErrInvalidEnum) {
			log.Fatal("single channel value mode can only be used with red, green, and blue channels")
		} else if err != nil {
			log.Fatalf("error processing the image %q", err)
		}
		
		p := r 
		switch ch {
		case Red:
			p = r 
		case Green: 
			p = g
		case Blue:
			p = b
		}

		p8 := uint8(p)	
		return color.NRGBA{p8, p8, p8, 255}
	})
}

func Shades(m image.Image, shadeCount int) (image.Image, error) {
	return process(m, func(c color.Color) color.Color {
		r, g, b, _ := decodeColor(c)
		est := float64((int(r) + int(g) + int(b)) / 3)

		// we estimate the gray color for a given pixel using one technique ("algo" since we are programmers) or another
		// then we map that estimation to a value among a limited set of shades that is uniformally distributed over the
		// the range of [0, 255] based on [shadeCount]

		// e.g. in the case of two shades (black and white essentially) if our estimation is greater than or equal to the
		// `covnersionFactor / 2` the resultant color would be white (brighter) and black otherwise
		conversionFactor := float64(255 / (shadeCount-1))
		p := uint8(math.Floor(est / conversionFactor + .5) * conversionFactor)
		
		return color.NRGBA{p, p, p, 255}
	})
}

// prefer returning [color.NRGBA] from the [processor] callback for better performance as
// otherwise most of the processing time would go into converting it to such
func process(m image.Image, processor func(color.Color) color.Color) (image.Image, error) {
	if m == nil {
		return nil, errors.New("no image was passed to the function")
	}

	b := m.Bounds()
	out := image.NewNRGBA(image.Rect(b.Min.X, b.Min.Y, b.Max.X, b.Max.Y))

	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			c := processor(m.At(x, y))
 			out.Set(x, y, c) 
		}
	}

	return out, nil
}

func decodeColor(c color.Color) (r, g, b, a uint8) {
	rr, gg, bb, aa := c.RGBA()
	return uint8(rr >> 8), uint8(gg >> 8), uint8(bb >> 8), uint8(aa >> 8)
}

