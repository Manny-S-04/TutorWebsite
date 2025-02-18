package main

import (
	"fmt"
	"os"
	"path/filepath"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	projectRoot, _ := filepath.Abs(filepath.Join(".", "../"))

    fmt.Println(projectRoot)

	envPath := filepath.Join(projectRoot, "src/.env")
    fmt.Println(envPath)
	_ = godotenv.Load(envPath)

	envLoaded := os.Getenv("LOADED")
	if envLoaded != "YES" {
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
