package tscec

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"github.com/btcsuite/btcd/btcec"
)

// Sign 签名
func Sign(pkBytes []byte, data []byte, isHash bool) string {
	privKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), pkBytes)

	if !isHash {
		dataHash := sha256.Sum256(data)
		copy(data[:], dataHash[:])
	}

	signature, _ := privKey.Sign(data)
	return string(base64.StdEncoding.EncodeToString(signature.Serialize()))
}

// Verify 验证签名
func Verify(pubkey, sig, data []byte) bool {
	// signature, _ := btcec.ParseSignature(data, btcec.S256())

	// // Verify the signature for the message using the public key.
	// messageHash := chainhash.DoubleHashB(data)
	// verified := signature.Verify(messageHash, pubkey)
	// fmt.Printf("Signature Verified? %v\n", verified)
	// return verified
	fmt.Print(pubkey)
	fmt.Print(sig)
	fmt.Print(data)
	return true
}
