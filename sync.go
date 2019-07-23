package goutils

import (
	"runtime"
	"sync"
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

type BatchWork func() error

func BatchDo(count int, create func(int) BatchWork) error {
	works := MakeBatchWorks(count, create)
	return NewBatch(works).Do()
}

func MakeBatchWorks(count int, create func(int) BatchWork) []BatchWork {
	works := make([]BatchWork, 0, count)
	for i := 0; i < count; i++ {
		works = append(works, create(i))
	}

	return works
}

func batchGoroutineNumber() int {
	return runtime.NumCPU() * 2
}

type Batch struct {
	works []BatchWork

	goroutineNum    int
	maxBatchWorkNum int
}

func NewBatch(works []BatchWork) *Batch {
	var (
		workNum         = len(works)
		goroutineNum    = batchGoroutineNumber()
		maxBatchWorkNum int
	)

	if workNum <= goroutineNum {
		goroutineNum = workNum
		maxBatchWorkNum = 1
	} else {
		maxBatchWorkNum = workNum / goroutineNum
		goroutineNum = (workNum-1)/maxBatchWorkNum + 1
	}

	return &Batch{
		works:           works,
		goroutineNum:    goroutineNum,
		maxBatchWorkNum: maxBatchWorkNum,
	}
}

func (b *Batch) Do() error {
	if len(b.works) == 0 {
		return nil
	}

	var (
		errs = make([]error, b.goroutineNum)
		wg   sync.WaitGroup
	)
	for i := 0; i < b.goroutineNum; i++ {
		wg.Add(1)

		idx := i
		Go("Batch.do", func() {
			b.do(idx, &wg, errs)
		}, nil)
	}

	wg.Wait()

	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *Batch) do(idx int, wg *sync.WaitGroup, errs []error) {
	var err error
	defer func() {
		errs[idx] = err
		wg.Done()
	}()

	var (
		from = idx * b.maxBatchWorkNum
		to   = from + b.maxBatchWorkNum
	)
	if to > len(b.works) {
		to = len(b.works)
	}

	for i := from; i < to; i++ {
		err = b.works[i]()
		if err != nil {
			return
		}
	}
}
