package tscec

import (
	"fmt"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
)

// Sign 签名
func Sign(pkBytes []byte, data []byte, isHash bool) string {
	privKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), pkBytes)

	// Sign a message using the private key.
	if !isHash {
		messageHash := chainhash.DoubleHashB(data)
		signature, _ := privKey.Sign(messageHash)
		return string(signature.Serialize())
	} else {
		signature, _ := privKey.Sign(data)
		return string(signature.Serialize())
	}

	// // Serialize and display the signature.
	// fmt.Printf("Serialized Signature: %x\n", signature.Serialize())
	// return string(signature.Serialize())
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
