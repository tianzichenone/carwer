package scheduler

import "carwer/enginee"

type QueuedScheduler struct {
	RequestChan chan enginee.Request
	WorkChan    chan chan enginee.Request
}

func (q *QueuedScheduler) Submit(req enginee.Request) {
	q.RequestChan <- req
}

func (q *QueuedScheduler) WorkReady(work chan enginee.Request) {
	q.WorkChan <- work
}

func (s *QueuedScheduler) ConfigMasterWorkChan(work chan enginee.Request) {
	panic("Not implement")
}

func (q *QueuedScheduler) Scheduler() {
	go func() {
		q.RequestChan = make(chan enginee.Request)
		q.WorkChan = make(chan chan enginee.Request)
		requestQueue := []enginee.Request{}
		workChanQueue := [] chan enginee.Request{}
		for {
			var receiveRequest enginee.Request
			var readyWorkChan chan enginee.Request
			if len(requestQueue) > 0 && len(workChanQueue) > 0 {
				receiveRequest = requestQueue[0]
				readyWorkChan = workChanQueue[0]
			}
			select {
			case request := <-q.RequestChan:
				requestQueue = append(requestQueue, request)
			case workChan := <-q.WorkChan:
				workChanQueue = append(workChanQueue, workChan)
			case readyWorkChan <- receiveRequest:
				requestQueue = requestQueue[1:]
				workChanQueue = workChanQueue[1:]
			}
		}
	}()
}

