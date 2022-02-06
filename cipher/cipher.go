package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
)

func GetKey() (key string) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		panic(err.Error())
	}
	key = fmt.Sprintf("%x", bytes)
	return key
}

func GetHash(key string) (hashedKey string) {
	bytes := sha256.Sum256([]byte(key))
	hashedKey = fmt.Sprintf("%x", bytes)
	return hashedKey
}

func CheckKey(key string, hashedKey string) bool {
	hash := GetHash(key)
	if hash == hashedKey {
		return true
	} else {
		return false
	}
}

func Encrypt(text string, keyString string) (encryptedText string) {
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
	encryptedText = fmt.Sprintf("%x", cipherText)
	return encryptedText
}

func Decript(encryptedText string, keyString string) (decryptedText string) {
	key, err := hex.DecodeString(keyString)
	if err != nil {
		panic(err.Error())
	}
	enc, err := hex.DecodeString(encryptedText)
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
	nonceSize := aesGCM.NonceSize()
	nonce, cipherText := enc[:nonceSize], enc[nonceSize:]
	plaintext, err := aesGCM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		panic(err.Error())
	}
	decryptedText = fmt.Sprintf("%s", plaintext)
	return decryptedText
}
