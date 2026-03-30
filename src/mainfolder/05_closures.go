package main

// ─────────────────────────────────────────────────────────────────────────────
// 05 · CLOSURES & HIGHER-ORDER FUNCTIONS
//
// Topics covered:
//   • Anonymous function (literal)
//   • IIFE — Immediately Invoked Function Expression
//   • Assigning a function to a variable
//   • Higher-order function — passing a function as argument
//   • Higher-order function — returning a function
//   • Closure — function that captures its surrounding scope
//   • Goroutine with anonymous function
// ─────────────────────────────────────────────────────────────────────────────

import (
	"fmt"
	"sync"
)

// ─────────────────────────────────────────────
// Higher-order helpers defined at package level
// ─────────────────────────────────────────────

// applyToAll calls fn on every element and returns a new slice.
// The parameter type func(int) int means "any function that takes
// an int and returns an int" — this is what makes it higher-order.
func applyToAll(nums []int, fn func(int) int) []int {
	result := make([]int, len(nums))
	for i, n := range nums {
		result[i] = fn(n)
	}
	return result
}

// makeMultiplier returns a closure that multiplies by factor.
// The returned function "closes over" factor — it remembers it
// even after makeMultiplier has returned.
func makeMultiplier(factor int) func(int) int {
	return func(n int) int {
		return n * factor
	}
}

// makeCounter returns a closure with its own private counter.
// Each call to the returned function increments and returns the count.
// Two separate counters are completely independent of each other.
func makeCounter() func() int {
	count := 0 // captured by the closure below
	return func() int {
		count++
		return count
	}
}

// makeAdder returns a closure that keeps a running total.
func makeAdder() func(int) int {
	total := 0
	return func(n int) int {
		total += n
		return total
	}
}

// ─────────────────────────────────────────────
// ClosuresDemo — run all closure examples
// ─────────────────────────────────────────────
func ClosuresDemo() {

	// ── 1. IIFE — Immediately Invoked Function Expression ──
	// Defined and called in the same statement.
	// Useful for one-off logic that doesn't need a name.
	func() {
		fmt.Println("IIFE: I run immediately, no name needed!")
	}()

	// IIFE with arguments and a return value
	greeting := func(name string) string {
		return "Hello, " + name + "!"
	}("Gopher")
	fmt.Println("IIFE with args:", greeting)

	// ── 2. Anonymous function assigned to a variable ───────
	// The variable holds the function — call it like any function.
	square := func(x int) int {
		return x * x
	}
	fmt.Println("square(9):", square(9))

	// ── 3. Passing a function as an argument (Higher-Order) ─
	nums := []int{1, 2, 3, 4, 5}

	doubled := applyToAll(nums, func(n int) int { return n * 2 })
	fmt.Println("doubled:", doubled)

	squared := applyToAll(nums, square) // reuse the variable from above
	fmt.Println("squared:", squared)

	// ── 4. Returning a function (Closure / Higher-Order) ────
	triple := makeMultiplier(3)
	tenX := makeMultiplier(10)
	fmt.Println("triple(7):", triple(7))
	fmt.Println("tenX(4):", tenX(4))

	// ── 5. Closure — capturing state ───────────────────────
	// Each counter has its own independent 'count' variable.
	counterA := makeCounter()
	counterB := makeCounter()
	fmt.Println("counterA:", counterA(), counterA(), counterA()) // 1 2 3
	fmt.Println("counterB:", counterB(), counterB())             // 1 2 (independent)

	// Running total with closure
	adder := makeAdder()
	fmt.Println("adder(10):", adder(10)) // 10
	fmt.Println("adder(5):", adder(5))   // 15 (remembers 10)
	fmt.Println("adder(20):", adder(20)) // 35 (remembers 15)

	// ── 6. Anonymous function in a goroutine ───────────────
	// Goroutines are lightweight concurrent threads.
	// Anonymous functions are perfect for quick goroutine tasks.
	var wg sync.WaitGroup
	messages := []string{"alpha", "beta", "gamma"}

	for _, msg := range messages {
		wg.Add(1)
		go func(m string) { // capture m by parameter, NOT the loop variable
			defer wg.Done()
			fmt.Println("goroutine received:", m)
		}(msg)
	}
	wg.Wait() // block until all goroutines finish
}
