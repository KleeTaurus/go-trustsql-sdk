package identity

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"sort"
	"strconv"
	"time"

	"gopkg.in/go-playground/validator.v9"
)

var (
	client   *http.Client
	validate *validator.Validate
)

func init() {
	client = &http.Client{
		Timeout: 1 * time.Second,
	}
	validate = validator.New()
}

// Linter 接口封装了请求餐具数据拼接要求
// 1.参数名ASCII码从小到大排序（字典序）；
// 2.如果参数的值为空不参与签名；
// 3.参数名区分大小写；
type Linter interface {
	Lint(c Common) (string, error)
}

const (
	// RegisteUserURI 注册用户
	// RegisteUserURI = "https://baas.trustsql.qq.com/idm_v1/api/user_cert/register"
	RegisteUserURI = "http://39.107.26.141:8004/post"
	// GetUserInfoURI 获取用户信息
	GetUserInfoURI = "https://baas.trustsql.qq.com/idm_v1/api/user_cert/fetch"
	// RegisteAccountURI 创建用户账户
	RegisteAccountURI = "https://baas.trustsql.qq.com/idm_v1/api/account_cert/register"
	// GetAccountsURI 获取用户的账户地址列表
	GetAccountsURI = "https://baas.trustsql.qq.com/idm_v1/api/account_cert/fetch_list"
	// GetPubkeyOfAccountURI 获取用户的账户公钥
	GetPubkeyOfAccountURI = "https://baas.trustsql.qq.com/idm_v1/api/account_cert/fetch"
)

// Common 公共信息
type Common struct {
	MchID       string `json:"mch_id"       validate:"required,len=12"`
	ProductCode string `json:"product_code" validate:"required,len=12"`
	SeqNo       string `json:"seq_no"       validate:"required,len=32"`
	Sign        string `json:"sign"         validate:"required,len=64"`
	Type        string `json:"type"         validate:"required,len=12"`
	TimeStamp   int64  `json:"time_stamp"   validate:"required"`
	ReqData     string `json:"req_data"     validate:"required"`
}

// UserRegister 注册用户需要的信息
type UserRegister struct {
	UserID       string `json:"user_id"        validate:"required"`
	PublicKey    string `json:"public_key"     validate:"required"`
	UserFullName string `json:"user_fullName"  validate:"required"`
}

func (u *UserRegister) lint(c *Common) string {
	signMap := make(map[string]string)
	getCheckString(&signMap, reflect.ValueOf((*u)))
	getCheckString(&signMap, reflect.ValueOf((*c)))
	var keys []string
	for k := range signMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	lintString := ""
	first := true
	for k := range keys {
		if !first {
			lintString = lintString + "&" + keys[k] + "=" + signMap[keys[k]]
		} else {
			lintString = keys[k] + "=" + signMap[keys[k]]
		}
		first = false
	}
	log.Printf("lintString is %s", lintString)
	return lintString
}

func getCheckString(m *map[string]string, v reflect.Value) {
	for i := 0; i < v.NumField(); i++ {
		if "sign" == v.Type().Field(i).Tag.Get("json") {
			continue
		}
		tag := v.Type().Field(i).Tag.Get("json")
		switch v.Field(i).Kind() {
		case reflect.Int64:
			{
				(*m)[tag] = strconv.FormatInt(v.Field(i).Interface().(int64), 10)
				continue
			}
		}

		(*m)[tag] = v.Field(i).Interface().(string)
	}
}

// RegisteUser 注册用户
func RegisteUser(u *UserRegister, c *Common) ([]byte, error) {
	data, err := json.Marshal(u)
	c.ReqData = string(data)
	sign := u.lint(c)
	if err != nil {
		return nil, err
	}
	c.Sign = sign
	reqData, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	// 校验common是否符合标准
	err = validate.Struct(c)
	if err != nil {
		return nil, err
	}

	log.Printf("RegisteUser req data is %s", string(reqData))

	// send http request
	req, err := http.NewRequest("POST", RegisteUserURI, bytes.NewBuffer(reqData))
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	log.Printf("RegisteUser response body is %s", string(body))
	_ = resp.Body.Close()
	return body, nil
}

// parseBody 处理response body
func parseBody(body []byte) {

}

// GetUserInfo 获取用户信息
func GetUserInfo() {
}

// RegisteAccount 创建用户账户
func RegisteAccount() {
}

// GetAccounts 获取用户的账户地址列表
func GetAccounts() {
}

// GetPubkeyOfAccount 获取用户的账户公钥
func GetPubkeyOfAccount() {
}
