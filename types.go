package main

import "time"

//IndexPage represents the content of the index page, available on "/"
//The index page shows a list of all rivers stored on db
type IndexPage struct {
	AllRivers []River
}

//Riverage represents the content of the river page, available on "/river.html"
//The river page shows info about a given river
type RiverPage struct {
	TargetRiver River
}

//River represents a river object
type River struct {
	ID              int
	Name            string
	City          string
	PublicationDate time.Time
	Level           int
}

//PublicationDateStr returns a sanitized Publication Date in the format YYYY-MM-DD
func (b River) PublicationDateStr() string {
	return b.PublicationDate.Format("2006-01-02")
}

//ErrorPage represents shows an error message, available on "/river.html"
type ErrorPage struct {
	ErrorMsg string
}
