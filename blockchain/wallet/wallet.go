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
	privy_d := hexutil.EncodeBig(keys.D)
	return publicAddr, privy_d
}

func generateKeys() (ecdsa.PrivateKey, error) {
	pubkeyCurve := elliptic.P256()
	privateKey, err := ecdsa.GenerateKey(pubkeyCurve, rand.Reader)
	return *privateKey, err
}

func makePrivateKey(public_address string, priv_d string) *ecdsa.PrivateKey {
	pub_x, err1 := hexutil.DecodeBig(public_address[:66])
	if err1 != nil {
		fmt.Println(err1.Error())
	}
	pub_y, err2 := hexutil.DecodeBig(public_address[66:])
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	private_d, err3 := hexutil.DecodeBig(priv_d[:])
	if err3 != nil {
		fmt.Println(err3.Error())
	}
	key_curve := elliptic.P256()
	pub_key := new(ecdsa.PublicKey)
	pub_key.Curve = key_curve
	pub_key.X = pub_x
	pub_key.Y = pub_y
	priv_key := new(ecdsa.PrivateKey)
	priv_key.PublicKey = *pub_key
	priv_key.D = private_d

	return priv_key
}

func makePublicKey(publicAddress string) *ecdsa.PublicKey {
	pub_x, err1 := hexutil.DecodeBig(publicAddress[:66])
	if err1 != nil {
		fmt.Println(err1.Error())
	}
	pub_y, err2 := hexutil.DecodeBig(publicAddress[66:])
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	key_curve := elliptic.P256()
	pub_key := new(ecdsa.PublicKey)
	pub_key.Curve = key_curve
	pub_key.X = pub_x
	pub_key.Y = pub_y
	return pub_key
}
