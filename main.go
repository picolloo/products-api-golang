package main

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)


func main() { 

	app := App{}

	app.Initialize(
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	app.Run(":3000")
}