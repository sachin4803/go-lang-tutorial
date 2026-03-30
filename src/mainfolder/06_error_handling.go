package main

// ─────────────────────────────────────────────────────────────────────────────
// 06 · ERROR HANDLING — defer / panic / recover
//
// Topics covered:
//   • The error interface — idiomatic Go error handling
//   • fmt.Errorf — creating formatted errors
//   • errors.New — creating simple errors
//   • defer — schedule cleanup before a function exits
//   • Multiple defers — LIFO (stack) execution order
//   • panic — abort normal execution
//   • recover — catch a panic in a deferred function
// ─────────────────────────────────────────────────────────────────────────────

import (
	"errors"
	"fmt"
)

// ─────────────────────────────────────────────
// SENTINEL ERROR
// A package-level error value others can compare against.
// ─────────────────────────────────────────────
var ErrDivisionByZero = errors.New("division by zero")

// ─────────────────────────────────────────────
// safeDivide — idiomatic error return (no panic)
// Returns (result, nil) on success or (0, err) on failure.
// ─────────────────────────────────────────────
func safeDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivisionByZero // return sentinel error
	}
	return a / b, nil
}

// ─────────────────────────────────────────────
// openFile — simulates a function that must clean up a resource.
// defer guarantees the "close" runs even if an error occurs.
// ─────────────────────────────────────────────
func openFile(name string) {
	fmt.Println("opening file:", name)
	defer fmt.Println("closing file:", name) // runs when openFile returns

	fmt.Println("reading file:", name)
	// ... file processing here ...
}

// ─────────────────────────────────────────────
// deferOrder — shows that multiple defers execute in LIFO order.
// Think of it like a stack: last-in, first-out.
// ─────────────────────────────────────────────
func deferOrder() {
	fmt.Println("deferOrder: start")
	defer fmt.Println("deferred 1 — runs LAST")
	defer fmt.Println("deferred 2 — runs SECOND")
	defer fmt.Println("deferred 3 — runs FIRST")
	fmt.Println("deferOrder: end (defers fire now)")
}

// ─────────────────────────────────────────────
// riskyOperation — demonstrates panic + recover.
//
// panic: immediately stops the current function and begins
//
//	unwinding the call stack, running deferred functions.
//
// recover: called inside a deferred function to stop the unwind
//
//	and return the panic value — letting the program continue.
//
// ─────────────────────────────────────────────
func riskyOperation(input int) {
	// The deferred recovery MUST be set up BEFORE the panic.
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic:", r)
		}
	}()

	fmt.Println("riskyOperation: input =", input)
	if input < 0 {
		panic(fmt.Sprintf("negative input is not allowed: %d", input))
	}
	fmt.Println("riskyOperation: completed normally")
}

// ─────────────────────────────────────────────
// safeExec — generic wrapper: runs any function safely,
// catching any panic and returning it as an error.
// ─────────────────────────────────────────────
func safeExec(fn func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("caught panic: %v", r)
		}
	}()
	fn()
	return nil
}

// ─────────────────────────────────────────────
// ErrorHandlingDemo — run all error handling examples
// ─────────────────────────────────────────────
func ErrorHandlingDemo() {

	// ── 1. Idiomatic error handling ───────────
	result, err := safeDivide(10, 3)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("10 / 3 =", result)
	}

	// Compare against sentinel error
	_, err = safeDivide(5, 0)
	if errors.Is(err, ErrDivisionByZero) {
		fmt.Println("caught sentinel error:", err)
	}

	// ── 2. defer — guaranteed cleanup ─────────
	openFile("data.txt")

	// ── 3. Multiple defers — LIFO order ───────
	deferOrder()

	// ── 4. panic + recover ────────────────────
	riskyOperation(42) // normal execution
	riskyOperation(-1) // triggers panic → recovered → program continues

	// ── 5. Generic safe wrapper ───────────────
	err = safeExec(func() {
		fmt.Println("safeExec: running stable function")
	})
	fmt.Println("safeExec err (nil expected):", err)

	err = safeExec(func() {
		panic("something exploded inside safeExec!")
	})
	fmt.Println("safeExec err (caught):", err)
}
