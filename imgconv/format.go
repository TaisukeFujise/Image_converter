package imgconv

import (
	"fmt"
	"path/filepath"
	"strings"
)

type Format string

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
		err = fmt.Errorf("%s is not a valid format2", raw)
		return
	}
	err = nil
	return
}

// pathを受け取って、そのファイルの拡張子がinputで指定された拡張子と一致しているかを確認して、boolを返す
func (f Format) Match(path string) bool {
	formatedPath, err := ParseFormat(filepath.Ext(path)[1:])
	if err != nil {
		return false
	}
	if formatedPath != f {
		return false
	}
	return true
}
