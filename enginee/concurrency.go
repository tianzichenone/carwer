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
}

func (c Concurrency) Run(seds ...Request) {

	in := make(chan Request)
	out := make(chan ParserResult)
	c.Scheduler.ConfigMasterWorkChan(in)

	for i := 0; i < c.WorkCount; i ++ {
		c.createWorker(in, out)
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

func (c Concurrency) createWorker(in chan Request, out chan ParserResult) {
	go func() {
		for {
			request := <- in
			result, err := Work(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}