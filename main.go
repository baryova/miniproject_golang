package main

import (
	"miniproject_golang/config"
	"miniproject_golang/routes"
)

func main() {
	config.Init()

	e := routes.New()

	e.Logger.Fatal(e.Start(":8000"))
}
