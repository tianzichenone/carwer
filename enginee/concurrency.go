package enginee

import (
	"log"
)

type Concurrency struct {
	Scheduler Scheduler
	WorkCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigMasterWorkChan(chan Request)
	WorkReady(chan Request)
	Scheduler()
}

func (c Concurrency) Run(seds ...Request) {

	out := make(chan ParserResult)
	c.Scheduler.Scheduler()
	for i := 0; i < c.WorkCount; i ++ {
		c.createWorker(out, c.Scheduler)
	}

	for _, req := range seds {
		c.Scheduler.Submit(req)
	}

	for {
		result := <-out
		for _, req := range result.Requests {
			c.Scheduler.Submit(req)
		}
		for _, item := range result.Items {
			log.Printf("Get item: %v", item)
		}
	}

}

func (c Concurrency) createWorker(out chan ParserResult, s Scheduler) {
	in := make(chan Request)
	go func() {
		for {
			// workers should pass their request chan
			s.WorkReady(in)
			request := <- in
			result, err := Work(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}