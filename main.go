package main

import (
	"database/sql"
	"log"
	"net/http"
)

var db *sql.DB

func init() {
	tmpDB, err := sql.Open("postgres", "dbname=rivers_database user=postgres password=postgres host=localhost sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	db = tmpDB
}

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("www/assets"))))

	http.HandleFunc("/", handleListRivers)
	http.HandleFunc("/river.html", handleViewRiver)
	http.HandleFunc("/save", handleSaveRiver)
	http.HandleFunc("/delete", handleDeleteRiver)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
