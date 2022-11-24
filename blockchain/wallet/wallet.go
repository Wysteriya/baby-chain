package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
)

func VerifySignature(publicKey_ string, hash []byte, sig []byte) bool {
	publicKey := makePublicKey(publicKey_)
	isValid := ecdsa.VerifyASN1(publicKey, hash, sig)
	return isValid
}

func SignMessage(privateKey_ string, publicKey string, hash []byte) ([]byte, error) {
	privateKey := makePrivateKey(publicKey, privateKey_)
	sign, err := ecdsa.SignASN1(rand.Reader, privateKey, hash)
	return sign, err
}

func GeneratePublicAddressAndKey() (string, string, error) {
	keys, err := generateKeys()
	if err != nil {
		return "", "", err
	}
	addrX := keys.PublicKey.X.Bytes()
	addrY := keys.PublicKey.Y.Bytes()
	publicAddr := append(addrX, addrY...)
	privyD := keys.D.Bytes()
	return hex.EncodeToString(publicAddr), hex.EncodeToString(privyD), nil
}

func generateKeys() (ecdsa.PrivateKey, error) {
	publicKeyCurve := elliptic.P256()
	privateKey, err := ecdsa.GenerateKey(publicKeyCurve, rand.Reader)
	return *privateKey, err
}

func makePrivateKey(publicAddress string, privyD string) *ecdsa.PrivateKey {
	pubX, err1 := hex.DecodeString(publicAddress[:64])
	if err1 != nil {
		fmt.Println(err1.Error())
	}
	pubY, err2 := hex.DecodeString(publicAddress[64:])
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	privateD, err3 := hex.DecodeString(privyD[:])
	if err3 != nil {
		fmt.Println(err3.Error())
	}
	pubKeyX := new(big.Int)
	pubKeyY := new(big.Int)
	privateKeyD := new(big.Int)

	pubKeyX.SetBytes(pubX)
	pubKeyY.SetBytes(pubY)
	privateKeyD.SetBytes(privateD)

	pubKey := new(ecdsa.PublicKey)
	keyCurve := elliptic.P256()
	pubKey.Curve = keyCurve
	pubKey.X = pubKeyX
	pubKey.Y = pubKeyY

	privyKey := new(ecdsa.PrivateKey)
	privyKey.PublicKey = *pubKey
	privyKey.D = privateKeyD

	return privyKey
}

func makePublicKey(publicAddress string) *ecdsa.PublicKey {
	pubX, err1 := hex.DecodeString(publicAddress[:64])
	if err1 != nil {
		fmt.Println(err1.Error())
	}
	pubY, err2 := hex.DecodeString(publicAddress[64:])
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	pubKeyX := new(big.Int)
	pubKeyY := new(big.Int)

	pubKeyX.SetBytes(pubX)
	pubKeyY.SetBytes(pubY)

	pubKey := new(ecdsa.PublicKey)
	keyCurve := elliptic.P256()
	pubKey.Curve = keyCurve
	pubKey.X = pubKeyX
	pubKey.Y = pubKeyY

	return pubKey
}
