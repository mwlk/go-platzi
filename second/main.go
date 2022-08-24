// package main

// import (
// 	"fmt"
// 	"strconv"
// 	"time"
// )

// func main() {
// 	var x int
// 	x = 8
// 	y := 7

// 	fmt.Println(x, y)
// 	value, err := strconv.ParseInt("lala", 0, 64)

// 	if err != nil {
// 		fmt.Printf("%v\n", err)
// 	} else {
// 		fmt.Println(value)
// 	}

// 	m := make(map[string]int)

// 	m["key"] = 6
// 	fmt.Println(m["key"])

// 	s := []int{1, 2, 3}
// 	s = append(s, 16)
// 	for index, value := range s {
// 		fmt.Println(index, value)
// 	}

// 	// c := make(chan int)
// 	// go doSomething(c)
// 	// <-c

// 	g := 25
// 	h := &g
// 	fmt.Println(g, h, *h)
// }

// func doSomething(c chan int) {
// 	time.Sleep(3 * time.Second)

// 	fmt.Println("done")
// 	c <- 1
// }
