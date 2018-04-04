package identity

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/KleeTaurus/go-trustsql-sdk/tscec"
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
	k := tscec.GeneratePairkey()
	c := common()
	u := UserRegister{
		UserID:       "1111111",
		PublicKey:    "2222222",
		UserFullName: "3333333333",
	}
	_, err := RegisteUser(&u, &c, k)
	if err != nil {
		fmt.Println(err)
	}
}

func TestGetUserInfo(t *testing.T) {
	k := tscec.GeneratePairkey()
	c := common()
	u := UserInfo{
		UserID: "1111111",
	}
	_, err := GetUserInfo(&u, &c, k)
	if err != nil {
		fmt.Println(err)
	}
}

func TestRegisteAccount(t *testing.T) {
	k := tscec.GeneratePairkey()
	c := common()
	u := Account{
		UserID:    "1111111",
		PublicKey: "publicKey test data",
	}
	_, err := RegisteAccount(&u, &c, k)
	if err != nil {
		fmt.Println(err)
	}
}

func TestGetAccounts(t *testing.T) {
	k := tscec.GeneratePairkey()
	c := common()
	u := Accounts{
		UserID:    "1111111",
		State:     "state test data",
		BeginTime: "1111-11-11 22:22:22",
		EndTime:   "1111-11-11 22:22:22",
		Page:      213,
		Limit:     234,
	}
	_, err := GetAccounts(&u, &c, k)
	if err != nil {
		fmt.Println(err)
	}
}

func TestGetPubkeyOfAccount(t *testing.T) {
	k := tscec.GeneratePairkey()
	c := common()
	u := PubkeyOfAccount{
		UserID:         "1111111",
		AccountAddress: "accout_address test data",
	}
	_, err := GetPubkeyOfAccount(&u, &c, k)
	if err != nil {
		fmt.Println(err)
	}
}
