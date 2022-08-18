package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/nfnt/resize"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/spf13/pflag"
)

// TODO: replace all the `log.Fatal`s and split into functions

const defaultSuffix = ".preview.jpg"

var (
	flagOutfile = pflag.StringP("output", "o", "", fmt.Sprintf("Output filename. Default: input file + \"%s\"", defaultSuffix))
	flagPreview = pflag.BoolP("preview", "p", false, "Extract preview instead of thumbnail (preview is a larger, more detailed image, if present)")
	flagResize  = pflag.StringP("resize", "r", "", "Resize the output image. The format is [W[x[H]]] where W is width, H is height. Empty string: no resize. 'W' or 'Wx' resize width to W, maintaining the aspect ratio. WxH changes both width and height")
)

// parseResize parses a resize string of the format [W[x[H]]] where W and H are
// non-negative integers. It returns the parsed width, height, and an error if
// any. Examples:
// "" -> W=0, H=0
// "128" or "128x" -> W=128, H=0
// "128x128" -> W=128, H=128
// "-1x-1" returns an error
//
// An error is returned if the string has an invalid format or if the numbers
// are negative.
func parseResize(r string) (int, int, error) {
	if r == "" {
		return 0, 0, nil
	}
	parts := strings.Split(r, "x")
	switch len(parts) {
	case 0:
		return 0, 0, nil
	case 1:
		w, err := strconv.Atoi(parts[0])
		if err != nil {
			return 0, 0, fmt.Errorf("invalid width '%s'", parts[0])
		}
		if w < 0 {
			return 0, 0, fmt.Errorf("width must be non-negative")
		}
		return w, 0, nil
	case 2:
		w, err := strconv.Atoi(parts[0])
		if err != nil {
			return 0, 0, fmt.Errorf("invalid width '%s'", parts[0])
		}
		var h int
		if parts[1] == "" {
			h = 0
		} else {
			h, err = strconv.Atoi(parts[1])
			if err != nil {
				return 0, 0, fmt.Errorf("invalid height '%s'", parts[1])
			}
		}
		if w < 0 {
			return 0, 0, fmt.Errorf("width must be non-negative")
		}
		if h < 0 {
			return 0, 0, fmt.Errorf("height must be non-negative")
		}
		return w, h, nil
	default:
		return 0, 0, fmt.Errorf("invalid resize string '%s'", r)
	}
}

func main() {
	pflag.Parse()
	if len(pflag.Args()) < 1 {
		log.Fatalf("missing file name")
	}
	filename := pflag.Arg(0)
	fd, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	x, err := exif.Decode(fd)
	if err != nil {
		log.Fatal(err)
	}
	var im []byte
	if *flagPreview {
		log.Fatal("PreviewImage not supported yet")
		// TODO add support for PreviewImage TIFF tag (extension)
		//tag, err := x.Get(exif.Preview)
	} else {
		im, err = x.JpegThumbnail()
	}
	if err != nil {
		log.Fatal(err)
	}

	w, h, err := parseResize(*flagResize)
	if err != nil {
		log.Fatal(err)
	}
	outfile := filename + defaultSuffix
	if *flagOutfile != "" {
		outfile = *flagOutfile
	}
	if w > 0 || h > 0 {
		log.Printf("Resizing to %dx%d", w, h)
		img, _, err := image.Decode(bytes.NewReader(im))
		if err != nil {
			log.Fatalf("Failed to decode image: %v", err)
		}
		resizedImg := resize.Resize(uint(w), uint(h), img, resize.Lanczos3)
		wfd, err := os.Create(outfile)
		if err != nil {
			log.Fatalf("os.Create failed: %v", err)
		}
		defer wfd.Close()
		if err := jpeg.Encode(wfd, resizedImg, nil); err != nil {
			log.Fatalf("Failed to encode image: %v", err)
		}
	} else {
		if err := os.WriteFile(outfile, im, 0644); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println(outfile)
}
