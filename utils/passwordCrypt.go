package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Crypto(password string) string {
	hash := md5.Sum([]byte(password))
	str := hex.EncodeToString(hash[:])
	return str
}
