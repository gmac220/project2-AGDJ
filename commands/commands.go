package commands

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gmac220/project2-AGDJ/opendb"
)

// Users contains id, username, and password
type Users struct {
	id       int
	username string
	password string
}

// Containers have containername, port, and username
type Containers struct {
	containername string
	port          int
	username      string
}

// SignIn verifies if user credentials match database
func SignIn(username string) (int, string, string) {
	var usernamedb, passdb string
	var id int
	var row *sql.Row

	db := opendb.OpenDB()
	defer db.Close()
	row = db.QueryRow("SELECT * FROM users WHERE username = $1", username)
	row.Scan(&id, &usernamedb, &passdb)
	fmt.Println("Logged in with", usernamedb, passdb)
	return id, usernamedb, passdb
}

// CreateAccount for either a customer or employee
func CreateAccount(username string, password string) {
	db := opendb.OpenDB()
	defer db.Close()
	db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", username, password)
}

// QueryUser selects a specific user from the users table
func QueryUser(username string) (int, string, string) {
	var id int
	var uname, pw string

	db := opendb.OpenDB()
	defer db.Close()
	row := db.QueryRow("SELECT * FROM users WHERE username = $1", username)
	row.Scan(&id, &uname, &pw)

	return id, uname, pw
}

// QueryAllUsers of users table
func QueryAllUsers() []Users {
	var info []Users
	db := opendb.OpenDB()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var uname, pw string = "", ""
		var id int = 0
		rows.Scan(&id, &uname, &pw)
		info = append(info, Users{id: id, username: uname, password: pw})
	}
	return info
}

// DeleteUser deletes row from users table
func DeleteUser(username string) {
	db := opendb.OpenDB()
	defer db.Close()
	db.Exec("DELETE FROM users WHERE username = $1", username)
}

// CreateRunning creates a row in the running table
func CreateRunning(port int, username string) {
	db := opendb.OpenDB()
	defer db.Close()
	db.Exec("INSERT INTO running (port, username) VALUES ($1, $2)", port, username)
}

// QueryUser selects a specific user from the users table
func QueryRunning(username string) (string, int, string) {
	var cname, uname string
	var port int

	db := opendb.OpenDB()
	defer db.Close()
	row := db.QueryRow("SELECT * FROM running WHERE username = $1", username)
	row.Scan(&cname, &port, &uname)

	return cname, port, uname
}

// QueryAllRunning if username "" gets all containers in running table else gets all user's running containers
func QueryAllRunning(username string) []Containers {
	var containers []Containers
	var rows *sql.Rows
	var err error

	db := opendb.OpenDB()
	defer db.Close()

	if username == "" {
		rows, err = db.Query("SELECT * FROM running")
	} else {
		rows, err = db.Query("SELECT * FROM running WHERE username=$1", username)
	}

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var uname, cname string = "", ""
		var port int = 0
		rows.Scan(&cname, &port, &uname)
		containers = append(containers, Containers{containername: cname, port: port, username: uname})
	}
	return containers
}

// DeleteRunning deletes row from running table
func DeleteRunning(port int) {
	db := opendb.OpenDB()
	defer db.Close()
	db.Exec("DELETE FROM running WHERE port = $1", port)
}
