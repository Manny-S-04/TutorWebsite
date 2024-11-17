package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"website/pkg/models"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	errorLog *log.Logger
    infoLog *log.Logger
    reviews *models.ReviewModel
    templateCache map[string]*template.Template
}

type Config struct{
    Addr string
  StaticDir string
}

func main(){
       
	addr := flag.String("addr", ":4000", "HTTP network address")
    // user:pass@tcp(ip)/tableName
	dsn := flag.String("dsn", "web:pass@/reviews?parseTime=true", "MySQL db")

    flag.Parse()
/*
    file, err := os.Open("/logs/logs.txt")
    if err != nil{
        panic(err)
    }
    
    defer file.Close()

    multiWriter := io.MultiWriter(os.Stdout, file)
*/
    infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

    errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
    defer db.Close()

    templateCache, err := newTemplateCache()
    if err != nil{
        errorLog.Fatal(err)
    }

    app := &application{
        errorLog: errorLog,
        infoLog: infoLog,
        reviews: &models.ReviewModel{DB: db},
        templateCache: templateCache,
    }
    
    /*
    tlsConfig := &tls.Config{
        PreferServerCipherSuites: true,
        CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
    }
    */

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
        //TLSConfig: tlsConfig,
	}

    infoLog.Printf("Starting server on %s", *addr)
    err = srv.ListenAndServe()
    //err = srv.ListenAndServeTLS("./tls/cert.pem","./tls/key.pem")
    errorLog.Fatal(err)
}   


func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
    
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
