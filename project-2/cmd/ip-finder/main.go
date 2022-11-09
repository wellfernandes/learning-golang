package main

import (
	"fmt"
	"log"
	"os"
	"project-2/internal/app"
)

func main() {
	fmt.Println("Starting application...")

	application := app.Generate()
	err := application.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
