package wallet

import (
	"baby-chain/tools"
	"crypto/ed25519"
	"encoding/hex"
)

func GenerateKeys() (string, string, error) {
	publicKey, privateKey, err := generateKeys()
	if err != nil {
		return "", "", err
	}
	if len(publicKey) != 32 { // unstable publicKey
		return GenerateKeys() // regenerate keys
	}
	if len(privateKey) != 64 { // unstable privateKey
		return GenerateKeys() // regenerate keys
	}
	return hex.EncodeToString(publicKey), hex.EncodeToString(privateKey), nil
}

func SignHash(_privateKey string, hash tools.Hash) ([]byte, error) {
	privateKey, err := makePrivateKey(_privateKey)
	if err != nil {
		return nil, err
	}
	return ed25519.Sign(privateKey, hash[:]), nil
}

func VerifySignature(_publicKey string, hash tools.Hash, sign []byte) bool {
	publicKey, err := makePublicKey(_publicKey)
	if err != nil {
		return false
	}
	return ed25519.Verify(publicKey, hash[:], sign)
}

func generateKeys() (publicKey ed25519.PublicKey, privateKey ed25519.PrivateKey, err error) {
	publicKey, privateKey, err = ed25519.GenerateKey(nil)
	if err != nil {
		return nil, nil, err
	}
	return publicKey, privateKey, nil
}

func makePrivateKey(_privateKey string) (ed25519.PrivateKey, error) {
	privateKey, err := hex.DecodeString(_privateKey)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func makePublicKey(_publicKey string) (ed25519.PublicKey, error) {
	publicKey, err := hex.DecodeString(_publicKey)
	if err != nil {
		return nil, err
	}
	return publicKey, nil
}
