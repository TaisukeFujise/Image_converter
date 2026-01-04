package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/TaisukeFujise/Image_converter/imgconv"
)

var (
	rawInput  string
	rawOutput string
	input     imgconv.Format
	output    imgconv.Format
	err       error
	root      string
)

func init() {
	flag.StringVar(&rawInput, "i", "jpg", "input image format")
	flag.StringVar(&rawOutput, "o", "png", "output image format")
}

func main() {
	flag.Parse()
	root = parseArgs(flag.Args())
	// TODO
	input, err = imgconv.ParseFormat(rawInput)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
	output, err = imgconv.ParseFormat(rawOutput)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
	err := filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return fmt.Errorf("%s: %w", path, errors.Unwrap(err))
			}
			if info.IsDir() {
				return nil
			}
			if input.Match(path) {
				err = imgconv.Convert(path, output)
				if err != nil {
					return err
				}
				return nil
			}
			return fmt.Errorf("%s is not a valid file", path)
		})
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
	}
}

func parseArgs(args []string) string {
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "error: invalid argument\n")
		os.Exit(1)
	}
	return args[0]
}
