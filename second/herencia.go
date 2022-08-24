package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

//declaracion anonima
type FullTimeEmployee struct {
	Person
	Employee
}

func GetMessage(p Person) {
	fmt.Println(p.name, ": edad =>", p.age)
}
func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.age = 27
	ftEmployee.name = "mirko"

	fmt.Printf("%v\n", ftEmployee)
}
