package main

import (
	"net/http"
)

func main() {
	getDB()
	initTables()
	http.HandleFunc("/api/", router)
	http.ListenAndServe(":8666", nil)
}
