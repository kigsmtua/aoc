package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(mineAdventCoins("bgvyzdsv"))
}

func mineAdventCoins(input string) int {
	number := 1
	magicInput := input
	for {
		md5hash := getMD5Hash(magicInput + strconv.Itoa(number))
		if startsWithFiveZeros(md5hash) {
			break
		}
		number += 1
	}
	return number
}

// We check if the string has 5 preciding zeros
func startsWithFiveZeros(s string) bool {
	prefix := "00000"
	return strings.HasPrefix(s, prefix)
}

// Returns the md5 hash of a string
func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
