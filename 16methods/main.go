package main

import "fmt"

func main() {

	var myUser User = User{"John Doe", 30, "mail.com", true}
	var myPtr *User = &myUser

	/* Calling method with value receiver */
	myPtr.Greet()                                       // Output : Hello, John Doe
	myPtr.UpdateEmail("newmail.com")                    // Update the email of the user
	fmt.Println("Original User Email : ", myUser.Email) // Output : Original User Email :  newmail.com
}

type User struct {
	Name       string
	Age        int
	Email      string
	isLoggedIn bool
}

/* Method with value receiver */
/* The difference between a method and a function is that a method has a receiver argument.
/* The receiver is the type (struct) that the method is associated with.
/* In this case, the Greet method is associated with the User struct, and it can be called on any instance of User. */

func (u User) Greet() string {
	return "Hello, " + u.Name
}

/*
	IMPORTANT: Methods with value receivers operate on a copy of the struct,

so any changes made to the struct within the method will not affect the original struct.
To Avoid that, we can use pointer receivers instead of value receivers.
*/
func (receiver *User) UpdateEmail(newEmail string) {
	receiver.Email = newEmail
}
