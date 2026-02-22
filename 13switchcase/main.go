package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	/* Generate a random number from 1 to 6 */
	// rand.Seed(time.Now().Nanosecond()) // Deprecated
	rand.NewSource(time.Now().UnixNano()) // Seed unique time in nanoseconds
	dice := rand.Intn(5) + 1
	fmt.Println("Dice Roll : ", dice)

	switch dice {
	case 1:
		fmt.Println("Move 1 ")
	case 2:
		fmt.Println("Move 2 ")
	case 3:
		fmt.Println("Move 3 ")
		fallthrough // Run the next case also
	case 4:
		fmt.Println("Move 4")
	case 5:
		fmt.Println("Move 5 ")
		fallthrough
	case 6:
		fmt.Println("Move 6 ")
	default:
		fmt.Println("Unknown move")
	}
}
