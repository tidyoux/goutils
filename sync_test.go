package goutils

import (
	"testing"
)

func TestBatch(t *testing.T) {
	BatchDo(10, func(idx int) BatchWork {
		return func() error {
			t.Log("idx", idx)
			return nil
		}
	})
}
