package tscec

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"encoding/base64"
	"math/big"

	secp256k1 "github.com/toxeus/go-secp256k1"
)

// SignString 对一个字符串进行签名（通常用于生成通讯方签名）
func (k *KeyPair) SignString(s string) string {
	return Sign(k.PrivateKey, []byte(s))
}

// Sign 签名
func Sign(privateKey []byte, data []byte) string {
	secp256k1.Start()
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
	copy(dupPrivKey[:32], privateKey[:])
	val, _ := secp256k1.Sign(dataSHA256, dupPrivKey, nil)
	secp256k1.Stop()
	return string(base64.StdEncoding.EncodeToString(val))
}

// Verify 验证签名
func Verify(pubkey, sig, data []byte) bool {
	secp256k1.Start()
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

// IssSign 生成一个用于共享信息的签名.
// @param infoKey          自定义的信息单号
// @param infoVersion      共享信息版本
// @param state            共享信息状态编码
// @param content          存放content的json格式字段
// @param notes            存放notes的json格式字段
// @param commitTime       提交共享信息时间 格式为YYYY-MM-DD HH:mm:
// @param privateKey       存放共享信息发起方的私钥 长度必须为PRVKEY_DIGEST_LENGTH
// @return                 返回签名
func IssSign(infoKey, infoVersion, content, notes, commitTime, privateKey string, state int) string {
	return ""
}

// IssVerifySign           验证一个共享信息的签名.
// @param infoKey          自定义的信息单号
// @param infoVersion      共享信息版本
// @param state            共享信息状态编码
// @param content          存放content的json格式字段
// @param notes            存放notes的json格式字段
// @param commitTime       提交共享信息时间 格式为YYYY-MM-DD HH:mm:SS
// @param pubkey           存放共享信息发起方的公钥 长度必须为PUBKEY_DIGEST_LENGTH
// @param sign             存放签名 长度必须为SIGN_DIGEST_LENGTH
// @return                 返回验证签名结果
func IssVerifySign(infoKey, infoVersion, content, notes, commitTime, pubkey, sign string, state int) bool {
	return true
}
