package wallet

import (
	"baby-chain/tools"
	"testing"
)

func TestKeys(t *testing.T) {
	_publicKey, _privateKey, err := GenerateKeys()
	tools.TError(err, t)
	t.Logf("publicKey: %s\nprivateKey: %s", _publicKey, _privateKey)
}

func TestSignature(t *testing.T) {
	_publicKey, _privateKey, err := GenerateKeys()
	tools.TError(err, t)

	hash := tools.HashB([]byte("Test"))
	falseHash := tools.HashB()
	signature, err := SignHash(_privateKey, hash)
	tools.TError(err, t)
	falseSignature, err := SignHash(_privateKey, falseHash)
	tools.TError(err, t)

	if !VerifySignature(_publicKey, hash, signature) {
		t.Fatalf("valid signature is not being validated")
	}
	if VerifySignature(_publicKey, falseHash, signature) {
		t.Fatalf("in valid signature is being validated")
	}
	if VerifySignature(_publicKey, hash, falseSignature) {
		t.Fatalf("in valid signature is being validated")
	}
}
