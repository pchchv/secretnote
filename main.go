package main

import (
	"./cipher"
	"bufio"
	"fmt"
	"os"
)

func main() {
	key := cipher.GetKey()
	hashedKey := cipher.GetHash(key)
	fmt.Print("Enter the message: ")
	reader := bufio.NewReader(os.Stdin)
	message, err := reader.ReadString('\n')
	if err != nil {
		panic(err.Error())
	}
	encryptedText := cipher.Encrypt(message, key)
	fmt.Printf("Your secret key: %s\n", key)
	if cipher.CheckKey(key, hashedKey) {
		fmt.Print(cipher.Decript(encryptedText, key))
	} else {
		panic("Invalid Key")
	}
}
