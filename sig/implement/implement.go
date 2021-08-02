package main

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rockiecn/test-sig/sig/implement/api"
	"github.com/rockiecn/test-sig/sig/implement/utils"
)

var (
	signed  = false
	sigByte []byte
	msg     = []byte("hello")
	skHex   string
)

func main() {

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
		fmt.Println("addr:", fromAddress)

		switch cmd {
		case "1":
			if signed {
				fmt.Println("already signed, run 3 to clear signature first")
				continue
			}

			fmt.Println("input private key:")
			fmt.Scanf("%s", &skHex)
			// string to byte
			skByte := utils.Str2Byte(skHex)
			fmt.Println("skByte:", skByte)
			sigByte, err = api.Sign(msg, skByte)
			if err != nil {
				log.Println("sign err:", err)
				continue
			}
			signed = true

			// bytes to string
			fmt.Println("signature:", hexutil.Encode(sigByte)) //

		case "2":
			if signed {
				api.Verify(msg, sigByte, fromAddress)
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
