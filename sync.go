package goutils

import (
	"sync/atomic"
)

type AtomicBool int32

func (b *AtomicBool) Set(is bool) {
	if is {
		atomic.StoreInt32((*int32)(b), 1)
	} else {
		atomic.StoreInt32((*int32)(b), 0)
	}
}

func (b *AtomicBool) Is() bool {
	return atomic.LoadInt32((*int32)(b)) != 0
}
