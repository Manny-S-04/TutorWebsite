package main

import (
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
    app.render(w,r,"home.page.tmpl", &templateData{})
}

func (app *application) contactus(w http.ResponseWriter, r *http.Request){
    app.render(w,r,"contactus.page.tmpl", &templateData{})
}

func (app *application) services(w http.ResponseWriter, r *http.Request){
    app.render(w,r,"services.page.tmpl", &templateData{})
}

func (app *application) reviewsPage(w http.ResponseWriter, r *http.Request){
    reviews, err  := app.reviews.GetAll()
    if err != nil{
        app.serverError(w,err)
    }
    app.render(w,r,"reviews.page.tmpl", &templateData{Reviews: reviews})
}

func (app *application) aboutus(w http.ResponseWriter, r *http.Request){
    app.render(w,r,"aboutus.page.tmpl", &templateData{})
}

func (app *application) requestCallBack(w http.ResponseWriter, r * http.Request){

    err := r.ParseForm()
    if err != nil{
        app.clientError(w, http.StatusBadRequest)
        return
    }

    name := r.PostForm.Get("name")
    message := r.PostForm.Get("message")
    errors := make(map[string]string)

    if strings.TrimSpace(name) == ""{
        errors["name"] = "Name cannot be empty"
    } else if utf8.RuneCountInString(name) > 100 {
        errors["name"] = "Name is too long"
    }

    if strings.TrimSpace(message) == ""{
        errors["message"] = "Content cannot be empty"
    } else if utf8.RuneCountInString(name) > 350 {
        errors["message"] = "Content is too long"
    }

    if len(errors) > 0{
        app.render(w,r,"contactus.page.tmpl", &templateData{
            FormErrors: errors,
            FormData: r.PostForm,
        })
        return
    }
    
    sendMessage := [2]string{name, message}

    app.callbackService(sendMessage)

    http.Redirect(w,r,"/contact-us", http.StatusSeeOther)
}

func (app *application) createReview(w http.ResponseWriter, r * http.Request){
    err := r.ParseForm()
    if err != nil{
        app.clientError(w, http.StatusBadRequest)
        return
    }

    name := r.PostForm.Get("name")
    reviewContent := r.PostForm.Get("review-content")
    stars := r.PostForm.Get("stars")
    errors := make(map[string]string)

    if strings.TrimSpace(name) == ""{
        errors["name"] = "Name cannot be empty"
    } else if utf8.RuneCountInString(name) > 100 {
        errors["name"] = "Name is too long"
    }

    if strings.TrimSpace(reviewContent) == ""{
        errors["reviewContent"] = "Content cannot be empty"
    } else if utf8.RuneCountInString(name) > 350 {
        errors["reviewContent"] = "Content is too long"
    }


    starsInt, err := strconv.Atoi(stars)
    if err != nil{
        app.infoLog.Print("Failed to convert stars", err)
    }
    if starsInt <= 0 || starsInt >= 6{
        errors["stars"] = "Stars must be between 1 and 5"
    } 
    
    if len(errors) > 0{
        reviews, err  := app.reviews.GetAll()
        if err != nil{
            app.serverError(w,err)
        }
        app.render(w,r,"reviews.page.tmpl", &templateData{
            Reviews: reviews,
            FormErrors: errors,
            FormData: r.PostForm,
        })
        return
    }
    
    sendReview := [3]string{name, reviewContent, stars}

    app.emailService(sendReview)

    http.Redirect(w,r,"/reviews", http.StatusSeeOther)
}
