package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://www.google.com"

func main() {

	fmt.Println("Initiating Get request to google.com")
	// response is a pointer : *http.response
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	/*
		Even if http is stateless and we are not maintaining any connection, we still need to close the response body
		to free up resources and avoid memory leaks.
		Since new http requests can reuse this one if closed instead of creating a new one wih its resources.
	*/
	defer response.Body.Close() // Close the response body when the function returns

	if response.Status == "200 OK" {
		fmt.Println("Received Request Successfully ... ")
		// reader := bufio.NewReader(os.Stdin)
		// input, _ := reader.ReadString('\n')
		// fmt.Println(input)

		/*
			IMPORTANT :
			The response body is retreived on demand when requested
			It is read from the open tcp connection , thats why we request to read the body using io.ReadAll
		*/

		content, err := io.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println("Res Body : ", string(content))
	}
}
