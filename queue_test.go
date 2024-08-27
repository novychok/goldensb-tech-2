package main

import (
	"fmt"
	"sync"
	"testing"
)

func Benchmark_queue(b *testing.B) {
	numMessages := 10000

	messages := make([]string, numMessages)
	for i := 0; i < numMessages; i++ {
		messages[i] = fmt.Sprintf("message%d", i+1)
	}

	for n := 0; n < b.N; n++ {
		wg := &sync.WaitGroup{}
		queue := newQueue()

		b.StartTimer()

		wg.Add(1)
		go queue.Produce(wg, messages)
		go queue.Consume(wg)
		wg.Wait()

		b.StopTimer()
	}
}
