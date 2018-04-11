package tsiss

import (
	"encoding/json"
	"fmt"
)

type ret struct {
	Retcode string           `json:"retcode"`
	Retmsg  string           `json:"retmsg"`
	InfoNo  *json.RawMessage `json:"info_no"`
}

func responseUtil(body []byte) error {
	var response ret
	err := json.Unmarshal(body, &response)
	if err != nil {
		return fmt.Errorf("json unmarshal response body error: %v", err)
	}
	switch response.Retcode {
	case "0":
		return nil
	default:
		return fmt.Errorf("request error, error code is %s, msg is %s", response.Retcode, response.Retmsg)
	}
}
