package goutils

import (
	"log"
	"time"
)

const (
	enabled = true
)

// DeferLogTimeCost logs time cost.
func DeferLogTimeCost(tag string) func() {
	start := time.Now()
	return func() {
		if enabled {
			log.Printf("%s time cost: %s", tag, time.Now().Sub(start))
		}
	}
}

// WithLogTimeCost logs time cost.
func WithLogTimeCost(tag string, f func()) {
	defer DeferLogTimeCost(tag)()
	f()
}
