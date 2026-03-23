package main

import (
	"fmt"

	"examplegithub.com/src/packagefolder"
)

var a int

func init() {
	fmt.Println("Int func")
	a = 100
}

func main() {
	fmt.Println("Hello, World!")
	Hello()
	welcome()
	packagefolder.Hellofrompackage()
}
