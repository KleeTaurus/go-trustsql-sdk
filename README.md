GoTrustSQL
=======

[![license](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/KleeTaurus/go-trustsql-sdk/blob/master/LICENSE)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/KleeTaurus/go-trustsql-sdk)

基于 Golang 语言的 [TrustSQL](https://trustsql.qq.com/) SDK

## 概述

该 SDK 实现了底层密钥对生成、地址生成、签名/验签等基础功能，并对 TrustSQL 提供的三类（信息共享/身份管理） API 接口进行了封装。

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
func SendToTrustSQL(content map[string]interface{}) (*tsiss.IssAppendResponse, error) {
    privateKey := "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
    client, _ := trustsql.NewClient(privateKey)
    client.SetIssRequestTimeout(5 * time.Second)
    //testURI := ""
    //client.SetAppendIssURI(testURI)
    accountAddr := client.GetAddrByPubkey()
    pubKey := client.GetPublicKey()
    issAppend := &tsiss.IssAppend{
        Version:  "1.0",
        SignType: "ECDSA",
        MchID:    "gbxxxxxxxxxxxxxxx",
        //MchSign:     "",
        Account:    string(accountAddr),
        CommitTime: time.Now().Format("2006-01-02 15:04:05"),
        //Content:    map[string]interface{}{"c": "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks"},
        Content: content,
        InfoKey: bson.NewObjectId().Hex(),
        //InfoKey:     "1242123jjj",
        InfoVersion: "1",
        State:       "0",
        Notes:       map[string]interface{}{"notes": "comments"},
        PublicKey:   pubKey,
        Sign:        "",
        ChainID:     "ch_tencent_test",
        LedgerID:    "ld_tencent_iss",
    }
    signStr, err := client.GetIssSignStr(issAppend)
    if err != nil {
        fmt.Printf("get issSignStr error: %s\n", err)
        return nil, err
    }

    issAppend.Sign = client.SignString(signStr, true)
    appendRes, err := client.AppendIss(issAppend)
    if err != nil {
        fmt.Printf("append error: %s\n", err)
        return nil, err
    }
    //fmt.Printf("appendRes: %+v\n", appendRes)
    return appendRes, nil
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
