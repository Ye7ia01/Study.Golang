package main

import "fmt"

func main() {
	fmt.Println("Hello, Functions!")
	result := add(5, 10)
	fmt.Printf("The result of addition is : %d\n", result)

	/* Basi Variadic Function */
	total, hasNumbers := sum(1, 2, 3, 4, 5)
	fmt.Printf("Success status %v , total is : %v", hasNumbers, total)
}

/* Basic FUnction Declaration */
func add(a int, b int) int {
	return a + b
}

/* Variadic Function */
func sum(numbers ...int) (int, bool) {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total, len(numbers) > 0
}
