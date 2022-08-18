package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/spf13/pflag"
)

const defaultSuffix = ".preview.jpg"

var (
	flagOutfile   = pflag.StringP("output", "o", "", fmt.Sprintf("Output filename. Default: input file + \"%s\"", defaultSuffix))
	flagThumbnail = pflag.BoolP("thumbnail", "t", false, "Extract thumbnail instead of preview (preview is a larger, more detailed image, if present)")
)

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
	if !*flagThumbnail {
		log.Fatal("PreviewImage not supported yet")
		// TODO add support for PreviewImage TIFF tag (extension)
		//tag, err := x.Get(exif.Preview)
	} else {
		im, err = x.JpegThumbnail()
	}
	if err != nil {
		log.Fatal(err)
	}
	outfile := filename + defaultSuffix
	if *flagOutfile != "" {
		outfile = *flagOutfile
	}
	if err := os.WriteFile(outfile, im, 0644); err != nil {
		log.Fatal(err)
	}
	fmt.Println(outfile)
}
