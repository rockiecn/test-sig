package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func createKs() {
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "123123"
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) //
}

func importKs() {
	file := "./tmp/UTC--2021-07-01T07-05-14.130901482Z--9e0153496067c20943724b79515472195a7aedaa"
	ks := keystore.NewKeyStore(
		"./wallets",
		keystore.StandardScryptN,
		keystore.StandardScryptP)

	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal("readfile:", err)
	}

	password := "123123"
	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		log.Fatal("import:", err)
	}

	fmt.Println(account.Address.Hex()) //

	if err := os.Remove(file); err != nil {
		log.Fatal("remove:", err)
	}
}

func main() {
	//createKs()
	importKs()
}
