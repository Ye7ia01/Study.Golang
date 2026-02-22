package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {

	/* os.File Implements io.Reader */
	fmt.Println("Impemeneting io.Reader for File")
	file, _ := os.Open("./output.txt")
	// file.WriteString("Test")
	count, err := countAlphapets(file)
	if err != nil {
		fmt.Println("Error counting alphapets in file : ", err)
	}
	fmt.Println("Number of Alphapets in file : ", count)

	// /* bufio.Reader Implements io.Reader */
	// fmt.Println("Impemeneting io.Reader for bufio.Reader, Enter input : ")
	// reader := bufio.NewReader(os.Stdin)
	// input, _ := reader.ReadString('\n')
	// fmt.Println("Your Input : ", input)
	// count2, err2 := countAlphapets(reader)

	// if err2 != nil {
	// 	fmt.Println("Error counting alphapets in bufio.Reader : ", err2)
	// }
	// fmt.Println("Number of Alphapets in bufio.Reader : ", count2)

	/* strings.Reader Implements io.Reader */
	fmt.Println("Impemeneting io.Reader for strings.Reader")
	strReader := strings.NewReader("Hello, World!")
	count3, err3 := countAlphapets(strReader)
	if err3 != nil {
		fmt.Println("Error counting alphapets in strings.Reader : ", err3)
	}
	fmt.Println("Number of Alphapets in strings.Reader : ", count3)

	/* Http Response Body Implements io.Reader */
	fmt.Println("Impemeneting io.Reader for Http Response Body")
	res, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	count5, err5 := countAlphapets(res.Body)
	if err != nil {
		fmt.Println("Error counting alphapets in Http Response Body : ", err5)
	}
	fmt.Println("Number of Alphapets in Http Response Body : ", count5)

}

/* @Params : Interace implementation for io.Reader */
func countAlphapets(reader io.Reader) (int, error) {

	count := 0
	buf := make([]byte, 1024)

	for {
		// Keeps Internal Cursor
		// All we need to do is call the Read method with our buffer in a loop
		// Not safe in multithreaded environment as the internal cursor is shared between all threads
		n, err := reader.Read(buf)
		// Loop on read bytes and count the number of alphapets
		for _, val := range buf[:n] {
			if (val >= 'A' && val <= 'Z') || (val >= 'a' && val <= 'z') {
				count++
			}
		}

		if err == io.EOF {
			return count, nil
		}

		if err != nil {
			return 0, err
		}
	}

}
