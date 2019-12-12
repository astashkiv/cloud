package main

//IndexPage represents the content of the index page, available on "/"
//The index page shows a list of all Rivers stored on db
type IndexPage struct {
	AllRivers []River
}

//RiverPage represents the content of the River page, available on "/River.html"
//The River page shows info about a given River
type RiverPage struct {
	TargetRiver River
}

//River represents a River object
type River struct {
	ID     int
	Name   string
	City string
	Level int
	Date string
}

//ErrorPage represents shows an error message, available on "/River.html"
type ErrorPage struct {
	ErrorMsg string
}
