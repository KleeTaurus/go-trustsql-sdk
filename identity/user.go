package identity

import (
	"net/http"
	"time"

	"github.com/KleeTaurus/go-trustsql-sdk/tscec"
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

const (
	// RegisteUserURI 注册用户
	RegisteUserURI = "https://baas.trustsql.qq.com/idm_v1.1/api/user_cert/register"
	// GetUserInfoURI 获取用户信息
	GetUserInfoURI = "https://baas.trustsql.qq.com/idm_v1.1/api/user_cert/fetch"
	// RegisteAccountURI 创建用户账户
	RegisteAccountURI = "https://baas.trustsql.qq.com/idm_v1.1/api/account_cert/register"
	// GetAccountsURI 获取用户的账户地址列表
	GetAccountsURI = "https://baas.trustsql.qq.com/idm_v1.1/api/account_cert/fetch_list"
	// GetPubkeyOfAccountURI 获取用户的账户公钥
	GetPubkeyOfAccountURI = "https://baas.trustsql.qq.com/idm_v1.1/api/account_cert/fetch"
)

// Common 公共信息
type Common struct {
	MchID       string `json:"mch_id"       validate:"required"`
	ProductCode string `json:"product_code" validate:"required"`
	SeqNo       string `json:"seq_no"       validate:"required"`
	Sign        string `json:"sign"         validate:"required"`
	Type        string `json:"type"         validate:"required"`
	TimeStamp   int64  `json:"time_stamp"   validate:"required"`
	ReqData     string `json:"req_data"     validate:"required"`
}

// UserRegister 注册用户需要的信息
type UserRegister struct {
	PublicKey    string `json:"public_key"     validate:"required"`
	UserID       string `json:"user_id"        validate:"required"`
	UserFullName string `json:"user_fullName"  validate:"required"`
}

// RegisteUser 注册用户
func RegisteUser(u UserRegister, c Common, k tscec.KeyPair) ([]byte, error) {
	return send(RegisteUserURI, u, c, k)
}

// UserInfo 获取用户信息参数
type UserInfo struct {
	UserID string `json:"user_id" validate:"required"`
}

// GetUserInfo 获取用户信息参数
func GetUserInfo(u UserInfo, c Common, k tscec.KeyPair) ([]byte, error) {
	return send(GetUserInfoURI, u, c, k)
}

// Account 创建用户账户参数
type Account struct {
	UserID    string `json:"user_id"        validate:"required"`
	PublicKey string `json:"public_key"     validate:"required"`
}

// RegisteAccount 创建用户账户
func RegisteAccount(u Account, c Common, k tscec.KeyPair) ([]byte, error) {
	return send(RegisteAccountURI, u, c, k)
}

// Accounts 获取用户的账户地址列表参数
type Accounts struct {
	UserID    string `json:"user_id"        validate:"required"`
	State     string `json:"state"          validate:"omitempty"`
	BeginTime string `json:"begin_time"     validate:"omitempty"`
	EndTime   string `json:"end_time"       validate:"omitempty"`
	Page      int    `json:"page"           validate:"omitempty"`
	Limit     int    `json:"limit"          validate:"omitempty"`
}

// GetAccounts 获取用户的账户地址列表
func GetAccounts(u Accounts, c Common, k tscec.KeyPair) ([]byte, error) {
	return send(GetAccountsURI, u, c, k)
}

// PubkeyOfAccount 获取用户的账户公钥参数
type PubkeyOfAccount struct {
	UserID         string `json:"user_id"         validate:"required"`
	AccountAddress string `json:"account_address" validate:"required"`
}

// GetPubkeyOfAccount 获取用户的账户公钥
func GetPubkeyOfAccount(u PubkeyOfAccount, c Common, k tscec.KeyPair) ([]byte, error) {
	return send(GetPubkeyOfAccountURI, u, c, k)
}
