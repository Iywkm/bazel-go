package main

import (
	"log"

	"github.com/Iywkm/bazel-go/uuid"
)

func main() {
	id, err := uuid.Generate()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(id)
}