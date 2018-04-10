package tsiss

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// AppendIss 共享信息新增/追加
func AppendIss(appendIssURI string, iss *IssAppend) (*IssAppendResponse, error) {
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
	log.Printf("trustsql response body is %s", string(body))
	_ = resp.Body.Close()

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
