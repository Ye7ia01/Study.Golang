package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter Rating")
	var reader *bufio.Reader
	reader = bufio.NewReader(os.Stdin)

	/* ReadString retrns (string,error) */
	/* we can read input, error to parse error */
	/* This is simulation to 'Catch' in other languages */
	input, _ := reader.ReadString('\n')
	fmt.Println("Your Rating is : ", input)

	/* My Test */
	/* Conclusion:
	1) error is a datatype in GO
	2) (type,error) can be returne to apredefined variables, not necessary to use :=
	*/
	var input2 string
	var err error
	input2, err = reader.ReadString('\n')
	fmt.Println("Input 2 : ", input2)
	fmt.Println("Error : ", err)

}
