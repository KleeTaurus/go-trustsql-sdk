package tscec

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"log"

	secp256k1 "github.com/toxeus/go-secp256k1"
	"golang.org/x/crypto/ripemd160"
)

const version = byte(0x00)
const addressChecksumLen = 4
const privateKeyLen = 32

// KeyPair 公私钥对数据结构
type KeyPair struct {
	PrivateKey []byte
	PublicKey  []byte
}

// GetPublicKey 获取公钥的base64编码
func (kp *KeyPair) GetPublicKey() string {
	return base64.StdEncoding.EncodeToString(kp.PublicKey)
}

// GetPrivateKey 获取私钥的base64编码
func (kp *KeyPair) GetPrivateKey() string {
	return base64.StdEncoding.EncodeToString(kp.PrivateKey)
}

// GeneratePairkey 生成公私钥对
func GeneratePairkey() *KeyPair {
	privateKey, publicKey := newKeyPair()
	keyPair := KeyPair{privateKey, publicKey}

	return &keyPair
}

// GeneratePairkeyByPrivateKey 通过base64编码的私钥生成KeyPair
func GeneratePairkeyByPrivateKey(privateKey string) (*KeyPair, error) {
	privKey, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return nil, err
	}
	pubKey, err := GeneratePubkeyByPrvkey(privKey)
	if err != nil {
		return nil, err
	}
	keyPair := KeyPair{
		PrivateKey: privKey,
		PublicKey:  pubKey,
	}
	return &keyPair, nil
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

// GetAddrByPubkey 计算公钥对应的地址
func (kp *KeyPair) GetAddrByPubkey() []byte {
	return GenerateAddrByPubkey(kp.PublicKey)
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

// 生成公私钥对
func newKeyPair() ([]byte, []byte) {
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

// 计算公钥校验码
func checksum(payload []byte) []byte {
	firstSHA := sha256.Sum256(payload)
	secondSHA := sha256.Sum256(firstSHA[:])

	return secondSHA[:addressChecksumLen]
}
