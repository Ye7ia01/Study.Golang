package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// 1- Initalize Std in buffer reader
	reader := bufio.NewReader(os.Stdin)

	// 2- Read user input
	fmt.Println("Enter Rating")
	userInput, _ := reader.ReadString('\n')

	// 3- Type Case string to float
	// strconv is a package that has many exported methods for typecasting
	// strings is a powerful package that has all string operations
	userRating, error := strconv.ParseFloat(strings.TrimSpace(userInput), 64)

	// 4- Check for errors
	// nil is equivilent to NULL in other languages
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println("Thanks for the rating, ", userRating+1)
	}

}
