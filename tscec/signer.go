package tscec

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"encoding/base64"
	"math/big"

	secp256k1 "github.com/toxeus/go-secp256k1"
)

// Sign 签名
func Sign(privateKey []byte, data []byte, isHash bool) string {
	secp256k1.Start()
	defer secp256k1.Stop()
	var dataSHA256 [32]byte
	if !isHash {
		dataSHA256 = sha256.Sum256(data)
	} else {
		copy(dataSHA256[:32], data[:])
	}

	var dupPrivKey [32]byte
	copy(dupPrivKey[:32], privateKey[:])
	val, _ := secp256k1.Sign(dataSHA256, dupPrivKey, nil)

	// // DER encode
	// b, err := asn1.Marshal(val)
	// if err != nil {
	// 	// TODO err handing
	// 	fmt.Println("err der encode")
	// }
	// reg := regexp.MustCompile(`[\s*\t\n\r]`)
	// rep := []byte("")
	// val1 := reg.ReplaceAll(b, rep)
	// fmt.Println("----------------------------")
	// fmt.Println(base64.StdEncoding.EncodeToString(val))
	// fmt.Println(base64.StdEncoding.EncodeToString(val1))
	// fmt.Println("----------------------------")
	return string(base64.StdEncoding.EncodeToString(val))
}

// Verify 验证签名
func Verify(pubkey, sig, data []byte) bool {
	secp256k1.Start()
	defer secp256k1.Stop()
	dataSHA256 := sha256.Sum256(data)
	return secp256k1.Verify(dataSHA256, sig, pubkey)
}

// PrivateKeyFromBytes 根据[]byte构造返回ecdsa私钥
func PrivateKeyFromBytes(privateKey []byte) *ecdsa.PrivateKey {
	curve := elliptic.P256()
	x, y := curve.ScalarBaseMult(privateKey)

	privKey := &ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{
			Curve: curve,
			X:     x,
			Y:     y,
		},
		D: new(big.Int).SetBytes(privateKey),
	}

	return privKey
}
