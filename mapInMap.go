package main

import "fmt"

func main() {

	// We can store multiple items in a map as well
	superhero := map[string]map[string]string{
		"Superman": map[string]string{
			"realname": "Clark Kent",
			"city":     "Metropolis",
			"age":      "120",
		},

		"Batman": map[string]string{
			"realname": "Bruce Wayne",
			"city":     "Gotham City",
			"age":      "86",
		},
	}
	fmt.Println(superhero, len(superhero), "---->")

	//return
	// We can output data where the key matches Superman
	if temp, hero := superhero["Batman"]; hero {
		fmt.Println(temp["realname"], temp["city"])
	}

}
