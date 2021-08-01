package main

import (
	"math/big"

	"github.com/rockiecn/test-sig/tx/sendtx"
)

func main() {
	// prepair params
	eth := "http://localhost:8545"
	value := big.NewInt(1000000000000000000)
	toAddrHex := "0xd6071743390681c792cef53bedfef72a5a0cd8ef"
	var data []byte
	sk := "cb61e1519b560d994e4361b34c181656d916beb68513cff06c37eb7d258bf93d"
	// call
	sendtx.SendTransaction(eth, value, toAddrHex, data, sk)
}
