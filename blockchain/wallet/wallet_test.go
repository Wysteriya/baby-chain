package wallet

import (
	"blockchain/block"
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	publicKey, privateKey, _ := GeneratePublicAddressAndKey()
	fmt.Println(publicKey)
	fmt.Println(privateKey)
	if len(publicKey) != 64 {
		t.Fatalf("unstable publicKey")
	} else if len(privateKey) != 128 {
		t.Fatalf("unstable privateKey")
	}
	hash := block.HashB([]byte("Test"))
	signature := SignMessage(privateKey, hash[:])
	if !VerifySignature(publicKey, hash[:], signature) {
		t.Fatalf("valid signature is not being validated")
	}
}
