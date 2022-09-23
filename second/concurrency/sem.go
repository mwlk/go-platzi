package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	c := make(chan int, 5)

	for i := 0; i < 10; i++ {
		c <- 1
		wg.Add(1)

		go doSomething(i, &wg, c)
	}
	wg.Wait()
}

func doSomething(i int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	fmt.Println("Id %d started", i)
	time.Sleep(5 * time.Second)

	fmt.Println("end %d", i)
	<-c
}
