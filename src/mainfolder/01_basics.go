package main

// ─────────────────────────────────────────────────────────────────────────────
// 01 · BASICS
//
// Topics covered:
//   • What a package is and the role of package main
//   • Exported vs unexported identifiers
//   • Multiple files in the same package
//   • Using an external package (packagefolder)
// ─────────────────────────────────────────────────────────────────────────────

import (
	"fmt"

	"github.com/sachin4803/go-lang-tutorial/src/packagefolder"
)

// ─────────────────────────────────────────────
// EXPORTED vs UNEXPORTED
//
// In Go, capitalised names are EXPORTED (visible outside the package).
// Lowercase names are UNEXPORTED (private to the package).
// ─────────────────────────────────────────────

// Greet is exported — callable from other packages.
func Greet(name string) string {
	return "Hello, " + name + "!"
}

// greetInternal is unexported — only usable within this package.
func greetInternal() string {
	return "Hello from inside the package!"
}

// ─────────────────────────────────────────────
// BasicsDemo — run all basics examples
// ─────────────────────────────────────────────
func BasicsDemo() {
	// 1. Simple output
	fmt.Println("Hello, World!")

	// 2. Exported function
	fmt.Println(Greet("Gopher"))

	// 3. Unexported function (works fine inside the same package)
	fmt.Println(greetInternal())

	// 4. Using a function from an external package.
	//    Only exported names (capital letters) can be accessed here.
	packagefolder.Hellofrompackage()

	// 5. Using a package constant defined in packagefolder
	fmt.Println("Package version:", packagefolder.Version)

	// 6. Using a package utility function
	fmt.Println("Square of 7:", packagefolder.Square(7))
	fmt.Println("Max(12, 45):", packagefolder.Max(12, 45))
}
