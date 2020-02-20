package encryptor

import (
	"crypto/aes"
	"crypto/cipher"
	"log"
)

var IV = []byte("1234567812345678")

func createCipher(key []byte) cipher.Block{
	ciph, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalf("Failed to create the AES cipher: %s", err)
	}
	return ciph
}

func AES(file []byte, key []byte) []byte{
	blockCipher := createCipher(key)
	stream := cipher.NewCTR(blockCipher, IV)
	stream.XORKeyStream(file, file)
	return file
}
