package main

import (
	"errors"
	"fmt"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcutil"
)

var params = &chaincfg.MainNetParams

func CreatePrivateKey() (*btcutil.WIF, error) {
	secret, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		return nil, err
	}
	return btcutil.NewWIF(secret, params, true)
}

func ImportWIF(wifStr string) (*btcutil.WIF, error) {
	wif, err := btcutil.DecodeWIF(wifStr)
	if err != nil {
		return nil, err
	}
	if !wif.IsForNet(params) {
		return nil, errors.New("the wif string is not valid for the bitcoin network")
	}
	return wif, nil
}

func GetAddress(wif *btcutil.WIF) (*btcutil.AddressPubKey, error) {
	return btcutil.NewAddressPubKey(wif.PrivKey.PubKey().SerializeCompressed(), params)
}

func pubKeyHashToScriptHash(pubKeyHash []byte) []byte {
	script, err := txscript.NewScriptBuilder().
		AddOp(txscript.OP_0).AddData(pubKeyHash).Script()
	if err != nil {
		panic(err)
	}

	redeemScriptHash := btcutil.Hash160(script)
	script, err = txscript.NewScriptBuilder().
		AddOp(txscript.OP_HASH160).AddData(redeemScriptHash).
		AddOp(txscript.OP_EQUAL).Script()
	if err != nil {
		panic(err)
	}
	return script
}

func main() {
	wif, _ := CreatePrivateKey()
	//wif, _ := ImportWIF("your compressed privateKey Wif")
	address, _ := GetAddress(wif)

	fmt.Println("Common Address:", address.EncodeAddress())

	pubKey := wif.PrivKey.PubKey().SerializeCompressed()
	pubKeyHash := btcutil.Hash160(pubKey)

	scriptPubKey := pubKeyHashToScriptHash(pubKeyHash)

	w, err := btcutil.NewAddressScriptHashFromHash(scriptPubKey[2:22], params)
	if err != nil {
		panic(err)
	}
	fmt.Println("Segregated Witness Address:", w.String())
	fmt.Println("PrivateKeyWifCompressed:", wif.String())
}
