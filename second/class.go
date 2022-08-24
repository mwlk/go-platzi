// package main

// import "fmt"

// type Employee struct {
// 	id       int
// 	name     string
// 	vacation bool
// }

// func (e *Employee) SetId(id int) {
// 	e.id = id
// }
// func (e *Employee) SetName(name string) {
// 	e.name = name
// }

// func (e *Employee) GetId() int {
// 	return e.id
// }
// func (e *Employee) GetName() string {
// 	return e.name
// }

// func constructor(id int, name string, vacation bool) *Employee {
// 	return &Employee{
// 		id:       id,
// 		name:     name,
// 		vacation: vacation,
// 	}
// }

// func main() {
// 	e := Employee{}
// 	fmt.Printf("%v", e)
// 	e.SetId(5)
// 	e.SetName("set")
// 	fmt.Println(e.GetId(), e.GetName())

// 	e2 := constructor(1, "ctor", true)
// 	fmt.Println(e2)
// }
