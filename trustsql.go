package trustsql

import (
	"encoding/base64"
	"net/http"
	"time"

	"github.com/KleeTaurus/go-trustsql-sdk/tscec"
	"github.com/KleeTaurus/go-trustsql-sdk/tsiss"
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

// trustsql URI 地址
const (
	AppendIssURI          = "https://baas.trustsql.qq.com/cgi-bin/v1.0/trustsql_iss_append_v1.cgi" // 共享信息查询
	QueryIssURI           = "https://baas.trustsql.qq.com/cgi-bin/v1.0/trustsql_iss_query_v1.cgi"  // 共享信息查询
	RegisteUserURI        = "https://baas.trustsql.qq.com/idm_v1.1/api/user_cert/register"         // 注册用户
	GetUserInfoURI        = "https://baas.trustsql.qq.com/idm_v1.1/api/user_cert/fetch"            // 获取用户信息
	RegisteAccountURI     = "https://baas.trustsql.qq.com/idm_v1.1/api/account_cert/register"      // 创建用户账户
	GetAccountsURI        = "https://baas.trustsql.qq.com/idm_v1.1/api/account_cert/fetch_list"    // 获取用户的账户地址列表
	GetPubkeyOfAccountURI = "https://baas.trustsql.qq.com/idm_v1.1/api/account_cert/fetch"         // 获取用户的账户公钥
)

// Client 腾讯区块链sdk
type Client struct {
	PrivateKey            []byte
	PublicKey             []byte
	appendIssURI          string // 共享信息查询
	queryIssURI           string // 共享信息查询
	registeUserURI        string // 注册用户
	getUserInfoURI        string // 获取用户信息
	registeAccountURI     string // 创建用户账户
	getAccountsURI        string // 获取用户的账户地址列表
	getPubkeyOfAccountURI string // 获取用户的账户公钥
}

// GenRandomPairkey 生成随机公私钥对
func GenRandomPairkey() *Client {
	privateKey, publicKey := tscec.NewKeyPair()
	client := Client{
		PrivateKey:            privateKey,
		PublicKey:             publicKey,
		appendIssURI:          AppendIssURI,
		queryIssURI:           QueryIssURI,
		registeUserURI:        RegisteUserURI,
		getUserInfoURI:        GetUserInfoURI,
		registeAccountURI:     RegisteAccountURI,
		getAccountsURI:        GetAccountsURI,
		getPubkeyOfAccountURI: GetPubkeyOfAccountURI,
	}

	return &client
}

// NewClient 通过base64编码的私钥生成client
func NewClient(privateKey string) (*Client, error) {
	privKey, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return nil, err
	}
	pubKey, err := tscec.GeneratePubkeyByPrvkey(privKey)
	if err != nil {
		return nil, err
	}
	client := Client{
		PrivateKey:            privKey,
		PublicKey:             pubKey,
		appendIssURI:          AppendIssURI,
		queryIssURI:           QueryIssURI,
		registeUserURI:        RegisteUserURI,
		getUserInfoURI:        GetUserInfoURI,
		registeAccountURI:     RegisteAccountURI,
		getAccountsURI:        GetAccountsURI,
		getPubkeyOfAccountURI: GetPubkeyOfAccountURI,
	}
	return &client, nil
}

// SetAppendIssURI 设置URI
func (c *Client) SetAppendIssURI(appendIssURI string) {
	c.appendIssURI = appendIssURI
}

// SetQueryIssURI 设置URI
func (c *Client) SetQueryIssURI(queryIssURI string) {
	c.queryIssURI = queryIssURI
}

// SetRegisteUserURI  设置URI
func (c *Client) SetRegisteUserURI(registeUserURI string) {
	c.registeUserURI = registeUserURI
}

// SetGetUserInfoURI 设置URI
func (c *Client) SetGetUserInfoURI(getUserInfoURI string) {
	c.getUserInfoURI = getUserInfoURI
}

// SetRegisteAccountURI 设置URI
func (c *Client) SetRegisteAccountURI(registeAccountURI string) {
	c.registeAccountURI = registeAccountURI
}

// SetGetAccountsURI 设置URI
func (c *Client) SetGetAccountsURI(getAccountsURI string) {
	c.getAccountsURI = getAccountsURI
}

