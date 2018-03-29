package identity

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/KleeTaurus/go-trustsql-sdk/identity"
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

	c := identity.Common{
		MchID:       "xxxxxxxxxxxxxxxxx",
		ProductCode: "xxxxxxxxx",
		SeqNo:       "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		Sign:        "",
		Type:        "sign",
		TimeStamp:   time.Now().Unix(),
		ReqData:     "",
	}
	u := identity.UserRegister{
		PublicKey:    userPublicKey,
		UserID:       "xxxxxxxx",
		UserFullName: "xxxxxxxxxx",
	}
	_, err = identity.RegisteUser(u, c, keyPair)
	if err != nil {
		fmt.Println(err)
	}
}
