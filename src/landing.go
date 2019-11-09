package main

import (
	"html/template"
	"net/http"
)

func usrLandingGET(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("web/landing.html"))
	_ = t.Execute(w, "")

}

func UserLandingHandler(w http.ResponseWriter, r *http.Request) {

	push(w, "css/main.css")
	push(w, "css/util.css")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	usrLandingGET(w,r)

}