package main

import (
	"io/fs"
	"net/http"
	getEmbedded "website/ui"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler{

    standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

    mux := pat.New()
    mux.Get("/", http.HandlerFunc(app.home))
    mux.Get("/home", http.HandlerFunc(app.home))
    mux.Get("/home/", http.HandlerFunc(app.home))
    mux.Get("/services", http.HandlerFunc(app.services))
    mux.Get("/services/", http.HandlerFunc(app.services)) 
    mux.Get("/about-us", http.HandlerFunc(app.aboutus))
    mux.Get("/about-us/", http.HandlerFunc(app.aboutus)) 
    mux.Get("/contact-us", http.HandlerFunc(app.contactus))
    mux.Get("/contact-us/", http.HandlerFunc(app.contactus))
    mux.Post("/contact-us/create", http.HandlerFunc(app.requestCallBack))
    mux.Post("/contact-us/create/", http.HandlerFunc(app.requestCallBack)) 
    mux.Get("/reviews", http.HandlerFunc(app.reviewsPage))
    mux.Get("/reviews/", http.HandlerFunc(app.reviewsPage))
    mux.Post("/reviews/create", http.HandlerFunc(app.createReview))
    mux.Post("/reviews/create/", http.HandlerFunc(app.createReview)) 

    assets, err := fs.Sub(getEmbedded.GetEmbeddedStatic(), "static")
    if err != nil{
        app.errorLog.Println(err)
    }
    fileServer := http.FileServer(http.FS(assets))

    //fileServer := http.FileServer(http.Dir("./ui/static/"))
    //fileServer := http.FileServerFS(staticDir)
    mux.Get("/static/", http.StripPrefix("/static/", fileServer))
 
    return standardMiddleware.Then(mux)
}
