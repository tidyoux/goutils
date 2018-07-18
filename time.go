package goutils

import (
	"log"
	"time"
)

const (
	enabled = true
)

// LogTimeCost def
func LogTimeCost(tag string) func() {
	start := time.Now()
	return func() {
		if enabled {
			log.Printf("%s time cost: %s", tag, time.Now().Sub(start))
		}
	}
}

// WithLogTimeCost def
func WithLogTimeCost(tag string, f func()) {
	defer LogTimeCost(tag)()
	f()
}
