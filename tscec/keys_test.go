package tscec

import (
	"encoding/base64"
	"log"
	"testing"
)

func TestGeneratePairkey(t *testing.T) {
	keyPair := GeneratePairkey()

	/*
		log.Printf("Private Key: %s, len: %d\n", base64Encode(keyPair.PrivateKey), len(keyPair.PrivateKey))
		log.Printf("Public key : %s, len: %d\n", base64Encode(keyPair.PublicKey), len(keyPair.PublicKey))
		log.Printf("Address    : %s, len: %d\n", keyPair.GetAddress(), len(keyPair.GetAddress()))
	*/

	if len(keyPair.PrivateKey) != 32 {
		t.Errorf("Incorrect length of the private key, it should be 32 bytes\n")
	}

	if len(keyPair.PublicKey) != 33 {
		t.Errorf("Incorrect length of the public key, it should be 33 bytes\n")
	}

	if len(keyPair.GetAddrByPubkey()) != 34 && len(keyPair.GetAddrByPubkey()) != 33 {
		t.Errorf("Incorrect length of the address, it should be 34 or 33 bytes\n")
	}
}

func TestPublicKeyAndAddress(t *testing.T) {

	keyPairs := [][]string{
		[]string{"Pc2DI133bYC6xcsKD1+Wwb1XkTZSqALiAhcGrzTXStM=", "Aim0N7NBzpOlic/LWbAaqc54IY+mK25Acf9tylpUi3CX", "19pvQGqsp6fig1tc6bjqWBGZmaCJbCkK83"},
		[]string{"XaaMJXKPjXKQKAjEGgiXaBhqIifI2b8fsjI+1qDIN2Y=", "A1v+Zfu1kbUEfDfi3kGoOxxxdB1JdFmcI5xScpubpids", "17dTMucUBwPkUEHwiVs7NY1tYn327yPpYx"},
		[]string{"98opKP6MzyTlNPcSN2ELywFSlASuervz/5okTxkvC/E=", "A4HNHgvDKMMluv1akCFDAtF5rNISWIsJXYPdk3yeXxgh", "1LEi1KXDU9BWS2ZFYBVWLGMunqi8Hxvwx"}}

	for i := range keyPairs {
		privateKey := keyPairs[i][0]
		publicKey := keyPairs[i][1]
		address := keyPairs[i][2]

		derivedPublicKey, err := GeneratePubkeyByPrvkey(base64Decode(privateKey))
		if err != nil {
			log.Panic(err)
		}
		if publicKey != base64Encode(derivedPublicKey) {
			t.Errorf("Incorrect derived public key, target: %s, derived: %s\n", publicKey, base64Encode(derivedPublicKey))
		}

		keyPair := &KeyPair{base64Decode(privateKey), base64Decode(publicKey)}
		if address != string(keyPair.GetAddrByPubkey()) {
			t.Errorf("Incorrect derived address, target: %s, derived: %s\n", address, keyPair.GetAddrByPubkey())
		}
	}
}

func base64Encode(input []byte) string {
	return base64.StdEncoding.EncodeToString(input)
}

func base64Decode(input string) []byte {
	src, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		log.Panic(err)
	}

	return src
}
