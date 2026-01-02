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
	input  string
	output string
	args   []string
	root   string
)

func init() {
	flag.StringVar(&input, "i", "jpg", "input image format")
	flag.StringVar(&output, "o", "png", "output image format")
}

func main() {
	flag.Parse()
	args = flag.Args()
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "error: invalid argument\n")
		os.Exit(1)
	}
	root = args[0]
	err := filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return fmt.Errorf("%s: %w", path, errors.Unwrap(err))
			}
			if info.IsDir() {
				return nil
			}
			if filepath.Ext(path) == "."+input {
				imgconv.Convert(path, input, output)
				return nil
			}
			return fmt.Errorf("%s is not a valid file", path)
		})
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
	}
}
