package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
)

func main() {
	/* My Playground */
	// WriteToFileNewMethod()

	// Will write to terminal
	// will write to file
	// myIOTest()

	// myBufferTest()
	readerInterfaceImplementations()

}

func readerInterfaceImplementations() {

	/* os.File Implements io.Reader */
	file, _ := os.Create("./output.txt")
	file.Read(make([]byte, 8)) // The os.File implements both io.Reader and io.Writer interfaces.

	/* bufio.Reader Implements io.Reader */
	reader := bufio.NewReader(os.Stdin)
	reader.Read(make([]byte, 8)) // The bufio.Reader implements the io.Reader interface.

	/* strings.Reader Implements io.Reader */
	strReader := strings.NewReader("Hello, World!")
	strReader.Read(make([]byte, 8)) // The strings.Reader implements the io.Reader interface.

	/* net.Conn Implements io.Reader and io.Writer */
	// conn, _ := net.Dial("tcp", "google.com:80")
	// conn.Read(make([]byte, 8)) // The net.Conn interface implements both io.Reader and io.Writer interfaces.

	/* Http Response Body Implements io.Reader */
	res, _ := http.Get("https://www.google.com")
	defer res.Body.Close()
	// Will read only 8 bytes ( 8 characters ) from the response body stream
	byteSlice := make([]byte, 8)
	// To read all , use io.ReadAll or loop with res.Body.Read(byteSlice) until EOF
	res.Body.Read(byteSlice) // The http.Response.Body implements the io.Reader interface.
	fmt.Println("Read from Http Response Body : ", string(byteSlice))
}

func myIOTest() {
	// Writer can be on any type that implements the Write method, such as os.Stdout, a file, or a network connection.
	// Example of Stdout Writer
	var w io.Writer = os.Stdout // w is an io.Writer that writes to standard output
	io.WriteString(w, "Hello, World!\n")

	// Exmaple of file writer
	file, _ := os.Create("./output.txt")
	var y io.Writer = file // y is an io.Writer that writes to a file
	io.WriteString(y, "Test")
	// Example of network connection writer
	conn, _ := net.Dial("tcp", "google.com:443")
	var z io.Writer = conn // z is an io.Writer that writes to a network connection
	io.WriteString(z, "GET / HTTP/1.0\r\n\r\n")
}

func myBufferTest() {

	r1 := strings.NewReader("My New Reader Buffer") // Create a buffered reader for standard input
	buf := make([]byte, 8)                          // Create a byte slice to hold the data

	/* Without Buffer */
	/* Go Allocates a default buffer (new for each call) */
	io.Copy(os.Stdout, r1)

	/* With Buffer */
	/* We can reuse the same buffer for multiple callse to avoid too many allocations */
	/* Can Match the underlying device page size (4096 bytes) or its multiples for better performance */
	io.CopyBuffer(os.Stdout, r1, buf)
}

func myNetworkTest() {
	res, _ := http.Get("https://www.google.com")
	defer res.Body.Close()
	/*
		We use io.ReadAll (Reader) on response body
		since the response body is of type interface ReadCloser which is an interface
		wrapping 'io,Reader' and 'io.Closer' interfaces
		i.e. res.body is io.Reader so we can use io operations to read the stream.
	*/
	body, _ := io.ReadAll(res.Body)
	println(string(body))
}

func WriteToFileNewMethod() {

	myFile, err := os.Create("./test.txt")
	if err != nil {
		panic(err)
	}

	var write io.Writer = myFile // write := io.Writer(myFile)
	io.WriteString(write, "Test to File using io.Writer interface")
}

func bufferedReadOperation() {
	/* Buffered Read Operation */
	/* This is a common pattern to read large files without loading the entire file into memory */
	/* We can use bufio package for buffered reading */

	/* Example of using bufio to read a file line by line */
	file, err := os.Open("largefile.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	/* Create a new scanner to read the file */
	/* Set the buffer size */
	bufio.NewScanner(file).Buffer(make([]byte, 0, 64*1024), 1024*1024) // Set buffer size MAX to 1MB

	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// line := scanner.Text()
	// fmt.Println(line) // Process the line (e.g., print it)
	// }

}
