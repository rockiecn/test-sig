package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	fmt.Println(verifySig(
		"0x829814B6E4dfeC4b703F2c6fDba28F1724094D11",
		"0x53edb561b0c1719e46e1e6bbbd3d82ff798762a66d0282a9adf47a114e32cbc600c248c247ee1f0fb3a6136a05f0b776db4ac82180442d3a80f3d67dde8290811c",
		[]byte("hello"),
	))
}

// from address, signature, message
func verifySig(from, sigHex string, msg []byte) bool {
	fromAddr := common.HexToAddress(from)

	fmt.Println("from:", from)
	fmt.Println("fromAddr:", fromAddr)

	sig := hexutil.MustDecode(sigHex)
	fmt.Println("sigHex:", sigHex)
	fmt.Println("sig:", sig)
	if sig[64] != 27 && sig[64] != 28 {
		return false
	}
	sig[64] -= 27

	// signature to pubkey
	pubKey, err := crypto.SigToPub(signHash(msg), sig)
	if err != nil {
		return false
	}

	// pubkey to address
	recoveredAddr := crypto.PubkeyToAddress(*pubKey)

	return fromAddr == recoveredAddr
}

func signHash(data []byte) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256([]byte(msg))
}
