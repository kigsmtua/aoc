package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(mineAdventCoins("bgvyzdsv", 5))
	fmt.Println(mineAdventCoins("bgvyzdsv", 6))
}

func mineAdventCoins(input string, nums int) int {
	number := 1
	magicInput := input
	for {
		md5hash := getMD5Hash(magicInput + strconv.Itoa(number))
		if startsWithXZeros(md5hash, nums) {
			break
		}
		number += 1
	}
	return number
}

// We check if the string has 5 preciding zeros
func startsWithXZeros(s string, nums int) bool {
	prefix := strings.Repeat("0", nums)
	return strings.HasPrefix(s, prefix)
}

// Returns the md5 hash of a string
func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
