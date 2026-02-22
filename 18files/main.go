package main

import (
	"fmt"
	"os"
)

func main() {

	/* Manipulating Files */

	/* 1- Creating a file using OS */
	/* There are so many ways to create a file , OD is most common */

	fileDescriptor, err := os.Create("./myfile.txt")
	checkError(err)

	// Will run at the end
	// Even If panic happens
	defer func() {
		fmt.Println("Closing File Descriptor")
		fileDescriptor.Close()
	}()

	/* 2- Write Some data */
	/* There are so many ways to to write to file , I think its better woth FileDescriptor methods */
	content := "My New File Content"
	fileDescriptor.WriteString(content)

	/* 3- Read Data */

	databytes, err := os.ReadFile(fileDescriptor.Name())
	checkError(err)
	fmt.Println("My Data is ", string(databytes))

	// Option 1 (My Study Option) Same as Keyboard Reader
	// reader := bufio.NewReader(fileDescriptor)
	// data, err := reader.ReadString('\n')
	// checkError(err)
	// fmt.Println("Data : ", data)

	// Option 2 (My Study)
	// bytes := make([]byte, 6) // Create a byte slice to hold the data
	// _, err = fileDescriptor.Read(bytes)
	// checkError(err) // Read data into the byte slice
	// // result := string(bytes)                     // Convert byte slice to string
	// fmt.Println("Read form files : ", bytes)

}

func checkError(err error) {
	if err != nil {
		panic(err) // Exit application (= system.exit(error_code))
	}
}
