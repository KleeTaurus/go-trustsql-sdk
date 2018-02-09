package strustsql

import (
	_ "fmt"
	"testing"
)

func TestNewPrivateKey(t *testing.T) {
	privateKey := NewPrivateKey()
	if len(privateKey) != 32 {
		t.Errorf("incorrect length of private key, %d\n", len(privateKey))
	}
	// fmt.Println(B64Encode(privateKey))
}

func TestGetPublicKey(t *testing.T) {
	encodedPrivateKey := "EOKJBDAp6NW7suUWB+QllRQfbqu9AjrrU1YjMY9EFjE="
	encodedPublicKey := "BGFmP3bFIwPpD/Aq1kZGCmMNHvVCl690MF/7T2azjGLpbaOICP6oAMqkhrvp2cyV/FPW+bwrf6hag+GwSu3E9Jw="
		
	publicKey, err := GetPublicKey(B64Decode(encodedPrivateKey))
	if err != nil {
		t.Errorf("Generating public key from private key failed.")
	}
	if B64Encode(publicKey) != encodedPublicKey {
		t.Errorf("Calculate public key from private key failed.")
	}
}

func TestGetAddress(t *testing.T) {
	encodedPublicKey := "BGFmP3bFIwPpD/Aq1kZGCmMNHvVCl690MF/7T2azjGLpbaOICP6oAMqkhrvp2cyV/FPW+bwrf6hag+GwSu3E9Jw="
	b58Address := "13GD9aH1jD5rLnGTEaJAhyCMMnZwfriwCB"
	publicKey := B64Decode(encodedPublicKey)
	addr := GetAddress(publicKey)
	if addr != b58Address {
		t.Errorf("Calculate address from public key failed.")
	}
	// fmt.Println(addr)
}

func TestKeyPairAddGen(t *testing.T) {
	encodedPrivateKey := "Q5XkNH6LpJu0S2GaBc8q4LPnudbEmYuJXquTUzfU1RA="
	encodedPublicKey := "BIFxQpcxsYzGKX379diqKQ9RvjzqY51giRK9UbAqcBQUP8DaooyCS75pxocPl/bQ8JtfjRVwZwXm5KZ8IIbJ59E="
	encodedAddress := "12cka46pKYpF31WfEXPC3oXcTMpGQiG9qq"

	publicKey, _ := GetPublicKey(B64Decode(encodedPrivateKey))
	address := GetAddress(publicKey)

	if B64Encode(publicKey) != encodedPublicKey {
		t.Errorf("Caculate public from private key failed.")
	}

	if address != encodedAddress {
		t.Errorf("Calculate address from public key failed.")
	}
}
