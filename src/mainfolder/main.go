package main

import "fmt"

// ─────────────────────────────────────────────
// PACKAGE-LEVEL VARIABLE
// Declared outside any function — available to all files in this package.
// ─────────────────────────────────────────────
var globalMessage string

// ─────────────────────────────────────────────
// init()
// Runs automatically BEFORE main(), once per package.
// Used for one-time setup/initialisation.
// ─────────────────────────────────────────────
func init() {
	globalMessage = "Go Tutorial initialised!"
	fmt.Println("init() →", globalMessage)
}

// ─────────────────────────────────────────────
// main() — program entry point
// Call each demo function to run the tutorial sections.
// Comment/uncomment as needed while studying.
// ─────────────────────────────────────────────
func main() {
	fmt.Println("\n===== 01 · BASICS =====")
	BasicsDemo()

	fmt.Println("\n===== 02 · VARIABLES & TYPES =====")
	VariablesAndTypesDemo()

	fmt.Println("\n===== 03 · POINTERS =====")
	PointersDemo()

	fmt.Println("\n===== 04 · FUNCTIONS =====")
	FunctionsDemo()

	fmt.Println("\n===== 05 · CLOSURES & HIGHER-ORDER FUNCTIONS =====")
	ClosuresDemo()

	fmt.Println("\n===== 06 · ERROR HANDLING (defer / panic / recover) =====")
	ErrorHandlingDemo()
}
