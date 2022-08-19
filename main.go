package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Student struct {
	StudentId string `json:"studentId"`
	Name      string `json:"Name"`
	Age       int    `json:"Age"`
	Class     string `json:"Class"`
	Subject   string `json:"Subject"`
}

var Students []Student

func main() {
	Students = []Student{
		{StudentId: "1", Name: "idk", Age: 6, Class: "nursery", Subject: "Poem"},
		{StudentId: "2", Name: "wth", Age: 8, Class: "one", Subject: "Maths"},
	}
	handleRequests()
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "homepage!")
	fmt.Println("Endpoint hit1")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homepage)
	myRouter.HandleFunc("/all", returnAllStudents)
	myRouter.HandleFunc("/student/{studentId}", returnSingleStudent)
	myRouter.HandleFunc("/student", addNewStudent).Methods("POST")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func addNewStudent(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	//fmt.Fprintf(w, "%+v", string(reqBody))
	var student Student
	json.Unmarshal(reqBody, &student)
	Students = append(Students, student)
	json.NewEncoder(w).Encode(student)
	fmt.Println("Endpoint hit2")
	if err != nil {
		fmt.Println("Could not add student")
	}
}

func returnAllStudents(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit3")
	json.NewEncoder(w).Encode(Students)
}

func returnSingleStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit4")
	vars := mux.Vars(r)
	key := vars["studentId"]
	//fmt.Fprintf(w, "Key: "+key)

	for _, student := range Students {
		if student.StudentId == key {
			json.NewEncoder(w).Encode(student)
		}
	}
}
