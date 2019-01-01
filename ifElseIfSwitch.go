package main

import "fmt"

func main() {

	// If Statement

	yourAge := 18
	for i := 1; i < 4; i++ {
		if i == 3 {
			yourAge = 14
		}

		if yourAge >= 18 {

			fmt.Println("You Can have lots of GF!")
		} else if yourAge >= 16 {
			fmt.Println("You Can Drive")

		} else {
			fmt.Println(yourAge)

		}
		yourAge = 17
	}

	//return

	// You can use else if perform different actions, but once a match
	// is reached the rest of the conditions aren't checked
	if yourAge >= 16 {
		fmt.Println("You Can Drive")
	} else if yourAge >= 18 {
		fmt.Println("You Can Vote")
	} else {
		fmt.Println("You Can Have Fun")
	}
	// Switch statements are used when you have limited options
	fmt.Println(yourAge)

	switch yourAge {
	case 16:
		fmt.Println("Go Drive")
	case 17:
		fmt.Println("Go Drive AGAIN")

	case 18:
		fmt.Println("Go Vote")
	default:
		fmt.Println("Go Have Fun")
	}

	return

}
