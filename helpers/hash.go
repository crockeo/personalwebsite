package helpers

import (
	"crypto/md5"
	"encoding/hex"
)

// Hashing an input string as MD5
func HashString(input string) string {
	hasher := md5.New()
	hasher.Write([]byte(input))
	return hex.EncodeToString(hasher.Sum(nil))
}
