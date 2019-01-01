package main

import "fmt"

func main() {

	// A Map is a collection of
	// Map[key] = value  pairs
	// Created with
	// varName := make(map[keyType] valueType)
	presAge := make(map[string]int)
	presAge["gary age is"] = 42
	presAge["kubera age is"] = 22
	presAge["hotcat"] = 11
	fmt.Println(presAge)
	fmt.Println(presAge["gary age is"])
	// Get the number of items in the Map
	fmt.Println(len(presAge))
	//return
	// The size changes when a new item is added
	presAge["Jack Xu"] = 55
	fmt.Println(len(presAge))
	// We can delete by passing the key to delete
	//no eror or exception when key didn't exist
	fmt.Println(presAge)
	delete(presAge, "John F. Kennedy")
	delete(presAge, "Jack Xu")
	fmt.Println(presAge)
	delete(presAge, "hotCat")
	fmt.Println(presAge)
	fmt.Println(len(presAge))

}
