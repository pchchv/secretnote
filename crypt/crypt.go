package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
)

type Note struct {
	HashedKey string `json:"key"`
	Text      string `json:"text"`
}

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
	plaintext := []byte(text)
	aesGCM := cryptHelper(keyString)
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	cipherText := aesGCM.Seal(nonce, nonce, plaintext, nil)
	encryptedText = fmt.Sprintf("%x", cipherText)
	return encryptedText
}

func Decrypt(keyString string, note Note) (decryptedText string) {
	if CheckKey(keyString, note.HashedKey) {
		enc, err := hex.DecodeString(note.Text)
		if err != nil {
			panic(err.Error())
		}
		aesGCM := cryptHelper(keyString)
		nonceSize := aesGCM.NonceSize()
		nonce, cipherText := enc[:nonceSize], enc[nonceSize:]
		plaintext, err := aesGCM.Open(nil, nonce, cipherText, nil)
		if err != nil {
			panic(err.Error())
		}
		decryptedText = fmt.Sprintf("%s", plaintext)
		return decryptedText
	} else {
		panic("Incorrect key")
	}
}

func cryptHelper(keyString string) cipher.AEAD {
	key, err := hex.DecodeString(keyString)
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
	return aesGCM
}
