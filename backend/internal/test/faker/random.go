package faker

import (
	"log"

	"github.com/bxcodec/faker/v3"
)

func RandomInt(min, max int) int {
	val, err := faker.RandomInt(min, max)
	if err != nil {
		log.Fatalf("Failed to generate random int: %v", err)
	}
	return val[0]
}
