package main

import (
	"fmt"
	"reflect"
)

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

	// Rune example (alias for int32, represents a Unicode code point)
	var m rune = '♥'
	fmt.Printf("Rune: %c, Unicode: %U, Integer: %d\n", m, m, m)

	// Complex number declaration and operations
	var c1 complex128 = 2 + 3i
	c2 := complex(4, -2) // complex128 by default
	sum := c1 + c2
	product := c1 * c2

	fmt.Printf("c1: %v, c2: %v\n", c1, c2)
	fmt.Printf("Sum: %v, Product: %v\n", sum, product)
	fmt.Printf("Real part: %v, Imaginary part: %v\n", real(product), imag(product))

	//pointers

	fmt.Println(a)

	var iptr *int = &a
	fmt.Println(iptr)
	fmt.Println(&a)
	fmt.Println(*iptr)

	var f5 float64 = 5.55
	var fptr *float64 = &f5

	fmt.Println(f5, fptr, *fptr)

	fmt.Println(reflect.TypeOf(f5))

	fmt.Println(reflect.TypeOf(f5).Kind())

	var d1 string = "Dhyan"
	fmt.Println(d1[0], d1[1], d1[2])
	fmt.Printf("%c, %c %c\n", d1[0], d1[1], d1[2])

	var sptr *string = &d1
	fmt.Println((*sptr)[0], (*sptr)[1])

}

func Calc(a int, b int) (c int, d int) {
	c = a + b
	d = a - b

	return
}
