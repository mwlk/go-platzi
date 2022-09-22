package main

func Sum(x, y int) int {
	return x + y
}

func getMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
