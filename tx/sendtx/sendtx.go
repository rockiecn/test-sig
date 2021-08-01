package sendtx

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
)

// send signed tx to given chain
func SendTransaction(eth string, value *big.Int, toAddrHex string, data []byte, sk string) error { // private key to sign tx
	client, err := ethclient.Dial(eth)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// prepair privatekey
	privateKey, err := crypto.HexToECDSA(sk)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// sk to pubkey
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return err
	}

	// from: pubkey to address
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("from:", fromAddress)
	fmt.Println("nonce:", nonce)

	// construct tx
	//value := big.NewInt(1000000000000000000) // in wei (1 eth)
	gasLimit := uint64(21000) // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
		return err
	}
	// to address
	toAddress := common.HexToAddress(toAddrHex)

	// var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	// get chain id
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("chain ID:", chainID)

	// sign tx
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// get signed tx bytes
	ts := types.Transactions{signedTx}
	rawTxBytes := ts.GetRlp(0)
	//rawTxHex := hex.EncodeToString(rawTxBytes)

	// decode bytes to Transaction and send it
	tx = new(types.Transaction)
	rlp.DecodeBytes(rawTxBytes, &tx)
	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Printf("tx hash: %s\n", tx.Hash().Hex()) //
	return nil
}