package decryptor
/*
import (
	"strings"
	"encoding/base64"
	"os"
	"bufio"
	"log"
	"io/ioutil"
	"path/filepath"
	"gucci/decryptor/sdecryptor"
	"gucci/decryptor/asdecryptor"
)

var (
	home, err = os.UserHomeDir()
	//desktop = filepath.Join(home, "Desktop")
	//pictures = filepath.Join(home, "Pictures")
	cwd = filepath.Join(home, ".Gucci")
)

func main() {
	f, err := os.Open(filepath.Join(cwd, "f.gucci"))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
		info1 := filepath.Join(cwd, "f.gucci")
		info2 := filepath.Join(cwd, "k.gucci")
		r := []string{info1, info2}

		for _, i := range r {
			err = os.Remove(i)
			if err != nil {
				log.Fatal(err)
			}
		}
	}()

	s := bufio.NewScanner(f)
	for s.Scan() {
		data := strings.Split(s.Text(), "   ")
		encryptedKey, err := base64.StdEncoding.DecodeString(data[0])
		if err != nil {
			log.Fatal(err)
		}
		file, err := base64.StdEncoding.DecodeString(data[1])
		if err != nil {
			log.Fatal(err)
		}

		privateKey, _ := ioutil.ReadFile("private.key")
		AESKey := asdecryptor.Decrypt(encryptedKey, privateKey)

		f, err := os.Open(string(file))
		if err != nil {
			log.Fatal(err)
		}
		reader := bufio.NewReader(f)
		buffer := make([]byte, 3000024)
		newFile := strings.TrimSuffix(string(file), ".gucci")

		for {
			count, err := reader.Read(buffer)
			if err != nil {
				if strings.Contains(err.Error(), "EOF") {
					break
				}
				log.Fatal(err)
			}

			f, err := os.OpenFile(newFile, os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				f, err = os.Create(newFile)
				if err != nil {
					log.Fatal(err)
				}

				_, err := f.Write(sdecryptor.Decrypt(buffer[:count], AESKey))
				if err != nil {
					f.Close()
					log.Fatal(err)
				}
				f.Close()
			} else {
				_, err := f.Write(sdecryptor.Decrypt(buffer[:count], AESKey))
				if err != nil {
					f.Close()
					log.Fatal(err)
				}

				f.Close()
			}
		}

		if err = f.Close(); err != nil {
			log.Fatal(err)
		}

		err = os.Remove(string(file))
		if err != nil {
			log.Fatal(err)
		}
	}

	if err = s.Err(); err != nil {
		log.Fatal(err)
	}
}*/