// SetGetPubkeyOfAccountURI 设置URI
func (c *Client) SetGetPubkeyOfAccountURI(getPubkeyOfAccountURI string) {
	c.getPubkeyOfAccountURI = getPubkeyOfAccountURI
}

// GetPrivateKey 获取私钥的base64编码
func (c *Client) GetPrivateKey() string {
	return base64.StdEncoding.EncodeToString(c.PrivateKey)
}

// GetPublicKey 获取公钥的base64编码
func (c *Client) GetPublicKey() string {
	return base64.StdEncoding.EncodeToString(c.PublicKey)
}

// GetAddrByPubkey 计算公钥对应的地址
func (c *Client) GetAddrByPubkey() []byte {
	return tscec.GenerateAddrByPubkey(c.PublicKey)
}

// SignString 对一个字符串进行签名（通常用于生成通讯方签名）
func (c *Client) SignString(s string) string {
	return tscec.Sign(c.PrivateKey, []byte(s))
}

// VerifySignature 对签名进行验证
func (c *Client) VerifySignature(sig, data []byte) bool {
	return tscec.Verify(c.PublicKey, sig, data)
}

// GetIssSignStr 共享信息新增/追加, 第一步获取待签名串
// 注意: 留空sign字段
func (c *Client) GetIssSignStr(ia *tsiss.IssAppend) (string, error) {
	lintString := []byte(Lint(nil, (*ia)))
	ia.MchSign = tscec.Sign(c.PrivateKey, lintString[:])

	signStr, err := tsiss.GetIssSignStr(c.appendIssURI, ia)
	if err != nil {
		return "", err
	}
	return signStr, nil
}

// AppendIss 共享信息新增/追加, 第二步将signstr加入到参数ia中,再次请求接口
func (c *Client) AppendIss(ia *tsiss.IssAppend) (*tsiss.IssAppendResponse, error) {
	lintString := []byte(Lint(nil, (*ia)))
	ia.MchSign = tscec.Sign(c.PrivateKey, lintString[:])

	isr, err := tsiss.AppendIss(c.appendIssURI, ia)
	if err != nil {
		return nil, err
	}
	return isr, nil
}

// QueryIss 共享信息查询
func (c *Client) QueryIss(iq *tsiss.IssQuery) (*tsiss.IssResponse, error) {
	lintString := []byte(Lint(nil, (*iq)))
	iq.MchSign = tscec.Sign(c.PrivateKey, lintString[:])

	isr, err := tsiss.QueryIss(c.queryIssURI, iq)
	if err != nil {
		return nil, err
	}
	return isr, nil
}

// RegisteUser 注册用户
func (c *Client) RegisteUser(u *UserRegister, common *Common) ([]byte, error) {
	body, err := send(c.registeUserURI, u, common, c.PrivateKey)
	if err != nil {
		return nil, err
	}
	// 检查返回值是否成功
	err = responseUtil(body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// GetUserInfo 获取用户信息参数
func (c *Client) GetUserInfo(u *UserInfo, common *Common) ([]byte, error) {
	body, err := send(c.getUserInfoURI, u, common, c.PrivateKey)
	if err != nil {
		return nil, err
	}
	// 检查返回值是否成功
	err = responseUtil(body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// RegisteAccount 创建用户账户
func (c *Client) RegisteAccount(u *Account, common *Common) ([]byte, error) {
	body, err := send(c.registeAccountURI, u, common, c.PrivateKey)
	if err != nil {
		return nil, err
	}
	// 检查返回值是否成功
	err = responseUtil(body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// GetAccounts 获取用户的账户地址列表
func (c *Client) GetAccounts(u *Accounts, common *Common) ([]byte, error) {
	body, err := send(c.getAccountsURI, u, common, c.PrivateKey)
	if err != nil {
		return nil, err
	}
	// 检查返回值是否成功
	err = responseUtil(body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// GetPubkeyOfAccount 获取用户的账户公钥
func (c *Client) GetPubkeyOfAccount(u *PubkeyOfAccount, common *Common) ([]byte, error) {
	body, err := send(c.getPubkeyOfAccountURI, u, common, c.PrivateKey)
	if err != nil {
		return nil, err
	}
	// 检查返回值是否成功
	err = responseUtil(body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
