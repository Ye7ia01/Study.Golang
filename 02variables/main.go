package main

import "fmt"

/* The first character in the variabe name is
the access modifier identifier of the variable
*/
const LoginToken string = "xyz" // equivelient to public loginToken

func main() {
	var username string = "Yehia"
	fmt.Println(username)
	fmt.Printf("Type of var is %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Type of var is %T \n", isLoggedIn)

	/* Default Values */
	var defaultInt int
	fmt.Println(defaultInt)
	fmt.Printf("Type of var is %T \n", defaultInt)

	/* Default Value of String is '' */
	var defaultString string
	fmt.Println(defaultString)
	fmt.Printf("Type of var is %T \n", defaultString)

	/* Implicit declarations */
	var website = "www.google.com" // Notice we did not specify data type
	fmt.Println(website)

	/* No var */
	numberOfUser := 3000
	fmt.Println(numberOfUser)

}
