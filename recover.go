package goutils

import (
	"fmt"
	"runtime/debug"

	log "github.com/sirupsen/logrus"
)

// DeferRecover defer recover from panic.
func DeferRecover(tag string, handlePanic func(error)) func() {
	return func() {
		if err := recover(); err != nil {
			log.Errorf("%s, recover from: %v\n%s\n", tag, err, debug.Stack())
			if handlePanic != nil {
				handlePanic(fmt.Errorf("%v", err))
			}
		}
	}
}

// WithRecover recover from panic.
func WithRecover(tag string, f func(), handlePanic func(error)) {
	defer DeferRecover(tag, handlePanic)()

	f()
}

// Go is a wrapper of goruntine with recover.
func Go(name string, f func(), handlePanic func(error)) {
	go WithRecover(fmt.Sprintf("goroutine %s", name), f, handlePanic)
}
