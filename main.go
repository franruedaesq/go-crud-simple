package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type task struct {
	ID     int    `json:ID`
	Name   string `json:Name`
	Contet string `json:Contet`
}

type allTasks []task

var tasks = allTasks{
	{
		ID:     1,
		Name:   "Task 1",
		Contet: "Task 1 content",
	},
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(tasks)
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my API")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	http.ListenAndServe(":3000", router)
}
