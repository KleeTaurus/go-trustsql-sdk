package strustsql

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"log"
	_ "fmt"
	"golang.org/x/crypto/ripemd160"
	"math/big"

	secp256k1 "github.com/toxeus/go-secp256k1"
)

// Base64 编码
func B64Encode(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}
	
// Base64 解码
func B64Decode(dst string) []byte {
	src, _ := base64.StdEncoding.DecodeString(dst)
	return src
}

// Base58 编码，参考 https://en.bitcoin.it/wiki/Base58Check_encoding
func B58Encode(b []byte) string {
	const BASE58_TABLE = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

	/* 转换 big endian 到 big int */
	x := new(big.Int).SetBytes(b)

	r := new(big.Int)
	m := big.NewInt(58)
	zero := big.NewInt(0)

	s := ""

	/* 转换 big int 到字符串 */
	for x.Cmp(zero) > 0 {
		x.QuoRem(x, m, r)
		s = string(BASE58_TABLE[r.Int64()]) + s
	}

	return s
}

// 生成指定长度的随机数
func NewRandomBytes(size int) ([]byte, error) {
	randBytes := make([]byte, size)
	_, err := rand.Read(randBytes)
	if err != nil {
		return nil, err
	}
	return randBytes, nil
}

// 生成私钥
func NewPrivateKey() []byte {
	bytes, err := NewRandomBytes(32)
	if err != nil {
		log.Fatal(err)
	}
	return bytes
}

// 根据私钥生成公钥
func GetPublicKey(privateKey []byte) ([]byte, error) {
	var dupPrivateKey [32]byte
	for i := 0; i < 32; i++ {
		dupPrivateKey[i] = privateKey[i]
	}

	secp256k1.Start()
	publicKey, success := secp256k1.Pubkey_create(dupPrivateKey, false)
	if !success {
		return nil, errors.New("Failed to create public key from provided private key.")
	}
	secp256k1.Stop()

	return publicKey, nil
}

// 根据公钥生成地址
func GetAddress(publicKey []byte) string {
	hashSha256 := sha256.Sum256(publicKey)

	ripemd := ripemd160.New()
	ripemd.Reset()
	ripemd.Write(hashSha256[:])
	hashRipemd160 := ripemd.Sum(nil)

	hashTwiceSha256 := sha256.Sum256(hashRipemd160)
	hashTwiceSha256 = sha256.Sum256(hashTwiceSha256[:])

	result := make([]byte, 1)
	result = append(result, hashRipemd160...)

	result = append(result, hashTwiceSha256[:4]...)
	return "1" + B58Encode(result)
}

