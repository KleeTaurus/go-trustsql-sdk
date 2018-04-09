package trustsql

import (
	"encoding/base64"
	"fmt"

	"github.com/KleeTaurus/go-trustsql-sdk/tscec"
)

// KeyPair 公私钥对数据结构
type KeyPair struct {
	PrivateKey []byte
	PublicKey  []byte
}

// GeneratePairkey 生成公私钥对
func GeneratePairkey() *KeyPair {
	privateKey, publicKey := tscec.NewKeyPair()
	keyPair := KeyPair{privateKey, publicKey}

	return &keyPair
}

// GeneratePairkeyByPrivateKey 通过base64编码的私钥生成KeyPair
func GeneratePairkeyByPrivateKey(privateKey string) (*KeyPair, error) {
	privKey, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return nil, err
	}
	pubKey, err := tscec.GeneratePubkeyByPrvkey(privKey)
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
	return tscec.GenerateAddrByPubkey(kp.PublicKey)
}

// SignString 对一个字符串进行签名（通常用于生成通讯方签名）
func (kp *KeyPair) SignString(s string) string {
	return tscec.Sign(kp.PrivateKey, []byte(s))
}

// VerifySignature 对签名进行验证
func (kp *KeyPair) VerifySignature(sig, data []byte) bool {
	return tscec.Verify(kp.PublicKey, sig, data)
}

// AppendIss 共享信息新增/追加
// TODO
func (kp *KeyPair) AppendIss() {
	fmt.Println("hello")
}

// QueryIss 共享信息查询
// TODO
func (kp *KeyPair) QueryIss() {}
