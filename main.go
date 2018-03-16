package main

import (
	"log"

	"github.com/torreswoo/blog/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
