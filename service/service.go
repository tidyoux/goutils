package service

import (
	"time"

	log "github.com/sirupsen/logrus"
)

type Service struct {
	worker   Worker
	interval time.Duration
	closeCh  chan struct{}
}

func New(w Worker) *Service {
	return NewWithInterval(w, time.Second*5)
}

func NewWithInterval(w Worker, interval time.Duration) *Service {
	return &Service{
		worker:   w,
		interval: interval,
		closeCh:  make(chan struct{}),
	}
}

func (s *Service) Start() {
	err := s.worker.Init()
	if err != nil {
		log.Errorf("worker (%s) init failed, %v", s.worker.Name(), err)
		return
	}

	log.Infof("worker (%s) started", s.worker.Name())

	for {
		select {
		case <-s.closeCh:
			s.worker.Destroy()
			log.Infof("worker (%s) stopped", s.worker.Name())
			break
		default:
			s.work()
		}

		time.Sleep(s.interval)
	}
}

func (s *Service) Stop() {
	close(s.closeCh)
}

func (s *Service) work() {
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("recover from worker (%s), %v", s.worker.Name(), err)
		}
	}()

	s.worker.Work()
}
