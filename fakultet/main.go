package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Counter struct {
	Id    int `json:"Id"`
	Count int `json:"Count"`
}

func healthz(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "All good")
}

func initialize(w http.ResponseWriter, req *http.Request) {
	user := os.Getenv("user")
	pass := os.Getenv("password")
	host := os.Getenv("db_host")
	db, err := sql.Open("mysql", user+":"+pass+"@tcp("+host+":3306)/")
	if err != nil {
		fmt.Fprintln(w, err.Error())
	} else {
		_, err := db.Exec("CREATE DATABASE if not exists golang;")
		if err != nil {
			fmt.Fprintln(w, err.Error())
		}
		_, err = db.Exec("USE golang;")
		if err != nil {
			fmt.Fprintln(w, err.Error())
		}
		_, err = db.Exec("CREATE TABLE counter(Id int NOT NULL AUTO_INCREMENT,Count int,PRIMARY KEY(id));")
		if err != nil {
			fmt.Fprintln(w, err.Error())
		}

		defer db.Close()

		_, err = db.Exec("INSERT INTO golang.counter values(1,0);")
		if err != nil {
			fmt.Fprintln(w, err.Error())
		}
	}
	fmt.Fprintln(w, "Success!")
}

func query(w http.ResponseWriter, req *http.Request) {
	var counter Counter
	user := os.Getenv("user")
	pass := os.Getenv("password")
	host := os.Getenv("db_host")
	go_host := os.Getenv("go_host")
	fmt.Println(user + pass + host)
	db, err := sql.Open("mysql", user+":"+pass+"@tcp("+host+":3306)/")
	if err != nil {
		fmt.Fprintln(w, err.Error())
	} else {
		stmt, err := db.Prepare("SELECT Id,Count FROM golang.counter WHERE id = ?;")
		if err != nil {
			fmt.Fprintln(w, err.Error())
		}
		err = stmt.QueryRow(1).Scan(&counter.Id, &counter.Count)
		if err != nil {
			fmt.Fprintln(w, err.Error())
		}
		stmt2, err2 := db.Prepare("UPDATE golang.counter SET Count = ? WHERE Id = ?;")
		if err2 != nil {
			fmt.Fprintln(w, err.Error())
		}
		counter.Count = counter.Count + 1
		_, err = stmt2.Exec(counter.Count, 1)
		if err != nil {
			fmt.Fprintln(w, err.Error())
		}

		defer db.Close()
		defer stmt.Close()

		fmt.Fprintf(w, "Counter: %d\n", counter.Count)
		fmt.Fprintf(w, "Hello from %s\n", string(go_host))

	}
}

func main() {
	http.HandleFunc("/hz", healthz)
	http.HandleFunc("/init", initialize)
	http.HandleFunc("/query", query)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
