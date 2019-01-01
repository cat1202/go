package main

import "fmt"

func main() {

	// An Array holds a ***FIXED*** number of values of the same type
	var favNums2 [5 + 1]float64
	favNums2[0] = 163
	favNums2[1] = 78557
	favNums2[2] = 691
	favNums2[3] = 3.141
	favNums2[4] = 1.618
	fmt.Println(favNums2)
	// You access the value by supplying the index number
	fmt.Println(favNums2[3])

	// Another way of initializing an array
	favNums3 := [5]float64{1, 2, 3, 4, 5}
	//return
	// How to iterate through an array (Use _ if a value isn't used)
	for i, value := range favNums3 {
		fmt.Println("value is:", value, "index is:", i)
	}
	//return
	// Slices are like arrays but you leave out the size
	numSlice := []int{5, 4, 3, 2, 1}
	// You can create a slice by defining the first index value to
	// take through the last,NOT include last,
	numSlice2 := numSlice[3:5] // numSlice2 == [2,1]
	fmt.Println(numSlice2, len(numSlice2))
	//return //out: [2 1] 2

	fmt.Println("numSlice2[0] =", numSlice2[0])

	// If you don't supply the first index it defaults to 0
	// If you don't supply the last index it defaults to max
	fmt.Println("numSlice[:2] =", numSlice[:2]) //AS:5,4
	fmt.Println("numSlice[2:] =", numSlice[2:]) //AS:3,2,1
	//return

	// You can also create an empty slice and define the data type,
	//initial length(5) (receive value of 0), capacity (max size)
	numSlice3 := make([]int, 5, 10)
	fmt.Println(numSlice3, len(numSlice3)) //->[0,0,0,0,0,0] 5
	//numSlice3 := make([]int, 5, 10)
	//fmt.Println(numSlice3, len(numSlice3)) //->[0,0,0,0,0,0] 5

	//return
	// You can copy a slice to another
	copy(numSlice3, numSlice)
	fmt.Println(numSlice3[0], len(numSlice3))

	// Append values(any number of value) to the end of a slice
	numSlice3 = append(numSlice3, 0, -1, 1, 9)
	fmt.Println(numSlice3, len(numSlice3))
	fmt.Println(numSlice3[6])

}
