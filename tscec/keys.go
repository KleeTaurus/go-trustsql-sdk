package tscec

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"log"

	"github.com/btcsuite/btcd/btcec"

	"golang.org/x/crypto/ripemd160"
)

const (
	version            = byte(0x00)
	addressChecksumLen = 4
	privateKeyLen      = 32
)

// NewKeyPair 生成公私钥对
func NewKeyPair() ([]byte, []byte) {
	curve := elliptic.P256()
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}

	publicKey, err := GeneratePubkeyByPrvkey(privateKey.D.Bytes())
	if err != nil {
		log.Panic(err)
	}

	return privateKey.D.Bytes(), publicKey
}

// GeneratePubkeyByPrvkey 根据私钥计算公钥
func GeneratePubkeyByPrvkey(p []byte) ([]byte, error) {
	curve := btcec.S256()
	_, publicKey := btcec.PrivKeyFromBytes(curve, p)
	// publicKey := privateKey.PubKey()
	// fmt.Println("-----------------------")
	// fmt.Println(privateKey.PubKey())
	// fmt.Println("-----------------------")
	// return publicKey.SerializeUncompressed(), nil
	return publicKey.SerializeCompressed(), nil
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
