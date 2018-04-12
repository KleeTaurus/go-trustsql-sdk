package trustsql

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"gopkg.in/go-playground/validator.v9"
)

func common() Common {
	return Common{
		MchID:       "asdfsafsadaa",
		ProductCode: "123456789012",
		SeqNo:       "12345678901234567890123456789012",
		Sign:        "1111111111111111111111111111111111111111111111111111111111111111",
		Type:        "123456789012",
		TimeStamp:   time.Now().Unix(),
		ReqData:     "asdf",
	}

}

func TestCommonValidate(t *testing.T) {
	c := common()
	validate = validator.New()

	errs := validate.Struct(c)
	if errs != nil {
		fmt.Println(errs)
	}
	_, _ = json.MarshalIndent(c, "", "    ")
	// fmt.Println(string(jc))
}

func TestRegisteUser(t *testing.T) {
	client := GenRandomPairkey()

	c := common()
	u := UserRegister{
		UserID:       "1111111",
		PublicKey:    "2222222",
		UserFullName: "3333333333",
	}
	_, err := client.RegisteUser(&u, &c)
	if err != nil {
		fmt.Println(err)
	}
}

// func TestGetUserInfo(t *testing.T) {
// 	privateKey, _ := tscec.NewKeyPair()
// 	c := common()
// 	u := identity.UserInfo{
// 		UserID: "1111111",
// 	}
// 	_, err := identity.GetUserInfo(&u, &c, privateKey)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }

// func TestRegisteAccount(t *testing.T) {
// 	privateKey, _ := tscec.NewKeyPair()
// 	c := common()
// 	u := identity.Account{
// 		UserID:    "1111111",
// 		PublicKey: "publicKey test data",
// 	}
// 	_, err := identity.RegisteAccount(&u, &c, privateKey)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }

// func TestGetAccounts(t *testing.T) {
// 	privateKey, _ := tscec.NewKeyPair()
// 	c := common()
// 	u := identity.Accounts{
// 		UserID:    "1111111",
// 		State:     "state test data",
// 		BeginTime: "1111-11-11 22:22:22",
// 		EndTime:   "1111-11-11 22:22:22",
// 		Page:      213,
// 		Limit:     234,
// 	}
// 	_, err := identity.GetAccounts(&u, &c, privateKey)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }

// func TestGetPubkeyOfAccount(t *testing.T) {
// 	privateKey, _ := tscec.NewKeyPair()
// 	c := common()
// 	u := identity.PubkeyOfAccount{
// 		UserID:         "1111111",
// 		AccountAddress: "accout_address test data",
// 	}
// 	_, err := identity.GetPubkeyOfAccount(&u, &c, privateKey)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }
