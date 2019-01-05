package keys

import (
	"crypto/rand"
	"log"
)

func GenerateRandomBytes(bits int,) []byte {
	generated := make([]byte, bits)
	_, err := rand.Read(generated)
	if err != nil {
		log.Fatal(err)
	}

	return generated
}