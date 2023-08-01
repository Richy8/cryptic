package main

import (
	"fmt"

	"github.com/Richy8/cryptic/decrypt"
	"github.com/Richy8/cryptic/encrypt"
)

func main() {
	encrytedStr := encrypt.Nimbus("Richy Jones")
	decrytedStr := decrypt.Nimbus(encrytedStr)

	fmt.Println("Encrypted string => ", encrytedStr)
	fmt.Println("Decrypted string => ", decrytedStr)
}