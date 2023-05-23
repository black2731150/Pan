package common

import (
	"crypto/md5"
	"encoding/hex"
)

func StringMD5(s string) string {
	hash := md5.New()
	hash.Write([]byte(s))
	hashedBytes := hash.Sum(nil)
	md5str := hex.EncodeToString(hashedBytes)
	return md5str
}
