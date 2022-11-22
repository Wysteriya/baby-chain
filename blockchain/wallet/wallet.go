package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"github.com/ubiq/go-ubiq/common/hexutil"
)

func GenerateKeys() (ecdsa.PrivateKey, error) {
	pubkeyCurve := elliptic.P256()
	privateKey, err := ecdsa.GenerateKey(pubkeyCurve, rand.Reader)
	return *privateKey, err
}

func VerifySignature(pub *ecdsa.PublicKey, hash []byte, sig []byte) bool {
	isValid := ecdsa.VerifyASN1(pub, hash, sig)
	return isValid
}

func SignMessage(privy *ecdsa.PrivateKey, hash []byte) ([]byte, error) {
	sign, err := ecdsa.SignASN1(rand.Reader, privy, hash)
	return sign, err
}

func GeneratePublicAddressAndKey() (public_address string, private_key string) {
	keys, err := GenerateKeys()
	if err != nil {
		fmt.Println("error generating keys")
	}
	addr_x := hexutil.EncodeBig(keys.PublicKey.X)
	addr_y := hexutil.EncodeBig(keys.PublicKey.Y)
	public_addr := addr_x + addr_y
	privy_d := hexutil.EncodeBig(keys.D)
	return public_addr, privy_d
}
