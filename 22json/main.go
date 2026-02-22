package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"name"` // JSON tag to specify the key name in JSON
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Tags     []string `json:"tags,omitempty"` // If epmpty remove from json
}

func main() {

	// jsonBytes := EncodeJson()
	// DecodeJson(jsonBytes)
	DecodeJsonAsMap()
}

/* Encode = Convert Go data to JSON format */
func EncodeJson() []byte {

	myCourses := []course{
		{"GO Lang Course", 322, "udemy.com", []string{"web-dev", "js"}},
		{"React Js", 322, "coursera.com", nil},
	}

	jsonObject, err := json.Marshal(myCourses)
	if err != nil {
		fmt.Println(jsonObject)
	}

	fmt.Println(string(jsonObject))
	return jsonObject

}

func DecodeJson(bytejson []byte) {

	myCourses := []course{}

	// We can validate json format before parsing
	isValid := json.Valid(bytejson)
	if !isValid {
		fmt.Println("Invalid JSON Format")
		return
	}

	// we can also create dummy json as following
	// var mybyteJson []byte =  []byte(
	// `[{"name":"GO Lang Course","price":322,"website":"udemy.com","tags":["web-dev","js"]},{"name":"React Js","price":322,"website":"coursera.com"}]`
	// )

	err := json.Unmarshal([]byte(bytejson), &myCourses)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Json Decoded object as %#v", myCourses)

}

func DecodeJsonAsMap() {

	/* Dictionary of <string,Interface> */
	/* Since a value can be any type: string, int , even another object */
	var myNewCourses []map[string]interface{}

	// myNewCourses["test"] = "test"
	// myNewCourses["test2"] = "test2"

	fmt.Println("My Courses ", myNewCourses)

	var mybyteJson []byte = []byte(
		`[{
		  "name":"GO Lang Course",
		  "price":322,
		  "website":"udemy.com",
		  "tags":["web-dev","js"]
		 },
		 {
		 "name":"React Js",
		  "price":322,
		  "website":"coursera.com"
		 }
		 ]`,
	)

	isValid := json.Valid(mybyteJson)
	if !isValid {
		fmt.Println("Json is not valid")
		return
	}

	err := json.Unmarshal(mybyteJson, &myNewCourses)

	if err != nil {
		panic(err)
	}

	for key, value := range myNewCourses {
		fmt.Printf("Key is %v , Value is %v\n", key, value)
		for _, v := range value {
			fmt.Printf("Value is %v\n", v)

		}
	}

}
