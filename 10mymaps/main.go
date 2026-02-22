package main

import "fmt"

func main() {

	/* Maps ia HASH TABLE data structure */
	/* All CRUD ops should be O(1) in Best Case to O(n) in Worst Case */
	/* Hash Tables are Key, Value Entries */
	/* Each key is given to hash function that returns a hash code */
	/* This hash code is then mapped to an array index (probably % len) */

	/* In cases of a coliision (2 keys ended up with the same index) :
	The index (bucket) is now a linked list storing all entries with this key
	*/

	/* A map with KEY (int) & VALUE(STRING) */
	myMap := make(map[int]string)

	for i := 0; i < 10; i++ {
		myMap[i] = string(i) + "My Test"
	}

	fmt.Println("My Map : ", myMap)

	/* DELETE a map index by using the delete object */
	delete(myMap, 5)

	fmt.Println("My Map : ", myMap)

	// var mymap2 map[int]string
	// mymap2 = make(map[int]string)

}
