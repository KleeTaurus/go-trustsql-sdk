package identity

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

func TestCommonValidate(t *testing.T) {
	c := &Common{
		MchID:       "asdfsafsadaa",
		ProductCode: "123456789012",
		SeqNo:       "12345678901234567890123456789012",
		Sign:        "1111111111111111111111111111111111111111111111111111111111111111",
		Type:        "123456789012",
		TimeStamp:   time.Now(),
		ReqData:     "asdf",
	}
	validate = validator.New()

	errs := validate.Struct(c)
	if errs != nil {
		fmt.Println(errs)
	}
	jc, _ := json.Marshal(c)
	fmt.Println(string(jc))

}
