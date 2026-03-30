// Package packagefolder demonstrates how to create and use a Go package.
//
// Key rules:
//   - The package name matches the folder name.
//   - EXPORTED identifiers start with a capital letter — visible to importers.
//   - unexported identifiers start with a lowercase letter — private.
//   - Constants, variables, types and functions can all be exported.
package packagefolder

import (
	"fmt"
	"math"
	"strings"
)

// ─────────────────────────────────────────────
// EXPORTED CONSTANT
// Importers can read Version as packagefolder.Version
// ─────────────────────────────────────────────

// Version is the current version of this package.
const Version = "1.0.0"

// Pi is a higher-precision pi exposed by this package.
const Pi = math.Pi

// ─────────────────────────────────────────────
// EXPORTED TYPE
// ─────────────────────────────────────────────

// Point represents a 2D coordinate.
type Point struct {
	X, Y float64
}

// Distance returns the Euclidean distance from p to another Point.
func (p Point) Distance(other Point) float64 {
	dx := p.X - other.X
	dy := p.Y - other.Y
	return math.Sqrt(dx*dx + dy*dy)
}

// ─────────────────────────────────────────────
// MATH UTILITIES
// ─────────────────────────────────────────────

// Square returns n * n.
func Square(n int) int {
	return n * n
}

// Cube returns n * n * n.
func Cube(n int) int {
	return n * n * n
}

// Max returns the larger of a and b.
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Min returns the smaller of a and b.
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Abs returns the absolute value of n.
func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// IsPrime reports whether n is a prime number.
func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// ─────────────────────────────────────────────
// STRING UTILITIES
// ─────────────────────────────────────────────

// Repeat returns s repeated n times.
func Repeat(s string, n int) string {
	return strings.Repeat(s, n)
}

// Reverse returns the string s in reverse order.
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// IsPalindrome reports whether s reads the same forwards and backwards.
func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	return s == Reverse(s)
}

// ─────────────────────────────────────────────
// GREETING (original, kept + improved)
// ─────────────────────────────────────────────

// Hellofrompackage prints a greeting from this package.
func Hellofrompackage() {
	fmt.Println("Hello from packagefolder! (version", Version, ")")
}
