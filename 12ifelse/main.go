package main

import "fmt"

func main() {

	number := 10

	if number > 10 { // Cannot add endline here, will cause syntax error
		fmt.Println("> 10")
		// In GO, we can create a new variable , assign them and check (do operations) on them inside the condition
	} else if newNumber := number - 10; newNumber < 10 {
		fmt.Println(" < 5")
	} else {
		fmt.Println("10")
	}

	/* My Tests Conclusion */
	/* Using more than 1 variable assignement  */

	if newNumber := number + 10; /*secondNumber := newnewNumber+1;*/ newNumber > 10 { // Cannot add endline here, will cause syntax error
		fmt.Println("> 10")
		// In GO, we can create a new variable , assign them and check (do operations) on them inside the condition
	} else if test := make(map[int]string); len(test) < 10 && test != nil {
		fmt.Println(" < 5")
	} else {
		fmt.Println("10")
	}

}
