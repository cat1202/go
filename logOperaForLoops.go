package main

import "fmt"

func main() {

	// Logical Operators
	fmt.Println("true && false =", true && false)
	fmt.Println("true || false =", true || false)
	fmt.Println("!true =", !true)

	// For loops
	i := 1
	for i <= 5 {
		fmt.Println(i)
		// Shorthand for i = i + 1
		i++
	}
	// Relational Operators include ==, !=, <, >, <=, and >=
	// You can also define a for loop like this, but you need semicolons
	for j := 0; j < 3; j++ {

		fmt.Println(j)

	}
	fmt.Println("----")
	//next, u can reuse j again!!!
	for j := 3; j < 6; j++ {

		fmt.Println(j)

	}

}
