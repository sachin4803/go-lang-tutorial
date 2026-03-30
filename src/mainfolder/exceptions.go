package main

import (
	"fmt"
)

func exception() {
	fmt.Println("inside exception example")

	x3, e := PanicExample(100, 0)
	if e != nil {
		fmt.Println("There is an error")
	} else {
		fmt.Println(x3)
	}

	// Example 1: defer
	DeferExample()

	// Example 2 & 3: panic and recover
	PanicRecoverExample()

}

func PanicExample(x1 int, x2 int) (x3 int, e error) {
	if x2 != 0 {
		x3 := x1 / x2
		fmt.Println(x3)
		return x3, nil
	} else {
		return 0, fmt.Errorf("Division is not possible")
	}
}

//1. defer , 2. panic, 3. recover()

// Schedules a function to run after the current function finishes (even if there’s a panic)
func DeferExample() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Outside")

}

// panic: Causes the program to stop normal execution and start panicking.
// recover: Used inside a deferred function to catch a panic and allow the program to continue.
func PanicRecoverExample() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}

	defer func() {
		if rec := recover(); rec != nil {
			fmt.Println(rec)
		}
	}()

	panic("Raise Panic")

	//fmt.Println("OutSide") -> wont run

}
