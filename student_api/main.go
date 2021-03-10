package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

//Student struct
type Student struct {
	Name   string `json:"name"`
	Rollno string `json:rollno.`
	Marks  *Mark  `json:marks`
}

//Mark struct
type Mark struct {
	Maths     string `json:maths`
	Chemistry string `json:chemistry`
	Physics   string `json:physics`
}

//Init student var as a slice Student struct
var students []Student

//Get all students
func getStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

func main() {

	//Init new router
	r := mux.NewRouter()
	students = append(students, Student{Name: "John Abraham", Rollno: "1111", Marks: &Mark{Maths: "95", Chemistry: "75", Physics: "80"}})
	r.HandleFunc("/api/students", getStudents).Methods("GET")
	http.ListenAndServe(":8000", r)
}
