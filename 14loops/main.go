package main

import "fmt"

func main() {

	/* Loops in Go */

	/* Default For Loop */
	fmt.Println("Default For Loop")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	/* Iterating Over an Iterable (slice) */
	slice := make([]string, 0)
	slice = append(slice, "Go")
	slice = append(slice, "is")
	slice = append(slice, "awesome")

	fmt.Println("For loop (foreach) slice")
	for index, value := range slice {
		fmt.Printf("Index : %d, Value : %s\n", index, value)
	}

	/* While loops */
	fmt.Println("While Loop")
	someValue := 0
	for someValue < 5 {
		if someValue == 3 {

			break // Exit the loop when someValue is 3
		}

		if someValue == 1 {
			someValue++
			continue // Skip the rest of the loop when someValue is 1
		}

		if someValue == 2 {
			goto SkipPrint // Jump to the label SkipPrint when someValue is 2
		}

		fmt.Println(someValue)
		someValue++
	}

SkipPrint:
	fmt.Println("Goto Demo Done")

}
