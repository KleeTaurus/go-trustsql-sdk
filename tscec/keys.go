package tscec

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/base64"
	"log"
)

const version = byte(0x00)
const addressChecksumLen = 4
const privateKeyLen = 32

// KeyPair 公私钥对数据结构
type KeyPair struct {
	PrivateKey []byte
	PublicKey  []byte
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

// GetPrivateKey 获取私钥的base64编码
func (kp *KeyPair) GetPrivateKey() string {
	return base64.StdEncoding.EncodeToString(kp.PrivateKey)
}

// GetPublicKey 获取公钥的base64编码
func (kp *KeyPair) GetPublicKey() string {
	return base64.StdEncoding.EncodeToString(kp.PublicKey)
}

// GetAddrByPubkey 计算公钥对应的地址
func (kp *KeyPair) GetAddrByPubkey() []byte {
	return GenerateAddrByPubkey(kp.PublicKey)
}

// SignString 对一个字符串进行签名（通常用于生成通讯方签名）
func (kp *KeyPair) SignString(s string) string {
	return Sign(kp.PrivateKey, []byte(s))
}

// VerifySignature 对签名进行验证
func (kp *KeyPair) VerifySignature(sig, data []byte) bool {
	return Verify(kp.PublicKey, sig, data)
}
