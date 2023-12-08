package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/sfomuseum/go-svg"
)

func main() {

	var format string

	flag.StringVar(&format, "format", "png", "...")

	flag.Parse()

	ctx := context.Background()

	for _, path := range flag.Args() {

		new_path := fmt.Sprintf("%s.png", path)

		r, err := os.Open(path)

		if err != nil {
			log.Fatalf("Failed to open %s for reading, %v", path, err)
		}

		defer r.Close()

		wr, err := os.OpenFile(new_path, os.O_RDWR|os.O_CREATE, 0644)

		if err != nil {
			log.Fatalf("Failed to create %s for writing, %v", new_path, err)
		}

		err = svg.RasterizeAsPNG(ctx, r, wr)

		if err != nil {
			log.Fatalf("Failed to rasterize %s, %v", path, err)
		}

		err = wr.Close()

		if err != nil {
			log.Fatalf("Failed to close %s after writing, %v", new_path, err)
		}
	}
}
