package main

// ─────────────────────────────────────────────────────────────────────────────
// 03 · POINTERS
//
// Topics covered:
//   • What a pointer is (stores a memory address, not a value)
//   • Declaring a pointer  (*Type)
//   • & — address-of operator (get the address of a variable)
//   • * — dereference operator (get/set the value at an address)
//   • Nil pointer (zero value of a pointer)
//   • Pointers to different types
//   • Why use pointers — mutating values in functions
// ─────────────────────────────────────────────────────────────────────────────

import "fmt"

// ─────────────────────────────────────────────
// doubleByValue — receives a COPY of the value.
// The original variable is NOT changed.
// ─────────────────────────────────────────────
func doubleByValue(n int) int {
	n = n * 2
	return n
}

// ─────────────────────────────────────────────
// doubleByPointer — receives the ADDRESS of the variable.
// It modifies the original through the pointer.
// ─────────────────────────────────────────────
func doubleByPointer(n *int) {
	*n = *n * 2 // dereference n, then multiply
}

// ─────────────────────────────────────────────
// setName — modifies a string through a pointer
// ─────────────────────────────────────────────
func setName(s *string, newName string) {
	*s = newName
}

// ─────────────────────────────────────────────
// PointersDemo — run all pointer examples
// ─────────────────────────────────────────────
func PointersDemo() {

	// ── 1. Declare a pointer ──────────────────
	var x int = 10
	var p *int = &x // p holds the memory address of x

	fmt.Println("x:", x)
	fmt.Println("&x (address of x):", &x) // e.g. 0xc0000140a8
	fmt.Println("p (pointer value):", p)  // same address
	fmt.Println("*p (dereference):", *p)  // 10

	// ── 2. Modify through a pointer ───────────
	*p = 99                            // write 99 to the memory address p points to
	fmt.Println("x after *p = 99:", x) // x is now 99

	// ── 3. Nil pointer ────────────────────────
	// The zero value of any pointer type is nil.
	// Dereferencing a nil pointer causes a runtime panic — always guard!
	var nilPtr *int
	fmt.Println("nil pointer:", nilPtr) // <nil>
	if nilPtr == nil {
		fmt.Println("nilPtr is nil — safe to skip dereference")
	}

	// ── 4. Passing by value vs by pointer ─────
	num := 5
	doubled := doubleByValue(num)
	fmt.Println("byValue  — original:", num, "returned:", doubled) // original unchanged

	doubleByPointer(&num)
	fmt.Println("byPointer — original after call:", num) // original IS changed

	// ── 5. Pointer to float64 ─────────────────
	price := 9.99
	pPrice := &price
	fmt.Printf("price: %.2f, address: %p, via ptr: %.2f\n", price, pPrice, *pPrice)

	// ── 6. Pointer to string ──────────────────
	name := "Alice"
	setName(&name, "Bob")
	fmt.Println("name after setName:", name)

	// ── 7. Indexing a string through a pointer ─
	// You must DEREFERENCE the pointer before indexing.
	city := "Golang"
	sPtr := &city
	fmt.Printf("First char via *sPtr: %c\n", (*sPtr)[0]) // NOT sPtr[0]
}
