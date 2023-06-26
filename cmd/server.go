package main

import (
	"os"

	"github.com/AhmedEnnaime/GoTest/config"
)

func main() {
	a := config.App{}

	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	a.Run(":8080")

	print("connection success")

}
