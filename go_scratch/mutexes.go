package main

import (
	"fmt"
	"math/rand"
	"rand"
	"sync"
)

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) updateValue(jobId string, value string) {
	fmt.Println("Job ", jobId, " waiting for lock")
	c.mu.Lock()
	fmt.Println("Job ", jobId, " got lock")
	defer c.mu.Unlock()
	c.counters[value]++
}

func main() {
	mutexes()

}

type read struct {
	key  int
	resp chan int
}

type write struct {
	key   int
	value int
	resp  chan bool
}

func statefulGoroutines() {

	var readOps uint64
	var writeOps uint64

	reads := make(chan read)
	writes := make(chan write)

	go func() {
		state := make(map[int]int)

		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.value
				write.resp <- true
			}
		}
	}()

	for r := 0; r < 100; r++ {
		go func() {

			for {
				curr_read := read{key: rand.Intn(5), resp: make(chan int)}
				reads <- curr

			}
		}()
	}
}

func mutexes() {
	c := &Container{
		counters: map[string]int{"a": 0, "b": 0},
	}
	var wg sync.WaitGroup

	doIncrement := func(jobId string, name string, n int) {
		for i := 0; i < n; i++ {
			c.updateValue(name, jobId)
		}
		wg.Done()
	}
	wg.Add(3)

	go doIncrement("a", "a", 50000)
	go doIncrement("b", "b", 50000)
	go doIncrement("c", "a", 30000)
	// go doIncrement("a", 30000)

	wg.Wait()
	fmt.Println(c.counters)
}
