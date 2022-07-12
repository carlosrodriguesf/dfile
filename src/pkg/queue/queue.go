package queue

import "sync"

type (
	Queue interface {
		Run(func())
		Wait()
	}

	queue struct {
		waitGroup *sync.WaitGroup
		cQueue    chan interface{}
	}
)

func New(limit int) Queue {
	return &queue{
		waitGroup: new(sync.WaitGroup),
		cQueue:    make(chan interface{}, limit),
	}
}

func (q *queue) Run(f func()) {
	q.waitGroup.Add(1)
	go func() {
		defer q.waitGroup.Done()

		q.cQueue <- nil

		f()

		<-q.cQueue
	}()
}

func (q *queue) Wait() {
	q.waitGroup.Wait()
}
