package identity

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/KleeTaurus/go-trustsql-sdk/tscec"
	"gopkg.in/go-playground/validator.v9"
)

func common() *Common {
	return &Common{
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
	u := &UserRegister{
		UserID:       "1111111",
		PublicKey:    "2222222",
		UserFullName: "3333333333",
	}
	_, err := RegisteUser(u, c, k)
	if err != nil {
		fmt.Println(err)
	}
}
