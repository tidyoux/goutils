package goutils

import (
	"bufio"
	"fmt"
	"io"
)

// ReadBytes reads n bytes from reader.
func ReadBytes(reader io.Reader, n int) ([]byte, error) {
	if reader == nil || n <= 0 {
		return nil, fmt.Errorf("invalid args (%v, %v)", reader, n)
	}

	buf := make([]byte, n)
	pos := 0
	for {
		m, err := reader.Read(buf[pos:])
		if err != nil {
			if err == io.EOF {
				pos += m
				if pos < n {
					return buf[:pos], err
				}
				return buf, nil
			}
			return nil, err
		}

		pos += m
		if pos >= n {
			return buf, nil
		}
	}
}

// ForeachLine reads string from reader line-by-line.
func ForeachLine(reader io.Reader, fun func(line string) error) error {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		err := fun(scanner.Text())
		if err != nil {
			return err
		}
	}
	return scanner.Err()
}
