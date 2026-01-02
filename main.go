package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var input string
var output string
var args []string
var root string

func init() {
	flag.StringVar(&input, "i", "jpg", "input image format")
	flag.StringVar(&output, "o", "png", "output image format")
}

func main() {
	flag.Parse()
	args = flag.Args()
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "error: invalud argument\n")
		os.Exit(1)
	}
	root = args[0]
	err := filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return fmt.Errorf("error: %s: no such file or directory", root)
			}
			if info.IsDir() {
				return nil
			}
			if filepath.Ext(path) == "."+input {
				// convert .input -> .output
				return nil
			}
			return fmt.Errorf("error: %s is not a valid file", path)
		})
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err)
	}
}
