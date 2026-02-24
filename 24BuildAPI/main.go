package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Author *Author `json:"author"`
}

type Author struct {
	Name     string `json:"name"`
	Platform string `json:"platform"`
}

// Course Methods
// Always use pointers for updating real values and not copies
func (c *Course) IsEmpty() bool {
	return c.Id == "" && c.Name == ""
}

// fake DB
// global variable
var Courses []Course

func main() {

	// course1 :=  {"1", "React Js", 123.32, &Author{"Hiresh", "Youtube"},}

	Courses = append(Courses, Course{"1", "React Js", 123.32, &Author{"Hiresh", "Youtube"}})
	Courses = append(Courses, Course{"2", "Go lang", 123.32, &Author{"Hiresh", "Youtube"}})
	Courses = append(Courses, Course{"3", "Javascript", 123.32, &Author{"Hiresh", "Youtube"}})
	Courses = append(Courses, Course{"4", "Python", 123.32, &Author{"Hiresh", "Youtube"}})

	r := mux.NewRouter()
	/*
		Http.ResponseWriter
		An interface, Implementing io.Writer, has 3 important methods :
		Write , // writes response
		Header, // retrieves header
		WriteHeader // writes header

	*/
	r.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {

		// The header method returns
		w.Header().Set("content-Type", "text/html")
		w.Header().Set("AtEnd1", "value 1")
		w.Header().Set("AtEnd2", "value 2")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("<h1> Welcome to My First API </h1>"))
	})

	//
	r.HandleFunc("/courses", GetAllCourses).Methods("GET")

	r.HandleFunc("/courses/{Id}", GetCourseById).Methods("GET")

	r.HandleFunc("/courses", CreateCourse).Methods("POST")

	r.HandleFunc("/courses/{Id}", UpdateCourse).Methods("PUT")

	r.HandleFunc("/courses/{Id}", DeleteCourse).Methods("DELETE")

	http.ListenAndServe("localhost:4000", r)

}

func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Courses)
}

func GetCourseById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["Id"]

	for _, value := range Courses {
		if value.Id == id {
			// json.NewEncoder(w).Encode(value)
			// Decode/ENcode used for STREAMs like http request body, response body, file stream, etc.
			// unlike marshal/unmarshal which is used for in-memory data
			json.NewEncoder(w).Encode(value)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	// json.NewEncoder(w).Encode("No Data for the given Id")
}

func CreateCourse(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var course Course

	/*
		Remember this issue
		when reading a stream twice, the second time it will be at EOF
		conclusion, any read operation whether in network stream,
		memory stream, file stream, etc. will move the pointer to the end of the stream
		so the second read will be empty and at EOF
	*/
	// content, _ := io.ReadAll(r.Body)
	// fmt.Println(string(content))
	// fmt.Println(json.Valid(content))

	// Decode/ENcode used for STREAMs like http request body, response body, file stream, etc.
	// unlike marshal/unmarshal which is used for in-memory data
	json.NewDecoder(r.Body).Decode(&course)
	// json.Unmarshal(content, &course)

	// Check for duplication before Insertion
	for _, value := range Courses {
		if value.Name == course.Name {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode("Course Name Already Exists")
			return
		}
	}

	json.NewEncoder(w).Encode(course)

	rand.NewSource(time.Now().Unix())
	newId := strconv.Itoa(rand.Intn(100))
	course.Id = newId
	Courses = append(Courses, course)
	fmt.Println("Courses: ", Courses)

}

func UpdateCourse(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Received PUT Request")
	params := mux.Vars(r)

	for index, value := range Courses {
		if value.Id == params["Id"] {
			// remove target element
			Courses = append(Courses[:index], Courses[index+1:]...)

			var newElement Course
			json.NewDecoder(r.Body).Decode(&newElement)
			newElement.Id = params["Id"]
			Courses = append(Courses, newElement)
			json.NewEncoder(w).Encode(newElement)
			fmt.Println("Courses : ", Courses)
			return
		}
	}

	// If no ID found
	json.NewEncoder(w).Encode("Not a Valid Id")

}

func DeleteCourse(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for k, v := range Courses {
		if v.Id == params["Id"] {
			Courses = append(Courses[:k], Courses[k+1:]...)
			json.NewEncoder(w).Encode(Courses)
			return
		}
	}

	json.NewEncoder(w).Encode("Not Valid Id")
}
