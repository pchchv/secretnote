package main

import (
	"bufio"
	"fmt"
	"os"
	"secretnote/crypt"
)

func main() {
	var note crypt.Note
	fmt.Print("Enter the message: ")
	var reader = bufio.NewReader(os.Stdin)
	message, err := reader.ReadString('\n')
	if err != nil {
		panic(err.Error())
	}
	key := crypt.GetKey()
	note.Text = crypt.Encrypt(message, key)
	note.HashedKey = crypt.GetHash(key)
	fmt.Printf("Your secret key: %s\n", key)
	fmt.Print(crypt.Decrypt(key, note))
}
