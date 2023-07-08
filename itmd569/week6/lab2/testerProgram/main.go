package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Course struct {
	CRN         string     `json:"crn"`
	Title       string     `json:"title"`
	Professor   *Professor `json:"professor"`
	Description string     `json:"description"`
}

type Professor struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname`
}

func main() {
	testPUTRequest()    // update course
	testDELETERequest() // delete course
	testPOSTRequest()   // create new course

}

func testPOSTRequest() {
	fmt.Println("\nTesting POST Request (Creating Class)")

	professor := Professor{
		Firstname: "New Professor FN",
		Lastname:  "New Professor LN",
	}

	course := Course{
		CRN:         "11",
		Title:       "Cool new Go Course",
		Professor:   &professor,
		Description: "Best course ever 11",
	}

	courseJson, err := json.Marshal(course)
	if err != nil {
		panic(err)
	}

	// before creating course
	resp, err := http.Get("http://localhost:8080/courses/11")
	if err != nil {
		panic(err)
	}

	var emptyCourse Course
	_ = json.NewDecoder(resp.Body).Decode(&emptyCourse)
	fmt.Println("Before Creating: ", emptyCourse)

	// creating course via POST request
	_, err = http.Post("http://localhost:8080/courses/11", "application/json", bytes.NewBuffer(courseJson))

	// after creating course
	resp, err = http.Get("http://localhost:8080/courses/11")
	if err != nil {
		panic(err)
	}

	var newCourse Course
	_ = json.NewDecoder(resp.Body).Decode(&newCourse)
	fmt.Println("After Creating: ", newCourse)

}

func testDELETERequest() {
	fmt.Println("\nTesting DELETE Request (Deleting Class)")

	client := &http.Client{}
	// before deleting
	resp, err := http.Get("http://localhost:8080/courses/3")
	var oldCourse0 Course
	_ = json.NewDecoder(resp.Body).Decode(&oldCourse0)
	fmt.Println("Before Delete: ", oldCourse0)

	// Create delete request
	req, err := http.NewRequest("DELETE", "http://localhost:8080/courses/3", nil)
	if err != nil {
		fmt.Println(err)

	}

	// Fetch DELETE Request
	resp, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// After deleting
	resp, err = http.Get("http://localhost:8080/courses/3")
	var oldCourse Course
	_ = json.NewDecoder(resp.Body).Decode(&oldCourse)
	fmt.Println("After Delete: ", oldCourse)

}

func testPUTRequest() {
	fmt.Println("\nTesting PUT Request (Updating Class)")
	client := &http.Client{}

	professor := Professor{
		Firstname: "New Professor FN",
		Lastname:  "New Professor LN",
	}

	course := Course{
		CRN:         "2",
		Title:       "Updated title for CRN 2",
		Professor:   &professor,
		Description: "updated description",
	}

	courseJson, err := json.Marshal(course)
	if err != nil {
		panic(err)
	}

	// before updating
	resp, err := http.Get("http://localhost:8080/courses/2")
	var oldCourse Course
	_ = json.NewDecoder(resp.Body).Decode(&oldCourse)
	fmt.Println("Before Update: ", oldCourse)

	// update the course with CRN 2
	req, err := http.NewRequest(http.MethodPut, "http://localhost:8080/courses/2", bytes.NewBuffer(courseJson))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}

	// after updating
	resp, err = http.Get("http://localhost:8080/courses/2")
	var updatedCourse Course
	_ = json.NewDecoder(resp.Body).Decode(&updatedCourse)
	fmt.Println("After Update: ", updatedCourse)
}
