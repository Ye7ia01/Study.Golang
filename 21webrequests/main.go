package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const baseUrl = "http://localhost:8000/"

func main() {

	// GetRequest(baseUrl + "get")
	PostFormRequest(baseUrl + "postform")

}

/* First Letter is capital as Public Function */
func GetRequest(path string) {

	/* Intiate Get Request */
	fmt.Println("Initating Get Request to ", path)
	response, err := http.Get(path)

	/* Check Errors */
	if err != nil {
		// fmt.Println("Request Failed with Status Code ", response.StatusCode)
		panic(err)
	}

	/* Clean Up after finish */
	defer response.Body.Close()

	/* Read Content via strings.builder and ioutil*/
	var stringResponse strings.Builder
	content, err := io.ReadAll(response.Body)
	stringResponse.Write(content)

	if err != nil {
		fmt.Println("Failed reading Content body")
		panic(err)
	}

	fmt.Println("Response Content : ", stringResponse.String())
}

func PostRequest(path string) {

	fmt.Println("Initiating Post request to ", path)
	requestBody := strings.NewReader(`
	{
		"coursename":"reactjs"
	}
	`)

	/* Intiate Post Request */
	/* Note the body parameter is io.Reader */
	content, err := http.Post(path, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	body, _ := io.ReadAll(content.Body)

	defer content.Body.Close()

	fmt.Println("Response : ", string(body))

}

func PostFormRequest(path string) {

	// Map (Hash Table Data structure)
	formValues := url.Values{}
	formValues.Add("firstName", "Yehia")
	formValues.Add("lastName", "Abdelhady")
	formValues.Add("Age", "28")

	response, err := http.PostForm(path, formValues)
	if err != nil {
		panic(err)
	}

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))

}
