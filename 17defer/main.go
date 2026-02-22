package main

import (
	"fmt"
)

func main() {

	/* Defer Tutorial */
	fmt.Println("Defers Tutorial")

	/*
		Defer Keyword is telling the compiler to POSTPONE the execution of this line (expression)
		to just before ending the execution flow (in funciton just before returns).

		Defer executes as LIFO (Last In First Out), so if multiple defers are used
		Expect them to be executed in the reverse order they were declared.

		Defer must be a function call and cannot be expression.

		USE CASES :
			Clean Up without code duplication : connection.close() // will run before exit
			Clean Up even in Panic cases : Even if moved the logic at the end of funcitons , In case of Panics the code will nut run, but with defers It will run even if in panic mode

	*/

	fmt.Println("Defer Demo : ", deferDemo())

	defer fmt.Println("Last to be printed")
	defer fmt.Println("Second Last to be printed")
	defer fmt.Println("Third Last to be Printed")

}

/*
VERY IMPORTANT: Note that "result" is a return value and not parameter
*/
func deferDemo() (result int) {

	/*
		Useful Example showing exactly where defer runs
		Overwriting a return result in Defer will return the new value
		i.e. It runs exactly AFTER return BEFORE function exits.
		Evidence is that we returned 42 but yet the return is 100.
		This is because at the end of the function (before exit) the defer executed
	*/
	defer func(result *int) {
		fmt.Println("Executing Defer")
		*result = 100
	}(&result)

	return 42 // automatically returns 'result'
}

/* Example with and without using DEFERS to understand why its useful */
// WITHOUT DEFER - messy and error-prone
// func unsafe() {

// 	f := openFile()
// 	// ... code ...
// 	f.Close() // This won't run if panic occurs!
// }

// func safe() {
// 	f := openFile()
// 	defer f.Close() // This runs even during panic!
// 	// ... code ...
// 	panic("oops") // f.Close() still executes
// }
