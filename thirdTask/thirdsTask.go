package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type WebVisits struct {
	visits sync.Map
	mu     sync.Mutex
}

func (w *WebVisits) Increment(url string, i int64) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.visits.Store(url, i)
}

func (w *WebVisits) GetVisit(url string) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.visits.Range(func(key, value any) bool {
		if key == url {
			fmt.Println(key, value)
		}
		return true
	})
}

func main() {
	var wg sync.WaitGroup
	var counter int64

	w := WebVisits{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}
	wg.Wait()
	w.Increment("ssss", counter)
	go w.GetVisit("ssss")
	time.Sleep(100 * time.Millisecond)

}
