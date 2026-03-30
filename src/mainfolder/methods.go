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

	AnanamousFunction()
	Higher1()
	fmt.Println(ReturnFunction())
	fmt.Println(ReturnFunction()("Go", "Language"))

	pos := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i))
	}
}

func AnanamousFunction() {

	// 1. Immediately Invoked Function Expression (IIFE)
	func() {
		fmt.Println("This runs immediately!")
	}()

	// 2. Anonymous function assigned to a variable
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println("Sum using anonymous function:", add(3, 4))

	// 3. Passing an anonymous function as an argument
	numbers := []int{1, 2, 3, 4, 5}
	processNumbers(numbers, func(n int) {
		fmt.Println("Processing number:", n)
	})

	// 4. Returning an anonymous function (closure)
	incrementer := makeIncrementer()
	fmt.Println("Incrementer:", incrementer()) // 1
	fmt.Println("Incrementer:", incrementer()) // 2
}

// Helper function for example 3
func processNumbers(nums []int, f func(int)) {
	for _, n := range nums {
		f(n)
	}
}

// Helper function for example 4
func makeIncrementer() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

//Higher order functions

func HigherOrderr1(a func(p, q string) string) {
	fmt.Println(a("Go", "Lang"))
}

func Higher1() {
	value := func(p, q string) string {
		return p + q + "Program"
	}
	HigherOrderr1(value)
}

//Return function

func ReturnFunction() func(i, j string) string {
	mf := func(i, j string) string {
		return i + j + "Program"
	}

	return mf
}

//closure function

func adder() func(int) int {
	sum := 0

	return func(x int) int {
		sum += x
		return sum
	}

}
