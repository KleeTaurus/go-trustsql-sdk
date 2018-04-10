// +build ignore

package tsiss

import (
	"fmt"
)

func ExamplAppendIss() {
	ia := &IssAppend{
		Version:  "",
		SignType: "",
		MchID:    "",
		MchSign:  "",

		NodeID:      "",
		ChainID:     "",
		LedgerID:    "",
		InfoKey:     "",
		InfoVersion: "",
		State:       "",

		CommitTime: "",
		Account:    "",
		PublicKey:  "",
		Sign:       "",
	}

	isr, err := AppendIss(ia)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(isr)
}
