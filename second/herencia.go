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
	endDate string
}
func (ftEmployee FullTimeEmployee) getMessage() string {
	return "full time employee"
}


type TemporaryEmployee struct {
	Person
	Employee
	taxRate int
}

func (tEmpoyee TemporaryEmployee) getMessage() string {
	return "part time employee"
}

type PrintInfo interface{
	getMessage() string
}

func getMessage(p PrintInfo){
	fmt.Println(p.getMessage())
}
func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.age = 27
	ftEmployee.name = "mirko"

	fmt.Printf("%v\n", ftEmployee)

	tEmployee := TemporaryEmployee{}
	getMessage(tEmployee)

	getMessage(ftEmployee)
}
