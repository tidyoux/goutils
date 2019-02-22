package goutils

import (
	"os"
	"os/signal"

	log "github.com/sirupsen/logrus"
)

// RegisterSignalHandler registers a global system signal handler.
func RegisterSignalHandler(h func(os.Signal), sigs ...os.Signal) {
	Go("signal-handler", func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, sigs...)

		for {
			s := <-c
			log.Warnf("receive signal: %s", s)
			if h != nil {
				h(s)
			}
		}
	}, nil)
}
