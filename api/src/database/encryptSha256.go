package database

import (
	"crypto/sha256"
	"encoding/hex"
)

func EncryptSha256(text string) string {

	return hex.EncodeToString(sha256.New().Sum([]byte(text)))
}
