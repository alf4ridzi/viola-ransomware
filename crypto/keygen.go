package crypto

import (
	"crypto/rand"
	"encoding/hex"
)

func KeyGenerator(longByte int) string {
	key := make([]byte, longByte)

	if _, err := rand.Read(key); err != nil {
		panic(err.Error())
	}

	strKeys := hex.EncodeToString(key)

	return strKeys
}
