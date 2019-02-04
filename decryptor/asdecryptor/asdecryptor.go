package asdecryptor

import (
	"crypto/rsa"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"log"
	"errors"
)

func Decrypt(encryptedData []byte, privatePemKey []byte) []byte{
	key := parseRsaPrivateKeyFromPemStr(privatePemKey)
	label := []byte("OAEP Encrypted")
	data, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, key, encryptedData, label)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func parseRsaPrivateKeyFromPemStr(privPEM []byte) *rsa.PrivateKey {
    block, _ := pem.Decode(privPEM)
    if block == nil {
        log.Fatal(errors.New("PEM data is not found"))
    }

    priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
    if err != nil {
        log.Fatal(err)
    }

    return priv
}