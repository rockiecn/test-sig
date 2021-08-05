package main

import (
	"math/big"

	"github.com/rockiecn/test-sig/tx/sendtx"
)

func main() {
	// prepair params
	eth := "http://localhost:8545"
	value := big.NewInt(1000000000000000000)
	toAddrHex := "0xb213d01542d129806d664248a380db8b12059061"
	var data []byte
	// sk for addr: 9e0153496067c20943724b79515472195a7aedaa
	sk := "cb61e1519b560d994e4361b34c181656d916beb68513cff06c37eb7d258bf93d"
	// call
	sendtx.SendTransaction(eth, value, toAddrHex, data, sk)
}
