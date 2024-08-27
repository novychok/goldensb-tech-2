package main

import (
	"fmt"
	"sync"
)

type broadcast interface {
	Produce(wg *sync.WaitGroup, messages []string)
	Consume(wg *sync.WaitGroup)
	Stop()
}

type queue struct {
	messages chan string
	done     chan struct{}
}

func (q *queue) Produce(wg *sync.WaitGroup, messages []string) {

	for _, msg := range messages {
		q.messages <- msg
	}

	q.Stop()
}

func (q *queue) Consume(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		msg, ok := <-q.messages
		if !ok {
			fmt.Println("Consumer: No more messages to consume, exiting.")
			return
		}
		fmt.Println("Consumed:", msg)
	}
}

func (q *queue) Stop() {
	close(q.done)
	close(q.messages)
}

func newQueue() broadcast {
	return &queue{
		messages: make(chan string),
		done:     make(chan struct{}),
	}
}
