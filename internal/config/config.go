package config

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/lulzshadowwalker/grrrrrrr/pkg/enum"
	"github.com/lulzshadowwalker/grrrrrrr/pkg/grrrrrrr"
)

type Method string

const (
	Avg Method = "average"
	Luma = "luma"
	Desat = "desaturate" 
	DecomposeMin = "decomposemin"
	DecomposeMax = "decomposemax" 
	SingleChannel = "singlechannel"
	Shades = "shades"
)

var methods = []Method{Avg, Luma, Desat, DecomposeMin, DecomposeMax, SingleChannel, Shades}

func (m Method) Validate() error {
	err := enum.IsIn[Method](m, methods)
	if err != nil {
		return fmt.Errorf("%s is not valid conversion method %w", string(m), err)
	}

	return nil
}

var (
	src string
	dest string

	method Method
	shadeCount int
	colorChannel grrrrrrr.ColorChannel
)

var workingDir string

func init() {
	flag.StringVar(&src, "src", "", "source image file to be converted")
	flag.StringVar(&dest, "dest", "", "output destination directory")

	m := flag.String("method", "shades", 
`conversion methods:
- average
- luma
- desaturate
- decomposeMin
- decomposeMax
- singleChannel
- shades
`)
	flag.IntVar(&shadeCount, "shadeCount", -1, "determines the number of shades when using the [shades] conversion method")
	cch := flag.String("colorChannel", "red", `determines which color channel to use for --method="singleChannle"`)

	flag.Parse()
	
	method = Method(strings.ToLower(*m))
	
	if c, err := grrrrrrr.StringToColorChannel(*cch); err != nil && method == SingleChannel { 
		log.Fatal(err.Error())
	} else {
		colorChannel = c
	}

	if src == "" {
		wd, err := getWorkingDir()
		if err != nil {
			log.Fatal(err.Error())
		}

		src = path.Join(wd, "../../assets/images/fallen-angels-1995.jpeg") 
	}

	if dest == "" {
		d, err := getWorkingDir()
		if err != nil {
			log.Fatal(err.Error())
		}
		
		dest = d 
	}
	dest = path.Join(dest, "grrrrrrr.png")

	err := method.Validate()
	if err != nil {
		log.Fatal(err.Error())
	}

	if shadeCount != -1 {
		if method != Shades {
			fmt.Println(`warning :: "--shadeCount" only has an effect when used with "--method=shades"`)
			return
		}

		if shadeCount <= 1 {
			log.Fatal(`"--shadeCount has to be greater than or equal to 2`)
		} else if shadeCount > 256 {
			log.Fatal(`"--shadeCount" can have a maximum value of 256`)
		}
	} else {
		shadeCount = 8
	}
}

func getWorkingDir() (string, error) {
	if workingDir != "" {
		return workingDir, nil
	}

	wd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf(
`failed to get the current working dir
try assigning "--src" and/or "--dest" explicitly
error: %w`, err)
	}
	
	workingDir = wd
	return wd, nil
}

func GetSrc() string {
	return src
}

func GetDest() string {
	return dest
}

func GetShadeCount() int {
	return shadeCount
}

func GetMethod() Method {
	return method
}

func GetColorChannel() grrrrrrr.ColorChannel {
	return colorChannel
}
