package tsiss

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestGetPubkeyOfAccount(t *testing.T) {
	issQuery := IssQuery{
		Version:   "1.0",
		SignType:  "ECDSA",
		MchID:     "gbec7b7cece75c8a5",
		MchSign:   "MEYCIQDCoCYth2zGer2Z/kliD11jRXGKqLqLNk/vo18js+CvRwIhANTQ3PbN9vj9YjmaB+rma2Sz0D+30WgZPOHAO9ysRsj1",
		ChainID:   "aa",
		LedgerID:  "bb",
		Timestamp: 1503648096,
		NodeID:    "cc",
	}
	_, err := QueryIss(issQuery)
	if err != nil {
		t.Error(err)
	}
	// fmt.Println(body)
}

func TestIssResponse(t *testing.T) {
	dat, err := ioutil.ReadFile("tsiss_query_response.json")
	if err != nil {
		t.Error(err)
	}
	// fmt.Print(string(dat))
	issResponse := IssResponse{}
	json.Unmarshal(dat, &issResponse)
	t.Log(issResponse.MchID)
	t.Log(issResponse.Infos[0].InfoKey)
}
