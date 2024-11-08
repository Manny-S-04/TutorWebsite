package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler{

    standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

    mux := pat.New()
    mux.Get("/", http.HandlerFunc(app.home))
    mux.Get("/home", http.HandlerFunc(app.home))
    mux.Get("/services", http.HandlerFunc(app.services))
    mux.Get("/about-us", http.HandlerFunc(app.aboutus))
    mux.Get("/contact-us", http.HandlerFunc(app.contactus))
    mux.Post("/contact-us", http.HandlerFunc(app.requestCallBack))
    mux.Get("/reviews", http.HandlerFunc(app.reviewsPage))
    mux.Post("/reviews/create", http.HandlerFunc(app.createReview))
 
    fileServer := http.FileServer(http.Dir("./ui/static/"))
    mux.Get("/static/", http.StripPrefix("/static",fileServer))
 
    return standardMiddleware.Then(mux)
}
