package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func main() {
	var cmd string

	fmt.Println("intput cmd, 1 to crteate keystore, 2 to import keystore")

	fmt.Scanf("%s", &cmd)

	switch cmd {
	case "1":
		createKs()
	case "2":
		var keyFile string
		fmt.Println("Input key file name in ./tmp:")
		fmt.Scanf("%s", &keyFile)
		keyFile = "./tmp/" + keyFile
		fmt.Println("keyFile:", keyFile)
		importKs(keyFile)
	}
}

func createKs() {
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "123123"
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) //
}

func importKs(keyFile string) {
	ks := keystore.NewKeyStore(
		"./wallets",
		keystore.StandardScryptN,
		keystore.StandardScryptP)

	jsonBytes, err := ioutil.ReadFile(keyFile)
	if err != nil {
		log.Fatal("readfile:", err)
	}

	password := "123123"
	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		log.Fatal("import:", err)
	}

	//
	if err := os.Remove(keyFile); err != nil {
		log.Fatal("remove:", err)
	}

	fmt.Println("import account:", account.Address.Hex()) //

}
