package tsiss

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// send 共享信息新增/追加
func send(appendIssURI string, iss *IssAppend) ([]byte, error) {
	// 校验common是否符合标准
	err := validate.Struct(iss)
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(iss)

	log.Printf("trustsql append iss request data is %s", string(data))

	req, err := http.NewRequest("POST", appendIssURI, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	_ = resp.Body.Close()
	log.Printf("trustsql response body is %s", string(body))
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

type signStr struct {
	SignStr string `json:"sign_str"`
}

// GetIssSignStr 共享信息新增/追加
func GetIssSignStr(appendIssURI string, iss *IssAppend) (string, error) {
	body, err := send(appendIssURI, iss)
	if err != nil {
		return "", err
	}
	s := signStr{}
	err = json.Unmarshal(body, &s)
	if err != nil {
		return "", err
	}
	return s.SignStr, nil
}

// AppendIss 共享信息新增/追加
func AppendIss(appendIssURI string, iss *IssAppend) (*IssAppendResponse, error) {
	body, err := send(appendIssURI, iss)
	if err != nil {
		return nil, err
	}

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
