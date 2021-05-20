package factory

import (
	"geak/libs/log"
	"geak/libs/safe"
	"go.uber.org/zap"
)

type Result struct {
	Data interface{}
}

type Manager struct {
	WaitConsumeQueue chan interface{}
	Concurrency		 int
	ResultQueue 	 chan *Result
	Client			 Client
}
type Client interface {
	Consume(interface{})(*Result,error)
	Produce()(interface{},error)
}

func New(concurrency int,client Client)(*Manager) {
	return &Manager{
		WaitConsumeQueue: make(chan interface{}),
		Concurrency:      concurrency,
		ResultQueue:      make(chan *Result),
		Client:           client,
	}
}

func (this *Manager) Produce(){
	for {
		entity,err := this.Client.Produce()
		if err != nil {
			log.Error("produce error",zap.Error(err))
			continue
		}
		this.WaitConsumeQueue<-entity
	}
}

func (this *Manager) Consume(){
	pool := make(chan interface{}, this.Concurrency)
	for i := 0; i < this.Concurrency; i++ {
		pool <- i
	}
	for {
		entity := <-this.WaitConsumeQueue
		<-pool
		safe.Go(func() {
			var err error
			var result *Result
			if result,err = this.Client.Consume(entity);err != nil {
				log.Error("consume err",zap.Error(err))
			}
			this.ResultQueue <- result
			pool <- 1
		})
	}

}

