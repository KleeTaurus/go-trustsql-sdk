package trustsql

import (
	"fmt"
	"testing"
)

func TestNewPrivateKey(t *testing.T) {
	privateKey := NewPrivateKey()
	if len(privateKey) != 32 {
		t.Errorf("Incorrect length of private key, %d\n", len(privateKey))
	}
	fmt.Println("Private Key:", B64Encode(privateKey))
}

func TestGetPublicKey(t *testing.T) {
	encodedPrivateKey := "EOKJBDAp6NW7suUWB+QllRQfbqu9AjrrU1YjMY9EFjE="
	encodedPublicKey := "BGFmP3bFIwPpD/Aq1kZGCmMNHvVCl690MF/7T2azjGLpbaOICP6oAMqkhrvp2cyV/FPW+bwrf6hag+GwSu3E9Jw="
	
	publicKey, _ := GetPublicKey(B64Decode(encodedPrivateKey))
	if B64Encode(publicKey) != encodedPublicKey {
		t.Errorf("The derived public key is not identical to the target public key.")
	}
}

func TestGetAddress(t *testing.T) {
	encodedPrivateKey := "Q5XkNH6LpJu0S2GaBc8q4LPnudbEmYuJXquTUzfU1RA="
	// encodedPublicKey := "BIFxQpcxsYzGKX379diqKQ9RvjzqY51giRK9UbAqcBQUP8DaooyCS75pxocPl/bQ8JtfjRVwZwXm5KZ8IIbJ59E="
	encodedAddress := "12cka46pKYpF31WfEXPC3oXcTMpGQiG9qq"

	publicKey, _ := GetPublicKey(B64Decode(encodedPrivateKey))
	address := GetAddress(publicKey)

	if address != encodedAddress {
		t.Errorf("The derived address is not identical to the target address.")
	}
}
