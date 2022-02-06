package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func main() {
	var message string
	key := getKey()
	hashedKey := getHash(key)
	fmt.Print("Enter the message: ")
	if _, err := fmt.Fscan(os.Stdin, &message); err != nil {
		panic(err.Error())
	}
	fmt.Printf("%s\n", encrypt(message, key))
	fmt.Printf("Your secret key: %s\n", key)
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

func encrypt(text string, keyString string) (encryptedText string) {
	key, err := hex.DecodeString(keyString)
	plaintext := []byte(text)
	if err != nil {
		panic(err.Error())
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	cipherText := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return fmt.Sprintf("%x", cipherText)
}
