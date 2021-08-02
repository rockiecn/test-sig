package main

import (
	"bytes"
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	signed  = false
	sigByte []byte
	err     error
	msg     = []byte("hello")
	skHex   string
)

func main() {

	for cmd := "0"; true; {
		fmt.Println(">> Intput cmd, 1 to sign, 2 to verify, 3 to clear signature")
		fmt.Scanf("%s", &cmd)

		switch cmd {
		case "1":
			if signed {
				fmt.Println("already signed, run 3 to clear signature first")
				continue
			}

			fmt.Println("input private key:")
			fmt.Scanf("%s", &skHex)
			// string to byte
			skByte := str2byte(skHex)
			fmt.Println("skByte:", skByte)
			sigByte, err = sign(msg, skByte)
			if err != nil {
				log.Println("sign err:", err)
				continue
			}
			signed = true

			// bytes to string
			fmt.Println("signature:", hexutil.Encode(sigByte)) //

		case "2":
			if signed {
				verify(msg, sigByte)
				continue
			} else {
				fmt.Println("not signed yet, run 1 to signed first")
				continue
			}

		case "3":
			signed = false
			continue
		}
	}

}

//
func sign(msg []byte, skByte []byte) (sigRet []byte, err error) {

	fmt.Println(">> now start sign msg with private key")

	fmt.Println("msgHex:", hexutil.Encode(msg)) // bytes to string
	fmt.Println("skByte:", hexutil.Encode(skByte))

	// byte to string, then string to ecdsa
	privateKeyECDSA, err := crypto.HexToECDSA(byte2str(skByte))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// digest
	digest := crypto.Keccak256Hash(msg)
	// common.Hash to string
	fmt.Println("digest:", digest.Hex()) //
	//fmt.Println("digest:", digest)

	// sign to bytes
	sigByte, err := crypto.Sign(digest.Bytes(), privateKeyECDSA)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return sigByte, nil

}

func verify(msg []byte, sigByte []byte) (ok bool, err error) {
	fmt.Println(">> now start verify signature")

	// digest
	digest := crypto.Keccak256Hash(msg)

	// bytes to string
	fmt.Println("digest:", hexutil.Encode(digest.Bytes())) //
	fmt.Println("signature:", hexutil.Encode(sigByte))     //

	//=========== prepair something
	// get private key
	//privateKey, err := crypto.HexToECDSA("b91c265cabae210642d66f9d59137eac2fab2674f4c1c88df3b8e9e6c1f74f9f")
	privateKeyECDSA, err := crypto.HexToECDSA(skHex)
	if err != nil {
		log.Fatal(err)
		return false, err
	}

	// get pubkey from private key
	publicKeyCrypto := privateKeyECDSA.Public()
	publicKeyECDSA, ok := publicKeyCrypto.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return false, err
	}

	// pub ecdsa to bytes
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	//=================== verify
	// recover public key from signature
	sigPublicKey, err := crypto.Ecrecover(digest.Bytes(), sigByte)
	if err != nil {
		log.Fatal(err)
		return false, err
	}

	// compare 2 publickeys
	matches := bytes.Equal(sigPublicKey, publicKeyBytes)
	fmt.Println(matches) // true

	sigPublicKeyECDSA, err := crypto.SigToPub(digest.Bytes(), sigByte)
	if err != nil {
		log.Fatal(err)
		return false, err
	}

	sigPublicKeyBytes := crypto.FromECDSAPub(sigPublicKeyECDSA)
	matches = bytes.Equal(sigPublicKeyBytes, publicKeyBytes)
	fmt.Println(matches) // true

	signatureNoRecoverID := sigByte[:len(sigByte)-1] // remove recovery id
	verified := crypto.VerifySignature(publicKeyBytes, digest.Bytes(), signatureNoRecoverID)
	fmt.Println(verified) // true

	// signature to public key
	publicECDSA, err := crypto.SigToPub(digest.Bytes(), sigByte)
	if err != nil {
		log.Println("SigToPub err:", err)
		return false, err
	}
	// pub key to address
	recoveredAddr := crypto.PubkeyToAddress(*publicECDSA)
	fmt.Println("recovered Address:", recoveredAddr)

	return true, nil
}

func str2byte(str string) []byte {
	var ret []byte = []byte(str)
	return ret
}

func byte2str(data []byte) string {
	var str string = string(data[:len(data)])
	return str
}
