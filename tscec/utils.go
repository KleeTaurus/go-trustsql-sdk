package tscec

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"log"

	secp256k1 "github.com/toxeus/go-secp256k1"
	"golang.org/x/crypto/ripemd160"
)

// 计算公钥校验码
func checksum(payload []byte) []byte {
	firstSHA := sha256.Sum256(payload)
	secondSHA := sha256.Sum256(firstSHA[:])

	return secondSHA[:addressChecksumLen]
}

// IntToHex converts an int64 to a byte array
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

// ReverseBytes reverses a byte array
func ReverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

// GeneratePubkeyByPrvkey 根据私钥计算公钥
func GeneratePubkeyByPrvkey(privateKey []byte) ([]byte, error) {
	var dupPrivateKey [privateKeyLen]byte
	copy(dupPrivateKey[:], privateKey[:privateKeyLen])

	secp256k1.Start()
	// TODO 改成可配置
	// 此处生成压缩公钥
	publicKey, success := secp256k1.Pubkey_create(dupPrivateKey, true)
	if !success {
		return nil, errors.New("failed to create public key from the provided private key")
	}
	secp256k1.Stop()

	return publicKey, nil
}

// GenerateAddrByPubkey 计算公钥对应的地址
func GenerateAddrByPubkey(publicKey []byte) []byte {
	publicKeyHash := HashPublicKey(publicKey)

	// https: //en.bitcoin.it/wiki/Technical_background_of_version_1_Bitcoin_addresses

	versionPayload := append([]byte{version}, publicKeyHash...)
	checksum := checksum(versionPayload)

	fullPayload := append(versionPayload, checksum...)
	address := Base58Encode(fullPayload)

	return address
}

// HashPublicKey 哈希公钥
func HashPublicKey(publicKey []byte) []byte {
	publicKeySHA256 := sha256.Sum256(publicKey)

	RIPEMD160Hasher := ripemd160.New()
	_, err := RIPEMD160Hasher.Write(publicKeySHA256[:])
	if err != nil {
		log.Panic(err)
	}
	publicKeyRIPEMD160 := RIPEMD160Hasher.Sum(nil)

	return publicKeyRIPEMD160
}
