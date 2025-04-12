package main

import (
	"fmt"
	"iter"
	"sync"
	"time"
)

func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

func testSomething(s string, e string) int {

	return 1
}

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) AllElements() iter.Seq[T] {
	// var elems []T
	// for e := lst.head; e != nil; e = e.next {
	// 	elems = append(elems, e.val)
	// }
	// return elems

	return func(yield func(T) bool) {
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

func testprint(msg string, message chan string) {
	for i := 0; i < 3; i++ {
		fmt.Println(msg, i)
		message <- msg
	}
}

func main_1() {
	// var s = []string{"foo", "bar", "zoo"}
	//
	// fmt.Println("index of zoo:", SlicesIndex(s, "zoo"))
	//
	// _ = SlicesIndex(s, "zoo")
	//
	// lst := List[int]{}
	// lst.Push(10)
	// lst.Push(13)
	// lst.Push(23)
	// for e := range lst.AllElements() {
	// 	fmt.Println(e)
	// }

	// message := make(chan string)
	// go testprint("foo", message)
	// go testprint("bar", message)
	// go testGoroutineOutput()
	// time.Sleep(time.Second)
	// fmt.Println("done")
	//
	// fmt.Println("=======Channel Example=======")
	// channelexample()
	// testTimeout()
	// channelClosing()
	// timerTest()

	// jobs := make(chan string, 5)
	// results := make(chan bool, 5)

	// for i := 0; i < 3; i++ {
	// 	go worker(i, jobs, results)
	// }

	// for i := 0; i < 5; i++ {
	// 	jobs <- string(i)
	// }
	// close(jobs)

	// for i := 0; i < 5; i++ {
	// 	<-results
	// }

	waitGroupWorker()
}

func testGoroutineOutput() string {
	return "test"
}

func channelexample() {

	messages := make(chan string)
	go func() { messages <- "ping" }()
	msg := <-messages
	fmt.Println(msg)
}

func testTimeout() {

	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}

func channelClosing() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	// This await is necessary to ensure that the goroutine has finished before the program exits
	// This is called as synchronization

	val, _ok := <-done
	fmt.Println("done", val, _ok)

	_, ok := <-jobs
	fmt.Println("received more jobs:", ok)
}

func timerTest() {
	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C
	fmt.Println("Timer 1 expired")
}

func worker(id int, jobs <-chan string, results chan<- bool) {
	for j := range jobs {
		fmt.Println("Started Executing  worker", id, j)
		time.Sleep(time.Second)
		fmt.Println("Executing worker", id, j)
		results <- true
	}
}

func workertest(i int) {
	fmt.Println("Executing worker", i)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", i)
}

func waitGroupWorker() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			fmt.Println("Called now for ", i)
			defer wg.Done()
			workertest(i)
		}()

		fmt.Println("test now for ", i)
	}
	wg.Wait()
}
