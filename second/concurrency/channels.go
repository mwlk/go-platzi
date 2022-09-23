package main

import "fmt"

func main() {
	c := make(chan int, 3)

	c <- 1

	fmt.Println(<-c)

}
