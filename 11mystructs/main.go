package main

import "fmt"

func main() {

	/* IMPORTANT: Structs has no Interitance */

	/* Struct Delcaration (definition) */
	type User struct {
		// Fields (= properties in OOP)
		Name       string // field name is 'Name'
		Email      string // field name is 'Email'
		Age        int    // field name is 'Age'
		isLoggedIn bool   // field name is 'isLoggedIn'
	}

	/* Instantiation */
	var yehia User = User{"Yehia", "ye7iaabdelhady@gmail.com", 28, true}
	// %+v formats, print field names alongside the field values
	fmt.Printf("Yehia : %+v \n", yehia)

	ahmed := User{"Ahmed", "Ahmed@gmail.com", 28, true}
	fmt.Printf("Ahmed :  %+v \n", ahmed)

}
