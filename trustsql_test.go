package trustsql

import (
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

func TestAppendIss(t *testing.T) {
	keyPair := GeneratePairkey()
	keyPair.AppendIss()
}
