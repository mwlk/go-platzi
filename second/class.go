package main

import "fmt"

type Employee struct {
	id   int
	name string
}

func main() {
	e := Employee{}
	fmt.Printf("%v", e)
	e.id = 0
	e.name = "test"
	fmt.Printf("%v", e)
}
