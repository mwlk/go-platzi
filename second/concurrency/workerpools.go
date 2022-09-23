package main

import "fmt"

func main() {
	tasks := []int{2, 3, 4, 5, 7, 10, 12, 40}

	nWorkers := 3

	jobs := make(chan int, len(tasks))
	result := make(chan int, len(tasks))

	for i := 0; i < nWorkers; i++ {
		go Worker(i, jobs, result)
	}

	for _, value := range tasks {
		jobs <- value
	}
	close(jobs)

	for r := 0; r < len(tasks); r++ {
		<-result
	}
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func Worker(id int, jobs <-chan int, result chan<- int) {
	for job := range jobs {
		fmt.Println("Worker id ", id)
		fmt.Println("start fib with => ", job)

		fib := Fibonacci(job)
		fmt.Println("Worker %d, job %d, fib %d", id, job, fib)

		result <- fib
	}
}
