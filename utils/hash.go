package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashString(str string) string {
	byteData := []byte(str)
	hash := sha256.New()
	hash.Write(byteData)
	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	return hashString
}
