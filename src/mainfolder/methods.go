package main

import (
	"fmt"
)

// 1. Simple function, multiple returns, and panic error handler
func divide(a, b int) (int, int) {
	if b == 0 {
		panic("division by zero")
	}
	return a / b, a % b
}

// 2. Passing address (pointer) vs value
func incrementByPointer(x *int) {
	*x = *x + 1
}

func incrementByValue(x int) int {
	x = x + 1
	return x
}

// 3. Variadic function
func sumAll(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func meth() {
	// 1. Multiple returns and panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	quotient, remainder := divide(10, 3)
	fmt.Println("Quotient:", quotient, "Remainder:", remainder)
	// Uncomment to see panic: divide(10, 0)

	// 2. Pointer vs value
	a := 5
	incrementByPointer(&a)
	fmt.Println("After incrementByPointer:", a)
	b := 5
	b = incrementByValue(b)
	fmt.Println("After incrementByValue:", b)

	// 3. Variadic function
	fmt.Println("Sum:", sumAll(1, 2, 3, 4, 5))
}
