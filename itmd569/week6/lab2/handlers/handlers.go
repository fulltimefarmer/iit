package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/tsmith-41/week6/courseDatabase"
	"net/http"
)

// GET
func GetCourseByCRN(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, course := range courseDatabase.Courses {
		if course.CRN == params["crn"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	http.NotFound(w, r)
}

// POST
func CreateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var course courseDatabase.Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	courseDatabase.Courses = append(courseDatabase.Courses, course)
	json.NewEncoder(w).Encode(course)
}

// GET
func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(courseDatabase.Courses)
}

// PUT
func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	// implement
}

// DELETE
func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")  // set application/jsonheader
	params := mux.Vars(r)                               // retrieve the parameters in the request
	for index, course := range courseDatabase.Courses { // loop through the courses database
		if course.CRN == params["crn"] { // check if the course CRN matches the request parameters
			courseDatabase.Courses = append(courseDatabase.Courses[:index], courseDatabase.Courses[index+1:]...) // remove that course from the database
			break
		}
	}
	json.NewEncoder(w).Encode(courseDatabase.Courses) // encode the entire database and send response back to requester
}
