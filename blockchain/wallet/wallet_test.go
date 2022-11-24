package wallet

import (
	"blockchain/block"
	"testing"
)

func Test(t *testing.T) {
	publicKey, privateKey, _ := GeneratePublicAddressAndKey()
	if len(publicKey) != 128 {
		t.Fatalf("unstable publicKey")
	} else if len(privateKey) != 64 {
		t.Fatalf("unstable privateKey")
	}
	hash := block.HashB([]byte("Test"))
	signature, err := SignMessage(privateKey, publicKey, hash[:])
	if err != nil {
		t.Fatal(err)
	} else if !VerifySignature(publicKey, hash[:], signature) {
		t.Fatalf("valid signature is not being validated")
	}
}
