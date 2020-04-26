package dirwalk

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type FileInfo struct {
	Name string
	Path string
}

func (i *FileInfo) FullPath() string {
	return filepath.Join(i.Path, i.Name)
}

type Filter func(path string, info os.FileInfo) bool

var defaultFilter = func(string, os.FileInfo) bool { return true }

// Walk gathers file recursively from path.
func Walk(path string, filter Filter) (files []*FileInfo, err error) {
	if filter == nil {
		filter = defaultFilter
	}

	path = strings.TrimRight(path, string(os.PathSeparator))
	items, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("read dir %s failed, %v", path, err)
	}

	for _, i := range items {
		if !filter(path, i) {
			continue
		}

		if i.IsDir() {
			fs, err := Walk(filepath.Join(path, i.Name()), filter)
			if err != nil {
				return nil, err
			}

			files = append(files, fs...)
		} else {
			files = append(files, &FileInfo{
				Name: i.Name(),
				Path: path,
			})
		}
	}
	return
}
