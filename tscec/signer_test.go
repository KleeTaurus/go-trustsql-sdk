package tscec

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestSign(t *testing.T) {
	fmt.Println("\n======test sign begin==========")
	privateKey, publicKey := NewKeyPair()
	fmt.Printf("privkey:   %s\n", base64.StdEncoding.EncodeToString(privateKey))
	fmt.Printf("pubkey:    %s\n", base64.StdEncoding.EncodeToString(publicKey))

	signMsg := "hello world, you are welcome"
	ret := Sign(privateKey, []byte(signMsg))

	fmt.Printf("data:      %s\n", signMsg)
	fmt.Printf("signature: %s\n", ret)
	fmt.Println("======test sign end==========")
}

func TestVerify(t *testing.T) {
	message := "hello world, you are welcome"
	sig := "MEUCIQCmKnxXF32Ni5/jWHYcBn57fvjXIF4+fvL/Tix+LDItvAIgN00bzhgpd/eMjOteN+SfsqpdnXJFepWToZ37VrY5qzE="
	privateKey := "OvxHYMWUE31PtBohtbXdF9fOXafS2fe3GoZtE9SVG2I="
	pubKey := "BBJjY/x2Patnk3aFCh/4u3q5p8c30RUiKiVHBl4Mg4w1SvMprrS26Z47WaXEj1Fe68deacd63mWTImOxupoNvWI="

	fmt.Println("\n======test verify begin==========")
	fmt.Printf("privkey:   %s\n", privateKey)
	fmt.Printf("pubkey:    %s\n", pubKey)
	fmt.Printf("data:      %s\n", message)
	fmt.Printf("signature: %s\n", sig)

	pkey, _ := base64.StdEncoding.DecodeString(pubKey)
	signature, _ := base64.StdEncoding.DecodeString(sig)

	ret := Verify(pkey, signature, []byte(message))

	fmt.Printf("return:    %t\n", ret)
	fmt.Println("======test verify end==========")
}
