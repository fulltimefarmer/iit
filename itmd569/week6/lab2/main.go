package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/tsmith-41/week6/courseDatabase"
	"github.com/tsmith-41/week6/handlers"
)

func main() {

	courseDatabase.CreateCourses() // creates the array of courses, the "database"

	// setup your mux router and routes here
	router := mux.NewRouter()
	router.HandleFunc("/courses/{crn}", handlers.GetCourseByCRN).Methods("GET")  // implement
	router.HandleFunc("/courses/", handlers.GetAllCourses).Methods("GET")        // done already
	router.HandleFunc("/courses/{crn}", handlers.UpdateCourse).Methods("PUT")    // implement
	router.HandleFunc("/courses/{crn}", handlers.CreateCourse).Methods("POST")   // implement
	router.HandleFunc("/courses/{crn}", handlers.DeleteCourse).Methods("DELETE") // done already
	http.ListenAndServe(":8080", router)

}
