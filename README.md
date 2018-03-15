GoTrustSQL
=======

[![license](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/KleeTaurus/go-trustsql-sdk/blob/master/LICENSE)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/KleeTaurus/go-trustsql-sdk)

基于 Golang 语言的 [TrustSQL](https://trustsql.qq.com/) SDK

## 概述
该 SDK 实现了底层密钥对生成、地址生成、签名/验签等基础功能，后期会对 TrustSQL 提供的 API 接口进行封装。

## 特性
SDK 基础命令
1. 生成密钥对
2. 根据私钥生成公钥
3. 根据公钥生成地址
4. 利用私钥对数据签名
5. 利用公钥对数据和签名进行验证

SDK API 接口
1. 数字资产
2. 信息共享
3. 身份管理

## 环境依赖
* go version >= 1.9
* 需要单独安装, 详细过程见 [github.com/toxeus/go-secp256k1](https://github.com/toxeus/go-secp256k1)
* cd $GOPATH/src/github.com/KleeTaurus/go-trustsql-sdk && [govendor](https://github.com/kardianos/govendor) sync

## 参考资料
1. [Bitcoin Wiki](https://en.bitcoin.it/wiki/Main_Page)
2. [Base58Check encoding](https://en.bitcoin.it/wiki/Base58Check_encoding)
3. [Technical background of version 1 Bitcoin addresses](https://en.bitcoin.it/wiki/Technical_background_of_version_1_Bitcoin_addresses)
4. [Merkle Trees](https://hackernoon.com/merkle-trees-181cb4bc30b4)
5. [数据库那么便宜，为何还要死贵的区块链来存储数据？](https://mp.weixin.qq.com/s/ME_E1EA95XILD_yaFg1d8Q)

## License
GoTrustSQL is MIT licensed. See the included LICENSE file for more details.
