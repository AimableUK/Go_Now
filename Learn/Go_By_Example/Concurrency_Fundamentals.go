package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	go sayHello()

	time.Sleep(time.Second)

	ch := make(chan int)

	// === Channels ===

	go func() {
		ch <- 42
	}()

	val := <-ch
	fmt.Println(val)

	// === Buffered Channels ===
	ch1 := make(chan int, 3)

	ch1 <- 10
	ch1 <- 20
	ch1 <- 30

	goval := <-ch1

	fmt.Println(goval)

	// send-only
	// chan <- int

	// receive only
	// <-chan int

	// === Channel Syncroniation ===

	done := make(chan bool)

	go worker(done)
	<-done

	fmt.Println("Main finished", done)

	// === Range Over Channels ===
	chm := make(chan int)

	go func() {
		chm <- 10
		chm <- 20
		chm <- 30
		close(chm)
	}()
	for val := range chm {
		fmt.Println(val)
	}

	// === SELECT ===
	chA := make(chan string)
	chB := make(chan string)

	select {
	case msg := <-chA:
		fmt.Println(msg)

	case msg := <-chB:
		fmt.Println(msg)

	case <-time.After(2 * time.Second):
		fmt.Println("Timed out")
	}

	// === Non-Blocking Channel Operations ===
	select {
	case value := <-ch:
		fmt.Println("Received:", value)

	default:
		fmt.Println("Nothing available")
	}

	// === WaitGroups ===
	var wg sync.WaitGroup
	wg.Add(3)

	go work(1, &wg)
	go work(2, &wg)
	go work(3, &wg)

	wg.Wait()
	fmt.Println("All workers finished")

	// === MUTEX ===
	var counter int
	var mu sync.Mutex

	mu.Lock()
	defer mu.Unlock()

	counter++
}

func sayHello() {
	fmt.Println("Hello")
}

// ccan send but cant receive
func prod(ch chan<- int) {
	ch <- 100
}

// ===== Channel Syncroniation =====
func worker(done chan bool) {
	fmt.Println("Worker finished")
	done <- true
}

// ===== WaitGroups =====
func work(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("worker", id, "finished")
}
