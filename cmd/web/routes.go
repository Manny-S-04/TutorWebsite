package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler{

    standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

    // assign functions to routes here
    mux := pat.New()
    mux.Get("/", http.HandlerFunc(app.home)
    mux.Get("/maths", http.HandlerFunc()

 
    fileServer := http.FileServer(http.Dir("./ui/static/"))
    mux.Get("/static/", http.StripPrefix("/static",fileServer))
 
    return standardMiddleware.Then(mux)
}
