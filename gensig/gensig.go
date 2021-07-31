package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// sk
	privateKey, err := crypto.HexToECDSA("cb61e1519b560d994e4361b34c181656d916beb68513cff06c37eb7d258bf93d")
	if err != nil {
		log.Fatal(err)
	}

	// data
	data := []byte("hello")
	hash := crypto.Keccak256Hash(data)
	fmt.Println("hash:", hash.Hex()) //

	// sign
	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("signature:", hexutil.Encode(signature)) //
}
