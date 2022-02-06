package main

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func main() {
	key := getKey()
	hashedKey := getHash(key)
	fmt.Print(checkKey(key, hashedKey))
}

func getKey() (key string) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		panic(err.Error())
	}
	key = fmt.Sprintf("%x", bytes)
	return key
}

func getHash(key string) (hashedKey string) {
	bytes := sha256.Sum256([]byte(key))
	hashedKey = fmt.Sprintf("%x", bytes)
	return hashedKey
}

func checkKey(key string, hashedKey string) bool {
	hash := getHash(key)
	if hash == hashedKey {
		return true
	} else {
		return false
	}
}
