package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/tsmith-41/week6/courseDatabase"
	"net/http"
)

// GET
func GetCourseByCRN(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // set the application/json header
	params := mux.Vars(r) // retrieve the parameters from the request
	for _, course := range courseDatabase.Courses { // loop the courses
		if course.CRN == params["crn"] { // find out the same course
			json.NewEncoder(w).Encode(course) // write the response
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with CRN: " + params["crn"])
}

// POST
func CreateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // set the application/json header
	var course courseDatabase.Course // create new course instance
	_ = json.NewDecoder(r.Body).Decode(&course) // decode the response body into the course instance
	courseDatabase.Courses = append(courseDatabase.Courses, course) // append the new course to courseDatabase.Courses
	json.NewEncoder(w).Encode(courseDatabase.Courses) // write the response
}

// GET
func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(courseDatabase.Courses)
}

// PUT
func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // set the application/json header
	params := mux.Vars(r) // retrieve the parameters from the request
	for index, course := range courseDatabase.Courses { // loop the courses
		if course.CRN == params["crn"] { // find out the same course
			courseDatabase.Courses = append(courseDatabase.Courses[:index], courseDatabase.Courses[index+1:]...) // remove the course
			var updatedCourse courseDatabase.Course // create new course instance
			_ = json.NewDecoder(r.Body).Decode(&updatedCourse) // set param to the new instance
			updatedCourse.CRN = params["crn"]
			courseDatabase.Courses = append(courseDatabase.Courses, updatedCourse) // add new course instance into Courses
			json.NewEncoder(w).Encode(updatedCourse) // write the response
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with CRN: " + params["crn"])
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
