package trustsql

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestSign(t *testing.T) {
	keyPair := NewKeyPair()
	fmt.Printf("privkey: %s\n", base64.StdEncoding.EncodeToString(keyPair.PrivateKey))

	signMsg := "hello world, you are welcome"
	Sign(keyPair.PrivateKey, []byte(signMsg))
}
