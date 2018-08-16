package goutils

import (
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	enabled = true
)

// DeferLogTimeCost logs time cost.
func DeferLogTimeCost(tag string) func() {
	start := time.Now()
	return func() {
		if enabled {
			log.Infof("%s time cost: %s", tag, time.Now().Sub(start))
		}
	}
}

// WithLogTimeCost logs time cost.
func WithLogTimeCost(tag string, f func()) {
	defer DeferLogTimeCost(tag)()
	f()
}
