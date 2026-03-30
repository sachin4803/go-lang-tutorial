package main

// ─────────────────────────────────────────────────────────────────────────────
// 02 · VARIABLES & TYPES
//
// Topics covered:
//   • Variable declaration styles (var, :=, multiple)
//   • Zero values
//   • All basic types: bool, int family, float family, string, byte, rune, complex
//   • Type conversion
//   • Constants and iota
//   • Custom type aliases
//   • reflect.TypeOf — inspecting types at runtime
//   • Blank identifier _
// ─────────────────────────────────────────────────────────────────────────────

import (
	"fmt"
	"reflect"
)

// ─────────────────────────────────────────────
// CONSTANTS & IOTA
//
// const values are evaluated at compile time and cannot be changed.
// iota is an auto-incrementing integer counter inside a const block.
// ─────────────────────────────────────────────
const appName string = "GoTutorial"
const version = 1.0 // type inferred as float64

// iota — great for enumerations
type Direction int

const (
	North Direction = iota // 0
	East                   // 1
	South                  // 2
	West                   // 3
)

func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}

// ─────────────────────────────────────────────
// CUSTOM TYPE ALIAS
//
// 'type X Y' creates a new distinct type based on Y.
// Useful for domain modelling and type safety.
// ─────────────────────────────────────────────
type Celsius float64
type Fahrenheit float64

func celsiusToFahrenheit(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// ─────────────────────────────────────────────
// VariablesAndTypesDemo — run all variable and type examples
// ─────────────────────────────────────────────
func VariablesAndTypesDemo() {

	// ── 1. Declaration styles ─────────────────
	var a int = 10         // explicit type + value
	var b = 20             // type inferred from value (int)
	c := 30                // short declaration — only inside functions
	var d, e int           // multiple vars, zero value (0 for int)
	var f, g = 40, "hello" // multiple vars, inferred types

	fmt.Println("a:", a, "b:", b, "c:", c)
	fmt.Println("d (zero):", d, "e (zero):", e)
	fmt.Println("f:", f, "g:", g)

	// ── 2. Zero values ────────────────────────
	// Every uninitialized variable gets a default "zero value".
	var zeroInt int       // 0
	var zeroFloat float64 // 0.0
	var zeroBool bool     // false
	var zeroStr string    // "" (empty string)
	fmt.Println("zero values →", zeroInt, zeroFloat, zeroBool, zeroStr)

	// ── 3. Boolean ────────────────────────────
	var isActive bool = true
	isAdmin := false
	fmt.Println("bool →", isActive, isAdmin)

	// ── 4. Integer types ──────────────────────
	// int   — platform size (32 or 64 bit)
	// int8  — -128 to 127
	// int16 — -32768 to 32767
	// int32 — -2147483648 to 2147483647
	// int64 — very large range
	// uint  — unsigned (no negatives)
	var i8 int8 = 127
	var i16 int16 = 32767
	var i64 int64 = 9223372036854775807
	fmt.Printf("int8: %d, int16: %d, int64: %d\n", i8, i16, i64)

	// ── 5. Float types ────────────────────────
	// float32 — 32-bit, less precision (~7 decimal digits)
	// float64 — 64-bit, more precision (~15 decimal digits) — DEFAULT
	var f32 float32 = 3.14159
	var f64 float64 = 2.718281828459045
	literalFloat := 1.618 // inferred as float64
	fmt.Printf("float32: %f\nfloat64: %.15f\nliteral: %f\n", f32, f64, literalFloat)

	// ── 6. String & byte ─────────────────────
	// Strings in Go are immutable sequences of bytes (UTF-8 encoded).
	// byte is an alias for uint8 — represents a single raw byte.
	word := "Golang"
	fmt.Printf("string: %s, length: %d\n", word, len(word))
	fmt.Printf("byte at index 0: %d (%c)\n", word[0], word[0])

	// ── 7. Rune ───────────────────────────────
	// rune is an alias for int32 and represents a Unicode code point.
	// Use single quotes for rune literals.
	var r1 rune = 'A'
	var r2 rune = '♥'
	fmt.Printf("rune 'A' — char: %c, unicode: %U, int: %d\n", r1, r1, r1)
	fmt.Printf("rune '♥' — char: %c, unicode: %U, int: %d\n", r2, r2, r2)

	// ── 8. Complex numbers ────────────────────
	// complex64  — float32 real + float32 imaginary
	// complex128 — float64 real + float64 imaginary (DEFAULT)
	var c1 complex128 = 2 + 3i
	c2 := complex(4, -2) // complex128 by default
	sum := c1 + c2
	product := c1 * c2
	fmt.Printf("c1: %v, c2: %v\n", c1, c2)
	fmt.Printf("sum: %v, product: %v\n", sum, product)
	fmt.Printf("real: %v, imag: %v\n", real(product), imag(product))

	// ── 9. Type conversion ────────────────────
	// Go does NOT do implicit conversion — you must convert explicitly.
	var intVal int = 42
	floatVal := float64(intVal) // int → float64
	backToInt := int(floatVal)  // float64 → int (truncates decimal)
	fmt.Printf("int→float64: %f, float64→int: %d\n", floatVal, backToInt)

	// ── 10. Constants ─────────────────────────
	const pi = 3.14159265358979
	fmt.Println("appName:", appName, "version:", version, "pi:", pi)

	// ── 11. iota ──────────────────────────────
	fmt.Printf("Directions — North:%d East:%d South:%d West:%d\n",
		North, East, South, West)
	fmt.Println("Direction.String():", North.String(), South.String())

	// ── 12. Custom type alias ─────────────────
	boiling := Celsius(100)
	fmt.Printf("%.1f°C = %.1f°F\n", boiling, celsiusToFahrenheit(boiling))

	// ── 13. reflect.TypeOf — inspect types at runtime ─
	fmt.Println("type of f64:", reflect.TypeOf(f64))         // float64
	fmt.Println("kind of f64:", reflect.TypeOf(f64).Kind())  // float64
	fmt.Println("type of word:", reflect.TypeOf(word))       // string
	fmt.Println("type of boiling:", reflect.TypeOf(boiling)) // main.Celsius

	// ── 14. Blank identifier _ ────────────────
	// Use _ to discard values you don't need (avoids "declared and not used" error).
	sum2, _ := 10, 99 // we only care about sum2
	fmt.Println("sum2 (99 discarded with _):", sum2)
}
