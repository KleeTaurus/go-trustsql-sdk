package trustsql

import (
	"fmt"
	"testing"

	"github.com/KleeTaurus/go-trustsql-sdk/identity"
	"github.com/KleeTaurus/go-trustsql-sdk/tscec"
	"github.com/KleeTaurus/go-trustsql-sdk/tsiss"
)

const (
	issGetSignStrTestURI = "http://39.107.26.141:8007/trustsql/v1.0/iss_get_sign_str"
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

func TestGetIssSignStr(t *testing.T) {
	keyPair, err := GeneratePairkeyByPrivateKey("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	if err != nil {
		t.Error("GeneratePairkeyByPrivateKey err")
	}

	ia := tsiss.IssAppend{
		Version:  "1.0",
		SignType: "ECDSA",
		MchID:    "xxxxxxxxxxxxxxxxx",
		MchSign:  "",

		ChainID:     "xxxxxxxxxxxxxxx",
		LedgerID:    "xxxxxxxxxxxxxx",
		InfoKey:     "xxxxxxxxxxxxxxxx",
		InfoVersion: "1",
		State:       "1",
		Content:     map[string]interface{}{"content": "test"},
		Notes:       map[string]interface{}{"note": "test"},

		CommitTime: "2018-04-04 16:47:31",
		Account:    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		PublicKey:  "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
	}

	lintString := []byte(identity.Lint(nil, ia))
	ia.MchSign = tscec.Sign(keyPair.PrivateKey, lintString[:])

	signStr, err := tsiss.GetIssSignStr(issGetSignStrTestURI, &ia)
	if err != nil {
		t.Error("GetIssSignStr err")
	}

	fmt.Printf("signstr is: %+v\n", signStr)
}

func TestAppendIss(t *testing.T) {
	// keyPair := GeneratePairkey()
	// keyPair.AppendIss()
}
