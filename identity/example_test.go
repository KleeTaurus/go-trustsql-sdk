package identity

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/KleeTaurus/go-trustsql-sdk/tscec"
)

func ExampleRegisteUser() {
	privKey, err := base64.StdEncoding.DecodeString("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	if err != nil {
		fmt.Println("error")
	}

	pubKey, err := tscec.GeneratePubkeyByPrvkey(privKey)
	if err != nil {
		fmt.Println("error")
	}
	keyPair := tscec.KeyPair{
		PrivateKey: privKey,
		PublicKey:  pubKey,
	}
	// p := string(base64.StdEncoding.EncodeToString(pubKey))

	userKeyPair := tscec.GeneratePairkey()
	userPublicKey := string(base64.StdEncoding.EncodeToString(userKeyPair.PublicKey))

	c := Common{
		MchID:       "xxxxxxxxxxxxxxxxx",
		ProductCode: "xxxxxxxxx",
		SeqNo:       "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		Sign:        "",
		Type:        "sign",
		TimeStamp:   time.Now().Unix(),
		ReqData:     "",
	}
	u := UserRegister{
		PublicKey:    userPublicKey,
		UserID:       "xxxxxxxx",
		UserFullName: "xxxxxxxxxx",
	}
	_, err = RegisteUser(&u, &c, &keyPair)
	if err != nil {
		fmt.Println(err)
	}
}

func ExamplePairKey() {
	// 创建公私钥对, 公钥采用压缩格式
	keyPair := tscec.GeneratePairkey()
	fmt.Printf("Private Key: %s\n", base64.StdEncoding.EncodeToString(keyPair.PrivateKey))
	fmt.Printf("Public Key: %s\n", base64.StdEncoding.EncodeToString(keyPair.PublicKey))
	fmt.Printf("Address: %s\n", keyPair.GenerateAddrByPubkey())
}
