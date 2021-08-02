package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// sk
	privateKeyECDSA, err := crypto.HexToECDSA("13cfe6b22d567ecb85ae524fdf115ac1bfdcefdc67cfeb698b9a24c9e478a341")
	if err != nil {
		log.Fatal(err)
	}

	// data
	data := []byte("hello")
	digest := crypto.Keccak256Hash(data)
	fmt.Println("digest:", digest.Hex()) //

	// sign
	sigByte, err := crypto.Sign(digest.Bytes(), privateKeyECDSA)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("signature:", hexutil.Encode(sigByte)) //

	//===========================================
	fmt.Println("now verify signature")

	// private to pub
	publicKeyCrypto := privateKeyECDSA.Public()
	// crypto to ecdsa
	publicKeyECDSA, ok := publicKeyCrypto.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return
	}

	// pubkey to address
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("fromAddress:", fromAddress)

	// pub to byte
	publicKeyByte := crypto.FromECDSAPub(publicKeyECDSA)

	// verify
	ret := crypto.VerifySignature(publicKeyByte, digest.Bytes(), sigByte)
	fmt.Println("verify:", ret)

}
