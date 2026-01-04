package imgconv

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type Format string

// TODO
func ParseFormat(raw string) (result Format, err error) {
	switch strings.ToLower(raw) {
	case "jpg":
		result = ".jpg"
	case "jpeg":
		result = ".jpg"
	case "png":
		result = ".png"
	case "gif":
		result = ".gif"
	default:
		result = ""
		err = fmt.Errorf("%s is not a valid format", raw)
		return
	}
	err = nil
	return
}

// pathを受け取って、そのファイルの拡張子がinputで指定された拡張子と一致しているかを確認して、boolを返す
func (f Format) Match(path string) bool {
	if filepath.Ext(path) != string(f) {
		return false
	}
	return true
}

func Convert(path string, output Format) error {
	reader, err := os.Open(path)
	if err != nil {
		return err
	}
	defer reader.Close()
	writer, err := os.Create(filepath.Base(path) + string(output))
	if err != nil {
		return err
	}
	defer writer.Close()
	switch output {
	case ".jpg":
		err = convertToJPEG(reader, writer)
	case ".png":
		err = convertToPNG(reader, writer)
	case ".gif":
		err = convertToGIF(reader, writer)
	}
	if err != nil {
		return err
	}
	return nil
}

func convertToJPEG(r io.Reader, w io.Writer) error {
	img, _, err := image.Decode(r)
	if err != nil {
		return err
	}
	return jpeg.Encode(w, img, nil)
}

func convertToPNG(r io.Reader, w io.Writer) error {
	img, _, err := image.Decode(r)
	if err != nil {
		return err
	}
	return png.Encode(w, img)
}

func convertToGIF(r io.Reader, w io.Writer) error {
	img, _, err := image.Decode(r)
	if err != nil {
		return err
	}
	return gif.Encode(w, img, nil)
}
