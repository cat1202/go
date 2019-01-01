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
	//fmt.Println(superhero, len(superhero), "---->")
	fmt.Println("--------")
	mytmp, myhero := superhero["Batman"]
	fmt.Println(mytmp, "\n\n")
	fmt.Println(myhero, "\n\n")

	//return

	// We can output data where the key matches Superman
	//1.if find item, assgin item to temp,and 'true' to hero
	//2; hero, is true ,then print content...
	if temp, hero := superhero["Batman"]; hero {
		fmt.Println(temp["realname"], temp["city"])
		fmt.Println(hero, "\n\n")
	}
	//if temp  := superhero["Batman"]  {
	//	fmt.Println(temp["realname"], temp["city"])
	//}
	tmp := superhero["Batman"]
	fmt.Println(tmp)
	//if (fmt == ]) {
	//	fmt.Println(tmp["realname"], tmp["city"])
	//}
	tmp2 := superhero["Batman222"]
	fmt.Println(tmp2)

}
