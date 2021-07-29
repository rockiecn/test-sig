package main

import (
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// generate key
	key, err := crypto.GenerateKey()
	if err != nil {
		fmt.Println(err)
	}

	// sk
	privateKey := hex.EncodeToString(key.D.Bytes())
	fmt.Println("privateKey:", privateKey)

	// address
	address := crypto.PubkeyToAddress(key.PublicKey).Hex()
	fmt.Println("address:", address)
}
