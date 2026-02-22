package main

import (
	"fmt"
	"net/url"
)

const myURL = "https://www.google.com?query=golang+tutorial&payment_id=12345"

func main() {

	/* Package URL has methods to parse and manipulate URLs */
	result, err := url.Parse(myURL)
	if err != nil {
		panic(err)
	}

	/* After Parsing the URL, we can access all query parameters, host, protocl , etc.. */

	// fmt.Println(result.Scheme) // Protocol
	// fmt.Println(result.Host)   // Hostname
	// fmt.Println(result.Path)   // page route
	// fmt.Println(result.Port()) // port number

	queryParams := result.Query()

	for i, val := range queryParams {
		fmt.Println("Key : ", i)
		fmt.Println("Value : ", val)
	}

	/* We can also construct a URL from its components */

	/* The & sign is very Important , a rule of thum when manipulating data objects, always use poiters to make sure your not manipulating copies */
	constructedURL := &url.URL{
		Scheme: "https",
		Host:   "www.google.com",
		Path:   "/search",
	}

	fmt.Println("Constructed URL : ", constructedURL.String())

}
