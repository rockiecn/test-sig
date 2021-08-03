package main

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rockiecn/test-sig/sig/implement/sigapi"
	"github.com/rockiecn/test-sig/sig/implement/utils"
)

func main() {
	var msg = []byte("hello") // msg to be signed
	var signed = false        // flag
	var sigByte []byte        // store signature
	var skHex string          // for sk input

	for cmd := "0"; true; {
		fmt.Println(">> Intput cmd, 1 to sign, 2 to verify, 3 to clear signature")
		fmt.Scanf("%s", &cmd)

		// decode Hex string to []byte
		addrByte, err := hex.DecodeString("9e0153496067C20943724b79515472195A7aEDAa")
		if err != nil {
			log.Fatal("decode err.")
			return
		}
		// []byte to common.Address
		fromAddress := common.BytesToAddress(addrByte)
		fmt.Println("fromAddress:", fromAddress)

		switch cmd {
		case "1":
			if signed {
				fmt.Println("already signed, run 3 to clear signature first")
				continue
			}

			fmt.Println("Input private key:")
			fmt.Scanf("%s", &skHex)
			// string to byte
			skByte := utils.Str2Byte(skHex)
			// call sign
			sigByte, err = sigapi.Sign(msg, skByte)
			if err != nil {
				log.Println("sign err:", err)
				continue
			}

			signed = true

			// bytes to string
			fmt.Println("signature:", hexutil.Encode(sigByte)) //

		case "2":
			if signed {
				sigapi.Verify(msg, sigByte, fromAddress)
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
