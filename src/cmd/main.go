package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	_ = godotenv.Load("./.env")

    envLoaded := os.Getenv("LOADED")
    if envLoaded != "YES"{
        panic(fmt.Errorf("Failed to load env"))
    }

	e := echo.New()
	db := ConnectDatabase(e)
	defer db.db.Close()
	RegisterHandlers(e, db)
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}
	e.Logger.Fatal(e.Start(port))
}
