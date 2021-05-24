package main

import (
	"fmt"
	"net/http"
	"strings"
)

func getID(r *http.Request) string {
	pathSlice := strings.Split(r.URL.Path, "/")
	id := pathSlice[len(pathSlice)-1]
	return id
}

/*
CLIENTS CONTROLLER
*/
func getClients(w http.ResponseWriter, r *http.Request) {
	var clients []Client
	rows := getAll("clients")

	for rows.Next() {
		var c Client
		err := rows.Scan(&c.ID, &c.Name, &c.Email, &c.CreatedAt, &c.UpdatedAt, &c.DeletedAt)

		if err != nil {
			fmt.Println(err)
		}
		clients = append(clients, c)

	}

	fmt.Println(clients)
}
func getClient(w http.ResponseWriter, r *http.Request) {
	var c Client
	id := getID(r)
	// fmt.Println(id)
	row := getOne("clients", id)

	err := row.Scan(&c.ID, &c.Name, &c.Email, &c.CreatedAt, &c.UpdatedAt, &c.DeletedAt)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(c)
}

func addClient(w http.ResponseWriter, r *http.Request) {
	lastInsertID := addOne("clients", r)
	fmt.Println(lastInsertID)
}
func updateClient(w http.ResponseWriter, r *http.Request) {
	id := getID(r)
	updateOne("clients", id, r)

}
func deleteClient(w http.ResponseWriter, r *http.Request) {
	id := getID(r)
	deleteOne("clients", id)
}

/*
PROJECTS CONTROLLER
*/
func getProjects(w http.ResponseWriter, r *http.Request) {

}
func getProject(w http.ResponseWriter, r *http.Request) {

}

func addProject(w http.ResponseWriter, r *http.Request) {

}
func updateProject(w http.ResponseWriter, r *http.Request) {

}
func deleteProject(w http.ResponseWriter, r *http.Request) {

}

/*
TASKS CONTROLLER
*/
func getTasks(w http.ResponseWriter, r *http.Request) {

}
func getTask(w http.ResponseWriter, r *http.Request) {

}

func addTask(w http.ResponseWriter, r *http.Request) {

}
func updateTask(w http.ResponseWriter, r *http.Request) {

}
func deleteTask(w http.ResponseWriter, r *http.Request) {

}

/*
NOTES CONTROLLER
*/
func getNotes(w http.ResponseWriter, r *http.Request) {

}
func getNote(w http.ResponseWriter, r *http.Request) {

}

func addNote(w http.ResponseWriter, r *http.Request) {

}
func updateNote(w http.ResponseWriter, r *http.Request) {

}
func deleteNote(w http.ResponseWriter, r *http.Request) {

}
