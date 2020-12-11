package main

import (
	"log"

	"github.com/scottcalebc/gradeview"
)

func main() {
	log.Println("Main started")

	cat := gradeview.NewCategory("Test", 1.0)

	log.Println("Created Category: ", cat)
}
