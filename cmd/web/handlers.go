package main

import (
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
    app.render(w,r,"home.page.tmpl", &templateData{})
}


func (app *application) services(w http.ResponseWriter, r *http.Request){
    app.render(w,r,"services.page.tmpl", &templateData{})
}

func (app *application) reviewsPage(w http.ResponseWriter, r *http.Request){
    app.render(w,r,"reviews.page.tmpl", &templateData{})
}


func (app *application) aboutus(w http.ResponseWriter, r *http.Request){
    app.render(w,r,"aboutus.page.tmpl", &templateData{})
}
