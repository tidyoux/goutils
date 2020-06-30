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
	LineInRange(lineNum int) bool
	ParseTimestamp(data []byte) (time.Time, error)
	TimeInRange(t time.Time) bool
	ParseAll(lineNum int, timestamp time.Time, data []byte) (LogEntry, error)
	Filter
}

func Do(file io.Reader, p Processor, showTime bool) error {
	if showTime {
		defer goutils.DeferTimeCost(func(d time.Duration) {
			fmt.Printf("time: %v\n", d)
		})()
	}

	var (
		errBreak = errors.New("break")
		lineNum  int
	)

	err := goutils.ForeachLine(file, func(line string) error {
		lineNum++

		if p.Done(lineNum) {
			return errBreak
		}

		if !p.LineInRange(lineNum) {
			return nil
		}

		if len(line) == 0 {
			return nil
		}

		tm, err := p.ParseTimestamp([]byte(line))
		if err != nil {
			return fmt.Errorf("parse line %d timestamp failed, %v", lineNum, err)
		}

		if !p.TimeInRange(tm) {
			return nil
		}

		entry, err := p.ParseAll(lineNum, tm, []byte(line))
		if err != nil {
			return fmt.Errorf("parse line %d failed, %v", lineNum, err)
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
