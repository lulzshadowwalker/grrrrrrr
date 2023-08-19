package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"io"
	"log"
	"os"

	c "github.com/lulzshadowwalker/grrrrrrr/internal/config"
	"github.com/lulzshadowwalker/grrrrrrr/pkg/grrrrrrr"
)

func main() {
	file, err := os.Open(c.GetSrc())
	if err != nil {
		log.Fatalf("failed to open image file %q", err)
	}
	defer file.Close()

	reader := io.Reader(file)
	m, f, err := image.Decode(reader)
	if f != "jpeg" && f != "png" {
		log.Fatalf("%q format is not supported you can only use .jpeg/.png files", f)
	}
	if err != nil {
		log.Fatalf("could not decode image %q", err)
	}

	result, err := os.Create(c.GetDest())
	if err != nil {
		log.Fatalf("failed to create file at destination %q", err)
	}
	defer result.Close()

	var out image.Image
	switch c.GetMethod() {
	case c.Avg:
		out, err = grrrrrrr.Average(m)
	case c.Luma:
		out, err = grrrrrrr.Luma(m)
	case c.Desat:
		out, err = grrrrrrr.Desaturate(m)
	case c.DecomposeMin:
		out, err = grrrrrrr.DecomposeMin(m)
	case c.DecomposeMax:
		out, err = grrrrrrr.DecomposeMax(m)
	case c.SingleChannel:
		out, err = grrrrrrr.SingleChannel(m, c.GetColorChannel())
	case c.Shades: fallthrough
	
	default: 
		out, err = grrrrrrr.Shades(m, c.GetShadeCount())
	
	}
	if err != nil {
		log.Fatal(err.Error())
	}

	err = png.Encode(result, out)
	if err != nil {
		log.Fatalf("failed to save image at destination %q", err)
	}

	fmt.Printf("image converted successfully ( %s )\n", c.GetDest())
}

