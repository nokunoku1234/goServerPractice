package main

import (
	"log"

	"goServerPractice/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
