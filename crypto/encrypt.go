package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func Encrypt(key string, text string) (string, error) {
	plaintext := []byte(text)
	bytekey, err := hex.DecodeString(key)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(bytekey)
	if err != nil {
		fmt.Println("error creating AES block cipher ", err)
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println("error setting GCM mode", err)
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println("error generating nonce", err)
		return "", err
	}

	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
	enc := hex.EncodeToString(ciphertext)
	// fmt.Println("Enc data: ", enc)

	return enc, err
}
