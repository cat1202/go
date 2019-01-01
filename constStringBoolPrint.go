package main

//import "fmt"
import (
	"fmt"
	"strings"
	//"unicode/utf8"
)

func main() {

	// A constant is a variable with a value that can't be changed

	const pi float64 = 3.14159265359

	// You can declare multiple variables like this

	var (
		varA   = 2
		varB   = 3
		varCat = 4
	)

	fmt.Println(varA, varB, varCat)

	// Strings are a series of characters between " or `

	var myName string = "Derek Banas"
	catName := "this is a pretty cat!"

	// Get string length

	fmt.Println(len(myName), "\rlen of catName is:", len(catName))
	//var catMask string = "abc"
	catMask := strings.Repeat("#", len(catName))
	fmt.Println(catMask)

	//fmt.Println(50 * "---")
	/*
	   It has a function instead of an operator, strings.Repeat. Here's a port of your Python example, which you can run here:

	   package main

	   import (
	       "fmt"
	       "strings"
	       "unicode/utf8"
	   )

	   func main() {
	       x := "my new text is this long"
	       y := strings.Repeat("#", utf8.RuneCountInString(x))
	       fmt.Println(x)
	       fmt.Println(y)
	   }
	   Note that I've used utf8.RuneCountInString(x) instead of len(x); the former counts "runes" (Unicode code points), while the latter counts bytes. In the case of "my new text is this long", the difference doesn't matter since all the characters are only one byte, but it's good to get into the habit of specifying what you mean
	*/

	// You can combine strings with +

	fmt.Println(myName + " is a robot")

	// Strings between " can contain escape symbols like \n for newline

	fmt.Println("I like \n \n")

	fmt.Print("NO new line here<---")
	fmt.Println("Newlines")

	// Booleans can be either true or false

	var isOver40 bool = true
	catDied := false
	fmt.Println(isOver40, "is cat died:", catDied)
	// Printf is used for format printing (%f is for floats)
	fmt.Println(pi)
	fmt.Printf("%f \n", pi)

	// You can also define the decimal precision of a float
	fmt.Printf("%.3f \n", pi)

	// %T prints the data type
	fmt.Printf("%T \n", pi)
	fmt.Printf("%T \n", "Gary is great")
	fmt.Printf("%T\n", 12345)
	// %t prints booleans
	fmt.Printf("%t\n", isOver40)
	fmt.Printf("%T\n", isOver40)

	// %d is used for integers
	fmt.Printf("%d \n", 100)
	// %b prints in binary
	fmt.Printf("%b \n", 100)
	// %c prints the character associated with a keycode
	fmt.Printf("%c \n", 44)
	// %x prints in hexcode
	fmt.Printf("%x \n", 17)
	fmt.Printf("%d(d) %x(x) %b(b)\n", 17, 17, 17)
	fmt.Printf("%d(d) %x(x) %b(b)\n", 7, 7, 7)
	fmt.Printf("%b(b) %b(b) %b(b)\n", 6, 4, 5)
	fmt.Printf("%x(x) %x(x) %x(x)\n", 199, 237, 240)
	// %e prints in scientific notation

	fmt.Printf("%e \n", pi)

}
