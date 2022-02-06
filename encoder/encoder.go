package main

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func main() {
	bytesKey, key := getKey()
	hashedKey := getHash(bytesKey)
	fmt.Printf("key: %s\n", key)
	fmt.Printf("hashed key: %s", hashedKey)
}

func getKey() (bytes []byte, key string) {
	bytes = make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		panic(err.Error())
	}
	key = fmt.Sprintf("%x", bytes)
	return bytes, key
}

func getHash(key []byte) (hashedKey string) {
	bytes := sha256.Sum256(key)
	hashedKey = fmt.Sprintf("%x", bytes)
	return hashedKey
}
