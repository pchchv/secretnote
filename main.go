package main

import (
	"./cipher"
	"fmt"
	"os"
)

func main() {
	var message string
	key := cipher.GetKey()
	hashedKey := cipher.GetHash(key)
	fmt.Print("Enter the message: ")
	if _, err := fmt.Fscan(os.Stdin, &message); err != nil {
		panic(err.Error())
	}
	fmt.Printf("%s\n", cipher.Encrypt(message, key))
	fmt.Printf("Your secret key: %s\n", key)
	fmt.Print(cipher.CheckKey(key, hashedKey))
}
