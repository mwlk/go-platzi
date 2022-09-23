package main

import (
	"fmt"
	"time"
)

func main() {
	channel1 := make(chan int)
	channel2 := make(chan int)

	d1 := 4 * time.Second
	d2 := 2 * time.Second

	go doSomething(d1, channel1, 10)
	go doSomething(d2, channel2, 20)

	for i := 0; i < 2; i++ {
		select {
		case channelMsg1 := <-channel1:
			fmt.Println(channelMsg1)
		case channelMsg2 := <-channel2:
			fmt.Println(channelMsg2)
		}
	}

}

func doSomething(i time.Duration, c chan<- int, param int) {
	time.Sleep(i)
	c <- param
}
