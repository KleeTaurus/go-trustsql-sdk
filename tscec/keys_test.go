package tscec

import (
	"encoding/base64"
	"log"
	"testing"
)

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

		derivedAddress := string(GenerateAddrByPubkey(base64Decode(publicKey)))
		if address != derivedAddress {
			t.Errorf("Incorrect derived address, target: %s, derived: %s\n", address, derivedAddress)
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
