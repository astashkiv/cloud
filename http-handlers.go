package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

func handleSaveRiver(w http.ResponseWriter, r *http.Request) {
	var id = 0
	var err error

	r.ParseForm()
	params := r.PostForm
	idStr := params.Get("id")

	if len(idStr) > 0 {
		id, err = strconv.Atoi(idStr)
		if err != nil {
			renderErrorPage(w, err)
			return
		}
	}

	name := params.Get("name")
	citys := params.Get("citys")

	levelStr := params.Get("level")
	level := 0
	if len(levelStr) > 0 {
		level, err = strconv.Atoi(levelStr)
		if err != nil {
			renderErrorPage(w, err)
			return
		}
	}

	dateStr := params.Get("date")
	date := 0
	if len(dateStr) > 0 {
		date, err = strconv.Atoi(dateStr)
		if err != nil {
			renderErrorPage(w, err)
			return
		}
	}

	if id == 0 {
		_, err = insertRiver(name, citys, level, date)
	} else {
		_, err = updateRiver(id, name, citys, level, date)
	}

	if err != nil {
		renderErrorPage(w, err)
		return
	}

	http.Redirect(w, r, "/", 302)
}

func handleListRivers(w http.ResponseWriter, r *http.Request) {
	rivers, err := allRivers()
	if err != nil {
		renderErrorPage(w, err)
		return
	}

	buf, err := ioutil.ReadFile("www/index.html")
	if err != nil {
		renderErrorPage(w, err)
		return
	}

	var page = IndexPage{AllRivers: rivers}
	indexPage := string(buf)
	t := template.Must(template.New("indexPage").Parse(indexPage))
	t.Execute(w, page)
}

func handleViewRiver(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	idStr := params.Get("id")

	var currentRiver = River{}
	//currentRiver.Arrear = time.Now()

	if len(idStr) > 0 {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			renderErrorPage(w, err)
			return
		}

		currentRiver, err = getRiver(id)
		if err != nil {
			renderErrorPage(w, err)
			return
		}
	}

	buf, err := ioutil.ReadFile("www/river.html")
	if err != nil {
		renderErrorPage(w, err)
		return
	}

	var page = RiverPage{TargetRiver: currentRiver}
	riverPage := string(buf)
	t := template.Must(template.New("riverPage").Parse(riverPage))
	err = t.Execute(w, page)
	if err != nil {
		renderErrorPage(w, err)
		return
	}
}

func handleDeleteRiver(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	idStr := params.Get("id")

	if len(idStr) > 0 {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			renderErrorPage(w, err)
			return
		}

		n, err := removeRiver(id)
		if err != nil {
			renderErrorPage(w, err)
			return
		}

		fmt.Printf("Rows removed: %v\n", n)
	}
	http.Redirect(w, r, "/", 302)
}

func renderErrorPage(w http.ResponseWriter, errorMsg error) {
	buf, err := ioutil.ReadFile("www/error.html")
	if err != nil {
		log.Printf("%v\n", err)
		fmt.Fprintf(w, "%v\n", err)
		return
	}

	var page = ErrorPage{ErrorMsg: errorMsg.Error()}
	errorPage := string(buf)
	t := template.Must(template.New("errorPage").Parse(errorPage))
	t.Execute(w, page)
}
