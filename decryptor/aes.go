package decryptor

import (
	"log"
	"crypto/cipher"
	"crypto/aes"
)

var IV = []byte("1234567812345678")

func createCipher(key []byte) cipher.Block{
	ciph, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}
	return ciph
}

func AES(data []byte, key []byte) []byte{
    blockCipher := createCipher(key)
    stream := cipher.NewCTR(blockCipher, IV)
    stream.XORKeyStream(data, data)
    return data
}
