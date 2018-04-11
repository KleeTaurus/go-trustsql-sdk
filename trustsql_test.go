package trustsql

import (
	"fmt"
	"testing"

	"github.com/KleeTaurus/go-trustsql-sdk/tsiss"
)

const (
	issGetSignStrTestURI = "http://39.107.26.141:8007/trustsql/v1.0/iss_get_sign_str"
)

func TestGeneratePairkey(t *testing.T) {
	client := GenRandomPairkey()

	/*
		log.Printf("Private Key: %s, len: %d\n", base64Encode(c.PrivateKey), len(c.PrivateKey))
		log.Printf("Public key : %s, len: %d\n", base64Encode(c.PublicKey), len(c.PublicKey))
		log.Printf("Address    : %s, len: %d\n", c.GetAddrByPubkey(), len(c.GetAddrByPubkey()))
	*/

	if len(client.PrivateKey) != 32 {
		t.Errorf("Incorrect length of the private key, it should be 32 bytes\n")
	}

	if len(client.PublicKey) != 33 {
		t.Errorf("Incorrect length of the public key, it should be 33 bytes\n")
	}

	if len(client.GetAddrByPubkey()) != 34 && len(client.GetAddrByPubkey()) != 33 {
		t.Errorf("Incorrect length of the address, it should be 34 or 33 bytes\n")
	}
}

func TestGetIssSignStr(t *testing.T) {
	client, err := NewTrustSQLByPrivateKey("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	if err != nil {
		t.Error("GeneratePairkeyByPrivateKey err")
	}
	client.SetAppendIssURI(issGetSignStrTestURI)

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

	signStr, err := client.GetIssSignStr(&ia)
	if err != nil {
		t.Errorf("GetIssSignStr failed %s", err)
	}

	fmt.Printf("signstr is: %+v\n", signStr)
}

func TestAppendIss(t *testing.T) {
}

func TestQueryIss(t *testing.T) {
	client := GenRandomPairkey()
	client.QueryIss()
}
