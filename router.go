package main

import (
	"fmt"
	"net/http"
	"strings"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello")
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")

}

func router(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	method := r.Method
	path := r.URL.Path
	path = strings.Replace(path, "/api/", "", -1)
	r.URL.Path = path
	pathSlice := strings.Split(path, "/")
	switch method {
	case "GET":
		switch pathSlice[0] {
		case "clients":
			if len(pathSlice) > 1 {
				getClient(w, r)
			} else {
				getClients(w, r)
			}
		case "projects":
			if len(pathSlice) > 1 {
				getProject(w, r)
			} else {
				getProjects(w, r)
			}
		case "tasks":
			if len(pathSlice) > 1 {
				getTask(w, r)
			} else {
				getTasks(w, r)
			}
		case "notes":
			if len(pathSlice) > 1 {
				getNote(w, r)
			} else {
				getNotes(w, r)
			}
		default:
			fmt.Fprintf(w, "GET METHOD: %s DOES NOT EXIST", path)
		}

	case "POST":
		switch pathSlice[0] {
		case "clients":
			addClient(w, r)
		case "projects":
			addProject(w, r)
		case "tasks":
			addTask(w, r)
		case "notes":
			addNote(w, r)
		default:
			fmt.Fprintf(w, "POST METHOD: %s DOES NOT EXIST", path)
		}
	case "DELETE":
		switch pathSlice[0] {
		case "clients":
			deleteClient(w, r)
		case "projects":
			deleteProject(w, r)
		case "tasks":
			deleteTask(w, r)
		case "notes":
			deleteNote(w, r)
		default:
			fmt.Fprintf(w, "POST METHOD: %s DOES NOT EXIST", path)
		}
	case "PUT":

		switch pathSlice[0] {
		case "clients":
			updateClient(w, r)
		case "projects":
			updateProject(w, r)
		case "tasks":
			updateTask(w, r)
		case "notes":
			updateNote(w, r)

		default:
			fmt.Fprintf(w, "POST METHOD: %s DOES NOT EXIST", path)
		}

	default:
		fmt.Fprintf(w, "METHOD: %s IS NOT ALLOWED", method)
	}

}
