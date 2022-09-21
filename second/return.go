package main

import "fmt"

func sum(values ...int) int {
	total := 0
	for _, num := range values {
		total += num
	}

	return total
}

func print(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

// func getValues(x int) (int, int, int) {
// 	return 2 * x, 3 * x, 4 * x
// }

func getValues(x int) (double int, triple int, quad int) {
	double = 2 * x
	triple = 3 * x
	quad = 4 * x
	return
}

func main() {
	fmt.Println(sum(1))

	fmt.Println(3, 2, 1, 3)

	print("mirko", "ivo", "wlk")

	fmt.Println(getValues(2))
}
