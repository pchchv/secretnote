package main

import (
	"bufio"
	"fmt"
	"os"
	"secretnote/crypt"
)

func main() {
	key := crypt.GetKey()
	hashedKey := crypt.GetHash(key)
	fmt.Print("Enter the message: ")
	reader := bufio.NewReader(os.Stdin)
	message, err := reader.ReadString('\n')
	if err != nil {
		panic(err.Error())
	}
	encryptedText := crypt.Encrypt(message, key)
	fmt.Printf("Your secret key: %s\n", key)
	if crypt.CheckKey(key, hashedKey) {
		fmt.Print(crypt.Decrypt(encryptedText, key))
	} else {
		panic("Invalid Key")
	}
}
