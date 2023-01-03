package wallet

import (
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
)

func GeneratePublicAddressAndKey() (string, string, error) {
	publicKey, privateKey, err := generateKeys()
	if err != nil {
		return "", "", err
	}
	return hex.EncodeToString(publicKey), hex.EncodeToString(privateKey), nil
}

func SignMessage(privateKey_ string, hash []byte) []byte {
	privateKey := makePrivateKey(privateKey_)
	sign := ed25519.Sign(privateKey, hash)
	return sign
}

func VerifySignature(publicKey_ string, hash []byte, sig []byte) bool {
	publicKey := makePublicKey(publicKey_)
	isValid := ed25519.Verify(publicKey, hash, sig)
	return isValid
}

func generateKeys() (publicKey ed25519.PublicKey, privateKey ed25519.PrivateKey, err error) {
	publicKey, privateKey, err = ed25519.GenerateKey(nil)
	if err != nil {
		return nil, nil, err
	}
	return publicKey, privateKey, nil
}

func makePrivateKey(_privateKey string) ed25519.PrivateKey {
	privateKey, err1 := hex.DecodeString(_privateKey)
	if err1 != nil {
		fmt.Println(err1.Error())
	}
	return privateKey
}

func makePublicKey(publicAddress string) ed25519.PublicKey {
	publicKey, err1 := hex.DecodeString(publicAddress)
	if err1 != nil {
		fmt.Println(err1.Error())
	}
	return publicKey
}
