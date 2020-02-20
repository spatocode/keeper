package encryptor

import (
	"crypto/rsa"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"log"
	"errors"
	"io/ioutil"
)

func RSA(data string, path string) []byte{
	publicPemKey := 	`
-----BEGIN RSA PUBLIC KEY-----
MIIBCgKCAQEAp7nDFXU1vESrwPasI4NOwfh6gzwdbkwgqFf0VOm1OjyCubiLzfzE
VRhKmQgxnX/7iLRhxC+URg6uTQd5absiNpeeeFUcQcbWfKxRxpF5sn7wDmNwVYGM
YU2khX/5WRfhXXZBqsxfIT3y2jQ3bJC4qaxOUI/F6+DtLluIQPm1XMjwjzHmj/VV
cbtETuBq3uQIHOz4mk5Tm9HV2IhkTsXG/YxXSLzKczqKcs1jsxWuWgE3by/YIsQv
JzInEhBnzg8G9gh4xz2eyGfgCZLnOKNJzOJBx6Rcq3VChboRJ/mwsP3358uNy4Gv
TPgh/0UMEMafTKQaWuUjpW/nBa7Neej5awIDAQAB
-----END RSA PUBLIC KEY-----
	`
	err := ioutil.WriteFile(path, []byte(publicPemKey), 0644)
	if err != nil {
		log.Fatal(err)
	}
	
	key := parseRsaPublicKeyFromPemStr(publicPemKey)
	label := []byte("OAEP Encrypted")
	cipherText, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, key, []byte(data), label)
	if err != nil {
		log.Fatal("Encryption error", err)
	}
	return cipherText
}

func parseRsaPublicKeyFromPemStr(pubPEM string) *rsa.PublicKey {
    block, _ := pem.Decode([]byte(pubPEM))
    if block == nil {
        log.Fatal(errors.New("failed to parse PEM block containing the key"))
    }

    pub, err := x509.ParsePKCS1PublicKey(block.Bytes)
    if err != nil {
        log.Fatal("ParsePKIXPublicKey error: ", err)
    }

    return pub
}