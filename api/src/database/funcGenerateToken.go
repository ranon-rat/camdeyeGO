package database

import (
	crRand "crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"math/big"
	"strconv"
	"time"
)

func GenerateToken(name string) string {

	bigInt, _ := crRand.Int(crRand.Reader, big.NewInt(1000000000000000000))
	token := hex.EncodeToString(sha256.New().Sum([]byte(name + strconv.Itoa(int(time.Now().Unix())) + bigInt.String())))
	return token
}
