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

func (s *SimpleScheduler) ConfigMasterWorkChan(work chan enginee.Request) {
	s.WorkChan = work
}

