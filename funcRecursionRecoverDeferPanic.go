package main

import "fmt"

// Functions allow us to reuse code and provide structure
// func funcName(parametersPassed) returnType
// Functions don't have access to any variables aside from those
// passed into it
func addThemUp(numbers []float64) float64 {

	sum := 0.0
	for _, val := range numbers {
		// Shorthand for
		//sum = sum + val
		sum += val
	}
	return sum
}

// Go functions can return multiple values
func next2Values(number int) (int, int, bool) {

	return number + 1, number + 2, true

}

// You can receive an undefined number of values with args ...int
func subtractThem(args ...int) int {

	finalValue := 0

	for _, value := range args {
		finalValue -= value
	}

	return finalValue

}

func main() {

	listOfNums := []float64{1, 2, 3}
	fmt.Println("Sum :", addThemUp(listOfNums))
	//return
	// Get 2 values from a function
	num1, num2, _ := next2Values(5)
	_, _, resBool := next2Values(1)
	fmt.Println(num1, num2, resBool)
	//return //->6 7 true

	//variable
	// Send an undefined number of values to a function (Variadic Function) args ...int) int
	fmt.Println(subtractThem(1, 2, 3, 10))
	//return //->-16

	// You can create a function in a function. It has access to the
	// local variables of the containing function
	// A function like this with no local variables is a closure
	num3 := 3
	doubleNum := func() int {
		num3 *= 2
		//return num3
		return 1
	}
	fmt.Println(doubleNum())
	fmt.Println(doubleNum())
	fmt.Println(num3)
	//return //-> 6,12
	// Calling a recursive function
	fmt.Println(factorial(3))
	//return

	// Defer executes a function after the inclosing function finishes
	// Defer can be used to keep functions together in a logical way
	// but at the same time execute one last as a clean up operation
	// Ex. Defer closing a file after we open it and perform operations
	defer printBef2()
	defer printTwo()
	printOne()
	//return
	// Use recover() to catch a division by 0 error
	var retVal int = 333
	retVal = safeDiv(3, 0)
	fmt.Println("retval =", retVal)

	fmt.Println("after div by 0")
	//return
	retVal = safeDiv(3, 2)
	fmt.Println("safediv(3,2)=", retVal)
	// We can catch our own errors and recover with panic & recover
	//return

	demPanic()

}

// Example of recursion : Function calls itself
// factorial(3)
// 3 * factorial(2) == 3 * 2
// 2 * factorial(1) == 2 * 1
// factorial(0) == 1

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

// Used to demonstrate defer

func printOne() { fmt.Println("in printone ", 1) }

func printTwo()  { fmt.Println("executed by defer...", 2) }
func printBef2() { fmt.Println("executed by defer...", 1.5) }

// If an error occurs we can catch the error with recover and allow
// code to continue to execute
func safeDiv(num1, num2 int) int {
	defer func() {
		recover()
		fmt.Println("defer in safeDiv")
		//fmt.Println(recover())
	}()
	solution := num1 / num2
	return solution
}

// Demonstrate how to call panic and handle it with recover
func demPanic() {

	defer func() {
		// If I didn't print the message nothing would show
		fmt.Println(recover())
		//fmt.Println("in demPanic", recover())

	}()
	panic("here is a PANIC")

}
