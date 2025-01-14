package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/smtp"
	"runtime/debug"
	"strings"
    _ "embed"
)

//go:embed .env
var env string

func (app *application) getEnv () [2]string{
    lines := strings.Split(env, "\n")
    var envVars [2]string
    envVars[0] = lines[0] 
    envVars[1] = lines[1]

	return envVars
}

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

func (app *application) emailService(body [3]string) bool{
    name := body[0]
    review := body[1]
    stars := body[2]

    var envVars [2]string = app.getEnv()

    splitEmail := strings.Split(envVars[0], "=")
    email := splitEmail[1]
    splitPass := strings.Split(envVars[1], "=")

    from := email
    pass := splitPass[1]
    to := email

    msg := "\n" +
    "Subject: New Review" + "\n" +
    "Name: " + name + "\n" +
    "Review: " + review + "\n" +
    "Stars: " + stars + "\n"
    
    err := smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", from, pass, "smtp.gmail.com"), from, []string{to}, []byte(msg))

    if err != nil{
        app.errorLog.Fatal("SMTP failure", err)
        return false
    }

    return true
}

func (app *application) callbackService(body [2]string) bool{
    name := body[0]
    message := body[1]

    var envVars [2]string = app.getEnv()

    splitEmail := strings.Split(envVars[0], "=")
    email := splitEmail[1]
    splitPass := strings.Split(envVars[1], "=")

    from := email
    pass := splitPass[1]
    to := email

    msg := "\n" +
    "Subject: New Message" + "\n" +
    "Name: " + name + "\n" +
    "Message: " + message + "\n"
    
    err := smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", from, pass, "smtp.gmail.com"), from, []string{to}, []byte(msg))

    if err != nil{
        app.errorLog.Fatal("SMTP failure", err)
        return false
    }

    return true
}
