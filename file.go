package goutils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// FileExist determines whether a file exists
func FileExist(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

// WithOpenFile opens a file, do something, and close it.
func WithOpenFile(name string, flag int, perm os.FileMode, fun func(*os.File) error) error {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return err
	}

	defer f.Close()

	return fun(f)
}

// WithReadFile opens a file for read, and close it.
func WithReadFile(name string, fun func(*bufio.Reader) error) error {
	return WithOpenFile(name, os.O_RDONLY, 0, func(f *os.File) error {
		reader := bufio.NewReader(f)
		return fun(reader)
	})
}

// WithReadFileLineByLine opens a file, reads string line-by-line.
func WithReadFileLineByLine(name string, fun func(string) error) error {
	return WithReadFile(name, func(reader *bufio.Reader) error {
		return ForeachLine(reader, fun)
	})
}

// WithWriteFile opens a file for write, and close it.
func WithWriteFile(name string, fun func(*bufio.Writer) error) error {
	return WithOpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666, func(f *os.File) error {
		writer := bufio.NewWriter(f)
		err := fun(writer)
		if err != nil {
			return err
		}
		return writer.Flush()
	})
}

// CopyFile copys and transform data.
func CopyFile(from, to string, transformers ...Transformer) error {
	data, err := ioutil.ReadFile(from)
	if err != nil {
		return fmt.Errorf("read file %s failed, %v", from, err)
	}

	data, err = CombineTransformer(transformers).Transform(data)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(to, data, 0666)
	if err != nil {
		return fmt.Errorf("write file %s failed, %v", to, err)
	}

	return nil
}
