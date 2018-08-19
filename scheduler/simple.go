package scheduler

import (
	"carwer/enginee"
)

type SimpleScheduler struct {
	WorkChan chan enginee.Request
}

func (s *SimpleScheduler) Submit(req enginee.Request) {
	go func() {
		s.WorkChan <- req
	}()
}

func (s *SimpleScheduler) CreateWorkChan() chan enginee.Request {
	return s.WorkChan
}

func (s *SimpleScheduler) Scheduler() {
	s.WorkChan = make(chan enginee.Request)
}

func (s *SimpleScheduler) WorkReady(chan enginee.Request) {

}
