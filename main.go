package main

import (
	"database/sql"
	"log"
	"net/http"
)

var db *sql.DB

func init() {
	tmpDB, err := sql.Open("postgres", "dbname=students_database user=postgres password=JJaRnp5Ten host=postgresql-1575201111.default.svc.cluster.local sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	db = tmpDB
}

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("www/assets"))))

	http.HandleFunc("/", handleListStudents)
	http.HandleFunc("/student.html", handleViewStudent)
	http.HandleFunc("/save", handleSaveStudent)
	http.HandleFunc("/delete", handleDeleteStudent)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
