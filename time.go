package goutils

import (
	"fmt"
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

var (
	// ErrTimeout returns when timeout.
	ErrTimeout = fmt.Errorf("timeout")
)

// WithTimeout runs f with a timeout.
func WithTimeout(duration time.Duration, f func() error) error {
	var (
		c   = make(chan struct{})
		err error
	)

	Go("WithTimeout", func() {
		err = f()
		close(c)
	}, func(panicErr error) {
		err = panicErr
		close(c)
	})

	select {
	case <-c:
		return err
	case <-time.After(duration):
		return ErrTimeout
	}
}

// UpdateMonitor monitors for update interval.
type UpdateMonitor struct {
	maxInterval    time.Duration
	value          int64
	lastUpdateTime time.Time
	handleTimeout  func(int64, time.Duration)
}

// NewUpdateMonitor returns a new UpdateMonitor.
func NewUpdateMonitor(maxInterval time.Duration, handleTimeout func(int64, time.Duration)) *UpdateMonitor {
	return &UpdateMonitor{
		maxInterval:    maxInterval,
		value:          0,
		lastUpdateTime: time.Now(),
		handleTimeout:  handleTimeout,
	}
}

// Update updates monitor state.
func (m *UpdateMonitor) Update(value int64) {
	if value != m.value {
		m.value = value
		m.lastUpdateTime = time.Now()
	} else {
		interval := time.Now().Sub(m.lastUpdateTime)
		if interval > m.maxInterval {
			if m.handleTimeout != nil {
				m.handleTimeout(m.value, interval)
			}
		}
	}
}

// Value returns the value of UpdateMoniter.
func (m *UpdateMonitor) Value() int64 {
	return m.value
}
