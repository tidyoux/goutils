package goutils

import (
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
