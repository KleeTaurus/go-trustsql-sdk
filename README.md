GoTrustSQL
=======

[![license](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/KleeTaurus/go-trustsql-sdk/blob/master/LICENSE)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/KleeTaurus/go-trustsql-sdk)

基于 Golang 语言的 [TrustSQL](https://trustsql.qq.com/) SDK

## 概述

该 SDK 实现了底层密钥对生成、地址生成、签名/验签等基础功能，并对 TrustSQL 提供的三类（数字资产/信息共享/身份管理） API 接口进行了封装。

## 特性

SDK 基础命令
1. 生成密钥对
2. 根据私钥生成公钥（压缩公钥）
3. 根据公钥生成地址（压缩地址）
4. 利用私钥对数据签名
5. 利用公钥对数据和签名进行验签

SDK API 接口
1. 数字资产（暂未实现）
2. 信息共享
3. 身份管理

## 示例

下列示例演示了该 SDK 的基本使用方法。

```go
package trustsql

import (
	"encoding/base64"
	"fmt"

	"github.com/KleeTaurus/go-trustsql-sdk/tscec"
)

const (
	// AppendIssURI 共享信息查询
	AppendIssURI = "https://baas.trustsql.qq.com/cgi-bin/v1.0/trustsql_iss_append_v1.cgi"
	// QueryIssURI 共享信息查询
	QueryIssURI = "https://baas.trustsql.qq.com/cgi-bin/v1.0/trustsql_iss_query_v1.cgi"
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

```

## 环境依赖

* go version >= 1.9
* 需要单独安装, 详细过程见 [github.com/toxeus/go-secp256k1](https://github.com/toxeus/go-secp256k1)
* cd $GOPATH/src/github.com/KleeTaurus/go-trustsql-sdk && [govendor](https://github.com/kardianos/govendor) sync

## 参考资料

1. [Bitcoin Wiki](https://en.bitcoin.it/wiki/Main_Page)
2. [Base58Check encoding](https://en.bitcoin.it/wiki/Base58Check_encoding)
3. [Bitcoin Developer Reference](https://bitcoin.org/en/developer-reference#block-chain)
4. [Technical background of version 1 Bitcoin addresses](https://en.bitcoin.it/wiki/Technical_background_of_version_1_Bitcoin_addresses)
5. [Merkle Trees](https://hackernoon.com/merkle-trees-181cb4bc30b4)
6. [数据库那么便宜，为何还要死贵的区块链来存储数据？](https://mp.weixin.qq.com/s/ME_E1EA95XILD_yaFg1d8Q)
7. [Data Insertion in Bitcoin's Blockchain](https://digitalcommons.augustana.edu/cgi/viewcontent.cgi?article=1000&context=cscfaculty)

## License

GoTrustSQL is MIT licensed. See the included LICENSE file for more details.
