package trustsql

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/KleeTaurus/go-trustsql-sdk/tscec"
)

func send(URI string, u interface{}, c *Common, privateKey []byte) ([]byte, error) {
	data, err := json.Marshal(u)
	c.ReqData = string(data)

	sign := Lint(u, (*c))
	if err != nil {
		return nil, err
	}
	c.Sign = tscec.Sign(privateKey, []byte(sign), false)

	reqData, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	// 校验common是否符合标准
	err = validate.Struct(c)
	if err != nil {
		return nil, err
	}

	log.Printf("trustsql request data is %s", string(reqData))

	// send http request
	req, err := http.NewRequest("POST", URI, bytes.NewBuffer(reqData))
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
	return body, nil
}
