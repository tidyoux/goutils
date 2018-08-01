package goutils

import (
	"fmt"
	"log"
	"runtime/debug"
)

// DeferRecover recover from panic.
func DeferRecover(tag string) {
	if err := recover(); err != nil {
		log.Printf("%s, recover from: %v\n%s\n", tag, err, debug.Stack())
	}
}

// WithRecover recover from panic.
func WithRecover(tag string, f func()) {
	defer DeferRecover(tag)
	f()
}

// Go is a wrapper of goruntine with recover.
func Go(name string, f func()) {
	go WithRecover(fmt.Sprintf("goroutine %s", name), f)
}
