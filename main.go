package main

import (
	"log"

	"github.com/chenghung/buffalo_learning/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
