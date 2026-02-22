package main

import (
	"fmt"
	"sort"
)

func main() {

	/* Array data type vs slice data type */
	// Array
	var myarr [5]string
	fmt.Printf("Type of var %T \n", myarr)

	// slice
	// Difference is that slice does not take a 'typeLength' parameter
	var myslice []string
	fmt.Printf("Type of slice %T \n", myslice)

	// Another way to declare slices
	// Although we did provide the lenght in the make , this is still slice
	// the inital legnth is 10 but can be extended using the 'append' method
	mySecondSlice := make([]int, 0)
	fmt.Printf("Type of slice %T \n", mySecondSlice)

	mySecondSlice = append(mySecondSlice, 3, 4, 7, 1, 3, 8, 6, 4, 3, 2, 4, 6, 8, 6, 4)
	fmt.Println("My New Slice : ", mySecondSlice)

	sort.Ints(mySecondSlice)
	fmt.Println(mySecondSlice)

	/* Slicing the slices */
	// IMPORTANT, Last Index Number is NON-INCLUSIVE
	// i.e. this gets from index 1 to index
	myNewSlice := mySecondSlice[1:6]
	fmt.Println(myNewSlice)

	fmt.Println("Capacity ", cap(myNewSlice))

	// Removing Slice elements
	top5Langs := []string{"C", "JAVA", "JAVASCRIPT", "PYTHON", "GO"}

	// remove index 2 (javacript)
	top5Langs = append(top5Langs[:2], top5Langs[3:]...)
	fmt.Println(top5Langs)

}
