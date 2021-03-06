package enginee

import (
//"log"
)
import (
	"carwer/model"
)

type Concurrency struct {
	Scheduler     Scheduler
	WorkCount     int
	ItemChan      chan model.Item
	WorkerProcess CrawerProcess
}

type CrawerProcess func(Request) (ParserResult, error)

type Scheduler interface {
	WorkNotify
	Submit(Request)
	CreateWorkChan() chan Request
	Scheduler()
}

type WorkNotify interface {
	WorkReady(chan Request)
}

func (c *Concurrency) Run(seds ...Request) {

	out := make(chan ParserResult)
	c.Scheduler.Scheduler()
	for i := 0; i < c.WorkCount; i++ {
		c.createWorker(c.Scheduler.CreateWorkChan(), out, c.Scheduler)
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
			//log.Printf("Get item: %v", item)
			go func() {
				c.ItemChan <- item
			}()
		}
	}

}

func (c *Concurrency) createWorker(in chan Request, out chan ParserResult, n WorkNotify) {
	go func() {
		for {
			// workers should pass their request chan
			n.WorkReady(in)
			request := <-in
			result, err := c.WorkerProcess(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
