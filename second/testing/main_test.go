package main

import "testing"

// func TestSum(t *testing.T) {
// 	total := Sum(5, 5)

// 	if total != 10 {
// 		t.Errorf("sum was incorrect, got %d expected %d", total, 10)
// 	}
// }

func TestSum(t *testing.T) {
	tables := []struct {
		a int
		b int
		n int
	}{
		{1, 2, 3}, {2, 2, 4}, {25, 26, 51},
	}

	for _, item := range tables {
		total := Sum(item.a, item.b)

		if total != item.n {
			t.Errorf("Sum not valid, expected %d but result was %d", total, item.n)
		}
	}
}

func TestMax(t *testing.T) {
	tables := []struct {
		a int
		b int
		n int
	}{
		{4, 2, 4},
		{3, 2, 3},
		{2, 5, 5},
	}

	for _, item := range tables {
		max := getMax(item.a, item.b)

		if max != item.n {
			t.Errorf("max expected %d but result %d", max, item.n)
		}
	}
}

func TestFibonacci(t *testing.T) {
	table := []struct {
		a int
		n int
	}{
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}

	for _, item := range table {
		fib := fibonacci(item.a)

		if fib != item.n {
			t.Errorf("expected %d but result is %d", fib, item.n)
		}
	}
}
