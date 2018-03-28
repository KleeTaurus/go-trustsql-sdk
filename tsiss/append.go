package tsiss

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	// AppendIssURI 共享信息查询
	AppendIssURI = "https://baas.trustsql.qq.com/cgi-bin/v1.0/trustsql_iss_append.cgi"
)

// IssAppend 共享信息新增/追加,请求参数
type IssAppend struct {
	Version  string `json:"version"   validate:"required"` // 接口版本, 版本号，默认为1.0
	SignType string `json:"sign_type" validate:"required"` // 签名方式,签名类型，取值：ECDSA
	MchID    string `json:"mch_id"    validate:"required"` // 通讯方ID
	MchSign  string `json:"mch_sign"  validate:"required"` // 通讯方签名

	/// 业务参数
	NodeID      string                 `json:"node_id,omitempty"`             // 节点ID,可以通过baas平台->相应链->基本信息页面中获取
	ChainID     string                 `json:"chain_id" validate:"required"`  // 链ID,可以通过baas平台->相应链->基本信息页面中获取
	LedgerID    string                 `json:"ledger_id" validate:"required"` // 账本ID,可以通过baas平台->相应链->相应账本页面中获取
	InfoKey     string                 `json:"info_key,omitempty"`            // 信息标识,查询符合此信息标识号的记录
	InfoVersion string                 `json:"info_version,omitempty"`        // XXX 信息版本号, 查询符合此信息版本号的记录
	State       string                 `json:"state,omitempty"`               // XXX 记录状态, 由业务自行定义。
	Content     map[string]interface{} `json:"content,omitempty"`             // 记录内容,可以选填部分json字段，查询符合这些字段的记录
	Notes       map[string]interface{} `json:"notes,omitempty"`               // 记录注释

	CommitTime string `json:"commit_time,omitempty"` // 由业务自行定义。 格式：YYYY-MM-DD HH:mm:SS
	Account    string `json:"account,omitempty"`     // 记录方地址,查询符合此记录方地址的记录
	PublicKey  string `json:"public_key"`            // 记录方公钥
	Sign       string `json:"sign"`                  // 记录方签名,使用SDK函数IssSign进行签名
}

// IssAppendResponse 共享信息新增/追加, 返回参数
type IssAppendResponse struct {
	Version  string `json:"version"   validate:"required"` // 接口版本, 版本号，默认为1.0
	SignType string `json:"sign_type" validate:"required"` // 签名方式,签名类型，取值：ECDSA
	MchID    string `json:"mch_id"    validate:"required"` // 通讯方ID
	MchSign  string `json:"mch_sign"  validate:"required"` // 通讯方签名

	/// 业务参数
	Retcode     string                 `json:"retcode"`                // XXX 返回状态码，0表示成功，其它为失败
	Retmsg      string                 `json:"retmsg"`                 // 返回信息，如非空，为错误原因。
	InfoKey     string                 `json:"info_key,omitempty"`     // 信息标识,查询符合此信息标识号的记录
	InfoVersion string                 `json:"info_version,omitempty"` // XXX 信息版本号, 查询符合此信息版本号的记录
	State       string                 `json:"state,omitempty"`        // XXX 记录状态, 由业务自行定义。
	Content     map[string]interface{} `json:"content,omitempty"`      // 记录内容,可以选填部分json字段，查询符合这些字段的记录
	Notes       map[string]interface{} `json:"notes,omitempty"`        // 记录注释
	CommitTime  string                 `json:"commit_time,omitempty"`  // 由业务自行定义。 格式：YYYY-MM-DD HH:mm:SS
	Account     string                 `json:"account,omitempty"`      // 记录方地址,查询符合此记录方地址的记录
	PublicKey   string                 `json:"public_key"`             // 记录方公钥
	Sign        string                 `json:"sign"`                   // 记录方签名,使用SDK函数IssSign进行签名
	THash       string                 `json:"t_hash,omitempty"`       // 记录哈希
	BHeight     string                 `json:"b_height,omitempty"`     // XXX 区块高度
	BPrevHash   string                 `json:"b_prev_hash"`            // 前块哈希
	BHash       string                 `json:"b_hash"`                 // 本块哈希
	BTime       string                 `json:"b_time"`                 // 区块时间
}

// AppendIss 共享信息新增/追加
func AppendIss(iss IssAppend) (*IssAppendResponse, error) {
	// 校验common是否符合标准
	err := validate.Struct(iss)
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(iss)

	log.Printf("trustsql append iss request data is %s", string(data))

	// send http request
	req, err := http.NewRequest("POST", QueryIssURI, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	// log.Printf("trustsql response body is %s", string(body))
	_ = resp.Body.Close()

	// TODO delete these test code
	// body, err = ioutil.ReadFile("test_data/append_response.json")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// body = []byte(body)

	issAppendResponse := IssAppendResponse{}
	err = json.Unmarshal(body, &issAppendResponse)
	if err != nil {
		return nil, err
	}
	err = validate.Struct(issAppendResponse)
	if err != nil {
		return nil, err
	}
	return &issAppendResponse, nil
}
