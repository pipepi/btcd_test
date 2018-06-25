package addr

import (
	"fmt"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
)

func GenerateBTCTest() (string, string, error) {
	privKey, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		return "", "", err
	}
	privKeyWif, err := btcutil.NewWIF(privKey, &chaincfg.TestNet3Params, false)
	if err != nil {
		return "", "", err
	}
	pubKeySerial := privKey.PubKey().SerializeUncompressed()
	pubKeyAddress, err := btcutil.NewAddressPubKey(pubKeySerial, &chaincfg.TestNet3Params)
	if err != nil {
		return "", "", err
	}
	return privKeyWif.String(), pubKeyAddress.EncodeAddress(), nil
}
func main() {
	wifKey, address, _ := GenerateBTCTest()
	fmt.Println(address, wifKey)
	//out mtNELkqMFJUHyKYxdcGJN6mszNwYH1JHyW 92TNcm4wXSb6zyov6x8TFg1NuQRZ6j7DMPNXkW2FeAYDrxLhSE7
}
