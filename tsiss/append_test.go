package tsiss

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestAppendIss(t *testing.T) {
	issAppend := &IssAppend{
		Version:     "1.0",
		SignType:    "ECDSA",
		MchID:       "gbec7b7cece75c8a5",
		MchSign:     "MEYCIQDCoCYth2zGer2Z/kliD11jRXGKqLqLNk/vo18js+CvRwIhANTQ3PbN9vj9YjmaB+rma2Sz0D+30WgZPOHAO9ysRsj1",
		Account:     "1PPdKF7brmyiQF7zKtGe1KnEURGRb1jBE2",
		CommitTime:  "2017-08-20 15:00:00",
		Content:     map[string]interface{}{"content": "test"},
		InfoKey:     "teast123",
		InfoVersion: "1",
		State:       "0",
		Notes:       map[string]interface{}{"notes": "test"},
		PublicKey:   "BNTUTn8geMERremCxomOVXSGLitaGe8FnjpRNXCLiNtd8yl9G6vtqCyGlmpEz901t6WVHItMqZ9ozt5o/Xf6uL4=",
		Sign:        "MEQCIE4YbWYw4FyUMd12HHJqsAGfor9xKdb7e0feM+G0yklAAiBEy1wk1MFL5ZcrySzvO9g7pYVVDkJwKrwZ56Gdlzlybg==",

		ChainID:  "xxx test chain id",
		LedgerID: "xxx test ledger_id",
	}
	iss, err := AppendIss(issAppend)
	if err != nil {
		t.Error(err)
	}
	t.Log(iss.MchID)
}

func TestAppendIssValidate(t *testing.T) {
	dat, err := ioutil.ReadFile("test_data/append_response.json")
	if err != nil {
		t.Error(err)
	}
	// fmt.Print(string(dat))
	issResponse := IssAppendResponse{}
	json.Unmarshal(dat, &issResponse)
	t.Log(issResponse.MchID)
}
