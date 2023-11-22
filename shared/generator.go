package shared

import (
	"log"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

func GenerateCode(len int) string {
	charset := "abcdefghijklmnopqrstuvwxyz"

	id, err := gonanoid.Generate(charset, len)
	if err != nil {
		log.Fatalf("failed to generate code: %v", err)
	}

	return id
}
