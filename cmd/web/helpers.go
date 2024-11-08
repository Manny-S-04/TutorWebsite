package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/smtp"
	"runtime/debug"
)

func (app *application) addDefaultData(td *templateData, _ *http.Request) *templateData  {
    if td == nil{
        td = &templateData{}
    }
    return td
}

func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData)  {
    ts, ok := app.templateCache[name]
    if !ok{
        app.serverError(w,fmt.Errorf("The template %s does not exist",name))
        return
    }

    buf := new(bytes.Buffer)

    err  := ts.Execute(buf,app.addDefaultData(td,r))
    if err != nil{
        app.serverError(w,err)
    }

    buf.WriteTo(w)
}
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

    http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *application) emailService(body [3]string){
    name := body[0]
    review := body[1]
    stars := body[2]

    from := "mannybqckup@gmail.com"
    pass := "azew syaq goxr kwnk"
    to := "mannybqckup@gmail.com"

    msg := "\n" +
    "Subject: New Review" + "\n" +
    "Name: " + name + "\n" +
    "Review: " + review + "\n" +
    "Stars: " + stars + "\n"
    
    err := smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", from, pass, "smtp.gmail.com"), from, []string{to}, []byte(msg))

    if err != nil{
        app.errorLog.Fatal("SMTP failure", err)
        return
    }
}

func (app *application) callbackService(body [2]string){
    name := body[0]
    message := body[1]

    from := "mannybqckup@gmail.com"
    pass := "azew syaq goxr kwnk"
    to := "mannybqckup@gmail.com"

    msg := "\n" +
    "Subject: New Message" + "\n" +
    "Name: " + name + "\n" +
    "Message: " + message + "\n"
    
    err := smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", from, pass, "smtp.gmail.com"), from, []string{to}, []byte(msg))

    if err != nil{
        app.errorLog.Fatal("SMTP failure", err)
        return
    }
}
