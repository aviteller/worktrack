package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func initTables() {

	tasks := `CREATE TABLE tasks
	 (
		 id INTEGER PRIMARY KEY AUTOINCREMENT,
		 project_id INTEGER,
		 name TEXT,
		 start_datetime TEXT,
		 end_datetime TEXT,
		 last_pause_datetime TEXT,
		 complete INTEGER,
		 currently_active TEXT,
		 created_at TEXT,
		 updated_at TEXT,
		 deleted_at TEXT
		 )`

	getDB().Exec(tasks)
	projects := `CREATE TABLE projects
	(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		client_id INTEGER,
		name TEXT,
		pph TEXT,
			complete INTEGER,
		 created_at TEXT,
		 updated_at TEXT,
		 deleted_at TEXT
		 )`

	getDB().Exec(projects)
	clients := `CREATE TABLE clients
	(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		email TEXT,
		 created_at TEXT,
		 updated_at TEXT,
		 deleted_at TEXT
		 )`

	getDB().Exec(clients)
	notes := `CREATE TABLE notes
	(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		table_id INTEGER,
		table_name text,
		content text,
		 created_at TEXT,
		 updated_at TEXT,
		 deleted_at TEXT
		 )`

	getDB().Exec(notes)

}

func getNow() string {
	now := time.Now().UTC().Format("2006-01-02 03:04:05")
	return now
}

func getDB() *sql.DB {
	database, err := sql.Open("sqlite3", "./timer.db")
	if err != nil {
		fmt.Println(err)
	}
	return database
}

func getAll(table string) *sql.Rows {
	queryString := "SELECT * FROM " + table + " WHERE deleted_at IS NULL"

	rows, err := getDB().Query(queryString)

	if err != nil {
		panic(err)
	}

	return rows
}

func getOne(table string, id string) *sql.Row {
	queryString := "SELECT * FROM " + table + " WHERE id = $1"

	fmt.Println(queryString)

	row := getDB().QueryRow(queryString, id)

	return row

}

func addOne(table string, r *http.Request) int64 {
	var gi map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&gi)

	if err != nil {
		panic(err)
	}

	keys := make([]string, 0, len(gi))
	values := make([]interface{}, 0, len(gi))
	// var table interface{}
	for k, v := range gi {

		keys = append(keys, k)
		values = append(values, v)

	}
	// now := time.Now().UTC().Format("2006-01-02 03:04:05")
	keys = append(keys, "created_at")
	values = append(values, getNow())

	s := make([]string, len(values))
	for i, v := range values {
		s[i] = fmt.Sprint(v)
	}

	sqlStatment := "INSERT INTO " + table + "(" + strings.Join(keys, ",") + ") VALUES ('" + strings.Join(s, "','") + "')"

	stmt, err := getDB().Prepare(sqlStatment)

	if err != nil {
		panic(err)
	}
	res, err := stmt.Exec()

	if err != nil {
		panic(err)
	}

	lastInsertID, err := res.LastInsertId()

	if err != nil {
		panic(err)
	}
	return lastInsertID
}

func updateOne(table string, id string, r *http.Request) {
	var gi map[string]string
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&gi)

	if err != nil {
		panic(err)
	}
	var updateString string

	for k, v := range gi {
		updateString += k + " = '" + v + "' , "
	}
	// now := time.Now().UTC().Format("2006-01-02 03:04:05")

	updateString += " updated_at = '" + getNow() + "'"

	// updateString = strings.TrimSuffix(updateString, ", ")
	fmt.Println(updateString)

	sqlStmt := "UPDATE " + table + " SET " + updateString + " WHERE id = " + id

	stmt, err := getDB().Prepare(sqlStmt)

	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec()

	if err != nil {
		panic(err)
	}

}

func deleteOne(table string, id string) {
	// use function in future
	var checkId int64
	var deletedAt sql.NullString

	getOneQuery := "SELECT id,deleted_at FROM " + table + " WHERE id = " + id

	row := getDB().QueryRow(getOneQuery)

	err := row.Scan(&checkId, &deletedAt)

	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(checkId, deletedAt.Valid, deletedAt.String)

	if checkId != 0 {
		if deletedAt.Valid == true {
			// fmt.Println("perm delete")
			deletedStmt := "DELETE FROM " + table + " WHERE id = " + id
			getDB().Exec(deletedStmt)
		} else {
			// fmt.Println("update")
			updateStmt := "UPDATE " + table + " SET deleted_at = '" + getNow() + "' WHERE id = " + id
			fmt.Println(updateStmt)
			getDB().Exec(updateStmt)

		}
	}

}
