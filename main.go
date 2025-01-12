package main

import (
	"bufio"
	"crypto/ecdsa"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/lyonnee/key25519/bip32"
	"github.com/lyonnee/key25519/bip39"
	"github.com/lyonnee/key25519/bip44"
)

// main is the entry point of the program. It reads a seed phrase from the standard input,
// generates a BIP39 seed from the mnemonic, derives a master key using BIP32, and then
// derives a child key according to the BIP44 path "m/44'/501'/0'/0'". Finally, it converts
// the derived key to an Ed25519 private key and prints both the private and public keys.

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter seed phrase: ")
	mnemonic, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	seed := bip39.ToSeed(mnemonic, "")

	masterKey := bip32.GenerateMasterKey(seed)
	path := "m/44'/60'/0'/0/0" // Ethereum path for first address in MetaMask

	indexs, err := bip44.ParsePath(path)
	if err != nil {
		log.Fatal(err)
	}

	var newKey = masterKey
	for _, v := range indexs {
		newKey = bip32.CKDPriv(newKey, v)
	}

	privKey, err := crypto.ToECDSA(newKey.PrivKey)
	if err != nil {
		log.Fatal(err)
	}

	pubKey := privKey.PublicKey
	publicKeyBytes := crypto.FromECDSAPub(&pubKey)
	fmt.Printf("Public Key: %x\n", publicKeyBytes)

}

func PubKey(mnemonic string) ecdsa.PublicKey {

	seed := bip39.ToSeed(mnemonic, "")

	masterKey := bip32.GenerateMasterKey(seed)
	path := "m/44'/60'/0'/0/0" // Ethereum path for first address in MetaMask

	indexs, err := bip44.ParsePath(path)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("indexs: %v", indexs)

	var newKey = masterKey
	for _, v := range indexs {
		newKey = bip32.CKDPriv(newKey, v)
	}

	log.Printf("masterKey: %x", masterKey)

	log.Printf("newKey: %x", newKey)

	privKey, err := crypto.ToECDSA(newKey.PrivKey)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("privKey: %x", privKey)

	return privKey.PublicKey
	//publicKeyBytes := crypto.FromECDSAPub(&pubKey)
	//return fmt.Sprintf("%x", publicKeyBytes)
	//return publicKeyBytes
}
