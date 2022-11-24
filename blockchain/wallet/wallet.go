package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"github.com/ubiq/go-ubiq/common/hexutil"
)

func VerifySignature(publicAddress string, hash []byte, sig []byte) bool {
	publicKey := makePublicKey(publicAddress)
	isValid := ecdsa.VerifyASN1(publicKey, hash, sig)
	return isValid
}

func SignMessage(privateKeyString string, publicAddress string, hash []byte) ([]byte, error) {
	privateKey := makePrivateKey(publicAddress, privateKeyString)
	sign, err := ecdsa.SignASN1(rand.Reader, privateKey, hash)
	return sign, err
}

func GeneratePublicAddressAndKey() (publicAddress string, privateKey string) {
	keys, err := generateKeys()
	if err != nil {
		fmt.Println("error generating keys")
	}
	addrX := hexutil.EncodeBig(keys.PublicKey.X)
	addrY := hexutil.EncodeBig(keys.PublicKey.Y)
	publicAddr := addrX + addrY
	privyD := hexutil.EncodeBig(keys.D)
	return publicAddr, privyD
}

func generateKeys() (ecdsa.PrivateKey, error) {
	publicKeyCurve := elliptic.P256()
	privateKey, err := ecdsa.GenerateKey(publicKeyCurve, rand.Reader)
	return *privateKey, err
}

func makePrivateKey(publicAddress string, privyD string) *ecdsa.PrivateKey {
	pubX, err1 := hexutil.DecodeBig(publicAddress[:66])
	if err1 != nil {
		fmt.Println(err1.Error())
	}
	pubY, err2 := hexutil.DecodeBig(publicAddress[66:])
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	privateD, err3 := hexutil.DecodeBig(privyD[:])
	if err3 != nil {
		fmt.Println(err3.Error())
	}
	keyCurve := elliptic.P256()
	pubKey := new(ecdsa.PublicKey)
	pubKey.Curve = keyCurve
	pubKey.X = pubX
	pubKey.Y = pubY
	privyKey := new(ecdsa.PrivateKey)
	privyKey.PublicKey = *pubKey
	privyKey.D = privateD

	return privyKey
}

func makePublicKey(publicAddress string) *ecdsa.PublicKey {
	pubX, err1 := hexutil.DecodeBig(publicAddress[:66])
	if err1 != nil {
		fmt.Println(err1.Error())
	}
	pubY, err2 := hexutil.DecodeBig(publicAddress[66:])
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	keyCurve := elliptic.P256()
	pubKey := new(ecdsa.PublicKey)
	pubKey.Curve = keyCurve
	pubKey.X = pubX
	pubKey.Y = pubY
	return pubKey
}
