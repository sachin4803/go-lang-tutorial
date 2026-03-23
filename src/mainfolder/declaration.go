package main

import "fmt"

func declare() {

	//Declared in init
	fmt.Println("a variable from init:", a)

	// Variable declarations

	var a int = 10         // explicit type and value
	var b = 20             // type inferred
	c := 30                // short declaration (type inferred)
	var d, e int           // multiple variables, default zero value (0)
	var f, g = 40, "hello" // multiple variables, inferred types

	// Rune type (alias for int32, represents a Unicode code point)
	var r rune = 'A' // single quotes for rune/character
	fmt.Printf("Rune r: %c, Unicode: %U, Value: %d\n", r, r, r)

	// Print variables and constants
	fmt.Println("a:", a, "b:", b, "c:", c)
	fmt.Println("d:", d, "e:", e)
	fmt.Println("f:", f, "g:", g)

	// Constant declarations
	const pi = 3.14
	const name string = "GoLang"

	fmt.Println("pi:", pi, "name:", name)

	var x float32 = pi
	var y float64 = pi
	fmt.Println("Convert from int to float", x, y)

	//Type declaration
	type apples int
	var kasmirApples apples = 100
	var normalApples apples = 300
	fmt.Println("apple variables: ", kasmirApples, normalApples)

	c, d = Calc(int(normalApples), int(kasmirApples))
	fmt.Println("sum of apples: ", c, " and substraction of apples: ", d)

}

func Calc(a int, b int) (c int, d int) {
	c = a + b
	d = a - b

	return
}
