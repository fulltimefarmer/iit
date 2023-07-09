package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
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

func TestGetHandler(t *testing.T) {
	fmt.Println("\nTesting GET Request (Get Class by CRN)")

	resp, err := http.Get("http://localhost:8080/courses/1")
	if err != nil {
		t.Errorf("Error in GET request: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.Body == nil {
		t.Error("Response body is nil")
		return
	}

	var course Course
	err = json.NewDecoder(resp.Body).Decode(&course)
	if err != nil {
		t.Errorf("Error in JSON decoding: %v", err)
		return
	}

	fmt.Println("Get the Course: ", course)

	if &course == nil {
		t.Errorf("Response of http://localhost:8080/courses/1 is nil!")
	}
	if course.CRN != "1" {
		t.Errorf("CRN of http://localhost:8080/courses/1 doesn’t match expectation!")
	}
}

func TestPOSTRequest(t *testing.T) {
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
	if err != nil {
		t.Fatal(err)
	}

	// after creating course
	resp, err = http.Get("http://localhost:8080/courses/11")
	if err != nil {
		t.Fatal(err)
	}

	var newCourse Course
	_ = json.NewDecoder(resp.Body).Decode(&newCourse)

	fmt.Println("After Creating: ", newCourse)
	if newCourse.CRN != course.CRN {
		t.Errorf("Unexpected CRN. Got %s, want %s", newCourse.CRN, course.CRN)
	}
}

func TestDELETERequest(t *testing.T) {
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
		t.Fatal(err)
	}

	// Fetch DELETE Request
	resp, err = client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// After deleting
	resp, err = http.Get("http://localhost:8080/courses/3")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	var oldCourse Course
	_ = json.NewDecoder(resp.Body).Decode(&oldCourse)

	fmt.Println("After Delete: ", oldCourse)
	if oldCourse.CRN != "" {
		t.Errorf("Course not deleted successfully. CRN still exists: %s", oldCourse.CRN)
	}
}

func TestPUTRequest(t *testing.T) {
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
		t.Fatal(err)
	}

	// before updating
	resp, err := http.Get("http://localhost:8080/courses/2")
	var oldCourse Course
	_ = json.NewDecoder(resp.Body).Decode(&oldCourse)
	fmt.Println("Before Update: ", oldCourse)

	// update the course with CRN 2
	req, err := http.NewRequest(http.MethodPut, "http://localhost:8080/courses/2", bytes.NewBuffer(courseJson))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err = client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	// after updating
	resp, err = http.Get("http://localhost:8080/courses/2")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	var updatedCourse Course
	_ = json.NewDecoder(resp.Body).Decode(&updatedCourse)
	fmt.Println("After Update: ", updatedCourse)

	if updatedCourse.Title != course.Title {
		t.Errorf("Course title not updated successfully. Got %s, want %s", updatedCourse.Title, course.Title)
	}
}
