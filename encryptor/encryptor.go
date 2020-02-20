package encryptor

import (
	"bufio"
	"os"
	"strings"
)

func Encrypt(file, key string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}

	reader := bufio.NewReader(f)
	data := make([]byte, 300024)
	temp := file + ".kee"
	for {
		count, err := reader.Read(data)
		if err != nil {
			if strings.Contains(err.Error(), "EOF") {
				break
			}
			return err
		}

		f, err := os.OpenFile(temp, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			f, err = os.Create(temp)
			if err != nil {
				return err
			}

			_, err = f.Write(AES(data[:count], []byte(key)))
			if err != nil {
				f.Close()
				return err
			}
			f.Close()
		} else {
			_, err = f.Write(AES(data[:count], []byte(key)))
			if err != nil {
				f.Close()
				return err
			}
			f.Close()
		}
	}

	if err = f.Close(); err != nil {
		return err
	}

	err = os.Remove(file)
	if err != nil {
		return err
	}

	err = os.Rename(temp, file)
	if err != nil {
		return err
	}
	return nil
}
