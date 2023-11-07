package main

import (
	"crypto/md5"
	"encoding/hex"
)

// Returns the md5 hash of a string
func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
