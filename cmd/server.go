package main

import (
	"github.com/AhmedEnnaime/GoTest/config"
)

func main() {
	a := config.App{}

	a.Initialize("postgres", "3ea14367A4", "gotest")

	a.Run(":8080")

	print("connection success")

}
