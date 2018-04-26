package trustsql

import (
	"encoding/base64"

	"github.com/KleeTaurus/go-trustsql-sdk/identity"
	"github.com/KleeTaurus/go-trustsql-sdk/tscec"
	"github.com/KleeTaurus/go-trustsql-sdk/tsiss"
)

const (
	// AppendIssURI 共享信息查询
	AppendIssURI = "https://baas.trustsql.qq.com/cgi-bin/v1.0/trustsql_iss_append_v1.cgi"
	// QueryIssURI 共享信息查询
	QueryIssURI = "https://baas.trustsql.qq.com/cgi-bin/v1.0/trustsql_iss_query_v1.cgi"
)

// Client 腾讯区块链sdk
type Client struct {
	PrivateKey   []byte
	PublicKey    []byte
	AppendIssURI string
	QueryIssURI  string
}

// GenRandomPairkey 生成随机公私钥对
func GenRandomPairkey() *Client {
	privateKey, publicKey := tscec.NewKeyPair()
	client := Client{
		PrivateKey:   privateKey,
		PublicKey:    publicKey,
		AppendIssURI: AppendIssURI,
		QueryIssURI:  QueryIssURI,
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
		PrivateKey:   privKey,
		PublicKey:    pubKey,
		AppendIssURI: AppendIssURI,
		QueryIssURI:  QueryIssURI,
	}
	return &client, nil
}

// SetAppendIssURI 获取私钥的base64编码
func (c *Client) SetAppendIssURI(appendIssURI string) {
	c.AppendIssURI = appendIssURI
}

// SetQueryIssURI 获取私钥的base64编码
func (c *Client) SetQueryIssURI(queryIssURI string) {
	c.QueryIssURI = queryIssURI
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

// SetIssRequestTimeout 设置Iss的请求超时时间
func (c *Client) SetIssRequestTimeout(timeout time.Duration) {
	tsiss.SetRequestTimeout(timeout)
}

// GetIssSignStr 共享信息新增/追加, 第一步获取待签名串
// 注意: 留空sign字段
func (c *Client) GetIssSignStr(ia *tsiss.IssAppend) (string, error) {
	lintString := []byte(identity.Lint(nil, (*ia)))
	ia.MchSign = tscec.Sign(c.PrivateKey, lintString[:])

	signStr, err := tsiss.GetIssSignStr(c.AppendIssURI, ia)
	if err != nil {
		return "", err
	}
	return signStr, nil
}

// AppendIss 共享信息新增/追加, 第二步将signstr加入到参数ia中,再次请求接口
func (c *Client) AppendIss(ia *tsiss.IssAppend) (*tsiss.IssAppendResponse, error) {
	lintString := []byte(identity.Lint(nil, (*ia)))
	ia.MchSign = tscec.Sign(c.PrivateKey, lintString[:])

	isr, err := tsiss.AppendIss(c.AppendIssURI, ia)
	if err != nil {
		return nil, err
	}
	return isr, nil
}

// QueryIss 共享信息查询
func (c *Client) QueryIss(iq *tsiss.IssQuery) (*tsiss.IssResponse, error) {
	lintString := []byte(identity.Lint(nil, (*iq)))
	iq.MchSign = tscec.Sign(c.PrivateKey, lintString[:])

	isr, err := tsiss.QueryIss(c.QueryIssURI, iq)
	if err != nil {
		return nil, err
	}
	return isr, nil
}
