package identity

import (
	"bytes"
	"net/http"
	"time"
)

// Client 保存地址
type Client struct {
	RegisteUserURI        string
	GetUserInfoURI        string
	RegisteAccountURI     string
	GetAccountsURI        string
	GetPubkeyOfAccountURI string
}

// Common 公共信息
type Common struct {
	MchID       string    `json:"mch_id"       validate:"required,len=12"`
	ProductCode string    `json:"product_code" validate:"required,len=12"`
	SeqNo       string    `json:"seq_no"       validate:"required,len=32"`
	Sign        string    `json:"sign"         validate:"required,len=64"`
	Type        string    `json:"type"         validate:"required,len=12"`
	TimeStamp   time.Time `json:"time_stamp"   validate:"required"`
	ReqData     string    `json:"req_data"     validate:"required"`
}

var client *http.Client

func init() {
	client = &http.Client{
		Timeout: 1 * time.Second,
	}
}

// NewClient 创建客户端
func NewClient(registeUserURI, getUserInfoURI, registeAccountURI,
	getAccountsURI, getPubkeyOfAccountURI string) *Client {
	return &Client{
		RegisteUserURI:        registeUserURI,
		GetUserInfoURI:        getUserInfoURI,
		RegisteAccountURI:     registeAccountURI,
		GetAccountsURI:        getAccountsURI,
		GetPubkeyOfAccountURI: getPubkeyOfAccountURI,
	}

}

// RegisteUser 注册用户
func (c *Client) RegisteUser(data string) {
	req, err := http.NewRequest("POST", c.RegisteAccountURI, bytes.NewBuffer([]byte(data)))
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")

	resp, err := client.Do(req)
	if err != nil {
	}
	_ = resp.Body.Close()
}

// GetUserInfo 获取用户信息
func (c *Client) GetUserInfo() {
}

// RegisteAccount 创建用户账户
func (c *Client) RegisteAccount() {
}

// GetAccounts 获取用户的账户地址列表
func (c *Client) GetAccounts() {
}

// GetPubkeyOfAccount 获取用户的账户公钥
func (c *Client) GetPubkeyOfAccount() {
}
