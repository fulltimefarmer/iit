package courseDatabase

var Courses []Course

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

func CreateCourses() {
	professor1 := Professor{Firstname: "Billy", Lastname: "P"}
	professor2 := Professor{Firstname: "Travis", Lastname: "S"}
	professor3 := Professor{Firstname: "Bob", Lastname: "S"}

	// feel free to make these more interesting :(
	Courses = append(Courses, Course{CRN: "1", Title: "Golang 101", Professor: &professor1, Description: "The best class ever 1"})
	Courses = append(Courses, Course{CRN: "2", Title: "Golang 102", Professor: &professor3, Description: "The best class ever 2"})
	Courses = append(Courses, Course{CRN: "3", Title: "Golang 103", Professor: &professor2, Description: "The best class ever 3"})
	Courses = append(Courses, Course{CRN: "4", Title: "Golang 104", Professor: &professor2, Description: "The best class ever 4"})
	Courses = append(Courses, Course{CRN: "5", Title: "Golang 105", Professor: &professor3, Description: "The best class ever 5"})
	Courses = append(Courses, Course{CRN: "6", Title: "Golang 106", Professor: &professor1, Description: "The best class ever 6"})
	Courses = append(Courses, Course{CRN: "7", Title: "Golang 107", Professor: &professor1, Description: "The best class ever 7"})
	Courses = append(Courses, Course{CRN: "8", Title: "Golang 108", Professor: &professor2, Description: "The best class ever 8"})
	Courses = append(Courses, Course{CRN: "9", Title: "Golang 109", Professor: &professor3, Description: "The best class ever 9"})
	Courses = append(Courses, Course{CRN: "10", Title: "Golang 110", Professor: &professor1, Description: "The best class ever 10"})

}
