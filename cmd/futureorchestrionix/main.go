// cmd/futureorchestrionix/main.go
package main

import (
	"flag"
	"log"
	"os"

	"futureorchestrionix/internal/futureorchestrionix"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	app := futureorchestrionix.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
