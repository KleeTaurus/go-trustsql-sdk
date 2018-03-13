package trustsql

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/big"

	secp256k1 "github.com/toxeus/go-secp256k1"
)

func Sign(privateKey []byte, data []byte) string {
	// privKey := PrivateKeyFromBytes(privateKey)
	dataSHA256 := sha256.Sum256(data)

	/*
		r, s, err := ecdsa.Sign(rand.Reader, privKey, dataSHA256[:])
		if err != nil {
			log.Panic(err)
		}
		signature := append(r.Bytes(), s.Bytes()...)
	*/

	// val, success := secp256k1.Sign(dataSHA256, privateKey, nil)
	var dupPrivKey [32]byte
	copy(privateKey[:], dupPrivKey[:32])
	val, success := secp256k1.Sign(dataSHA256, dupPrivKey, nil)
	fmt.Println(val)
	fmt.Println(success)

	fmt.Println("sign:", base64.StdEncoding.EncodeToString(val))
	// return string(signature)
	return string(base64.StdEncoding.EncodeToString(val))
}

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
