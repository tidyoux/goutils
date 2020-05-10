package logparser

import (
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/tidyoux/goutils"
)

type LogEntry fmt.Stringer

type Processor interface {
	Done(lineNum int) bool
	InRange(lineNum int) bool
	Parse(lineNum int, data []byte) (LogEntry, error)
	Accept(entry LogEntry) bool
}

func Do(file io.Reader, p Processor, showTime bool) error {
	if showTime {
		defer goutils.DeferTimeCost(func(d time.Duration) {
			fmt.Printf("time: %v\n", d)
		})()
	}

	var (
		errBreak = errors.New("break")
		n        int
	)
	err := goutils.ForeachLine(file, func(line string) error {
		n++

		if p.Done(n) {
			return errBreak
		}

		if !p.InRange(n) {
			return nil
		}

		if len(line) == 0 {
			return nil
		}

		entry, err := p.Parse(n, []byte(line))
		if err != nil {
			return fmt.Errorf("parse line %d failed, %v", n, err)
		}

		if p.Accept(entry) {
			fmt.Println(entry)
		}
		return nil
	})
	if err != nil && err != errBreak {
		return err
	}
	return nil
}
