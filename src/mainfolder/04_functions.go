package main

// ─────────────────────────────────────────────────────────────────────────────
// 04 · FUNCTIONS
//
// Topics covered:
//   • Simple function
//   • Multiple return values
//   • Named return values (naked return)
//   • Value parameters vs pointer parameters
//   • Variadic functions  (func f(args ...T))
//   • Function types  (type Fn func(...))
//   • Blank identifier _ to discard a return value
// ─────────────────────────────────────────────────────────────────────────────

import "fmt"

// ─────────────────────────────────────────────
// 1. Simple function
// ─────────────────────────────────────────────
func add(a, b int) int {
	return a + b
}

// ─────────────────────────────────────────────
// 2. Multiple return values
// Common Go pattern — return a result and an error together.
// ─────────────────────────────────────────────
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide %.2f by zero", a)
	}
	return a / b, nil
}

// ─────────────────────────────────────────────
// 3. Named return values (naked return)
// Name the return variables in the signature — useful for clarity.
// A bare "return" sends back the current values of named variables.
// ─────────────────────────────────────────────
func minMax(nums []int) (min, max int) {
	min, max = nums[0], nums[0]
	for _, n := range nums[1:] {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return // naked return — returns min and max
}

// ─────────────────────────────────────────────
// 4. Value parameter vs Pointer parameter
//
// byValue    — receives a copy; original is unchanged.
// byPointer  — receives an address; modifies the original.
// ─────────────────────────────────────────────
func applyTaxByValue(price float64, rate float64) float64 {
	price += price * rate // operates on a copy
	return price
}

func applyTaxByPointer(price *float64, rate float64) {
	*price += *price * rate // modifies the original
}

// ─────────────────────────────────────────────
// 5. Variadic function
// Use ...T to accept any number of arguments of type T.
// Inside the function, args is a slice of T.
// ─────────────────────────────────────────────
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// You can also spread a slice into a variadic call with slice...
func sumSlice() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("sum via spread:", sum(nums...)) // spread slice
}

// ─────────────────────────────────────────────
// 6. Function type
//
// Define a named type for a function signature.
// This makes code clearer and allows functions to be
// stored, passed, and returned like any other value.
// ─────────────────────────────────────────────
type MathOp func(int, int) int

func compute(a, b int, op MathOp) int {
	return op(a, b)
}

// ─────────────────────────────────────────────
// FunctionsDemo — run all function examples
// ─────────────────────────────────────────────
func FunctionsDemo() {

	// ── 1. Simple function ────────────────────
	fmt.Println("add(3, 4):", add(3, 4))

	// ── 2. Multiple return values ─────────────
	result, err := divide(10, 3)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Printf("10 / 3 = %.4f\n", result)
	}

	// Use _ to discard a return value
	result2, _ := divide(20, 4)
	fmt.Println("20 / 4 =", result2)

	// Division by zero
	_, err2 := divide(5, 0)
	fmt.Println("divide by 0 error:", err2)

	// ── 3. Named returns ──────────────────────
	nums := []int{3, 1, 7, 2, 9, 4}
	lo, hi := minMax(nums)
	fmt.Printf("minMax(%v) → min: %d, max: %d\n", nums, lo, hi)

	// ── 4. Value vs pointer params ────────────
	price := 100.0
	taxed := applyTaxByValue(price, 0.18) // copy — original unchanged
	fmt.Printf("byValue  — original: %.2f, taxed copy: %.2f\n", price, taxed)

	applyTaxByPointer(&price, 0.18) // pointer — original changes
	fmt.Printf("byPointer — original after call: %.2f\n", price)

	// ── 5. Variadic ───────────────────────────
	fmt.Println("sum(1,2,3):", sum(1, 2, 3))
	fmt.Println("sum(1..10):", sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	sumSlice()

	// ── 6. Function type ──────────────────────
	multiply := MathOp(func(x, y int) int { return x * y })
	subtract := MathOp(func(x, y int) int { return x - y })

	fmt.Println("compute(6, 3, multiply):", compute(6, 3, multiply))
	fmt.Println("compute(6, 3, subtract):", compute(6, 3, subtract))
}
