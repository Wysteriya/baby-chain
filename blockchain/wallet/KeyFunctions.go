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
	privatekey, err := ecdsa.GenerateKey(pubkeyCurve, rand.Reader)
	return *privatekey, err
}

func VerifySignature(pub *ecdsa.PublicKey, hash []byte, sig []byte) bool {
	isValid := ecdsa.VerifyASN1(pub, hash, sig)
	return isValid
}

func SignMessage(priv *ecdsa.PrivateKey, hash []byte) ([]byte, error) {
	sign, err := ecdsa.SignASN1(rand.Reader, priv, hash)
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
	priv_d := hexutil.EncodeBig(keys.D)
	return public_addr, priv_d

}

// EncodePublic public key
// func EncodePublic(pubKey *ecdsa.PublicKey) (string, error) {
// 	encoded, err := x509.MarshalPKIXPublicKey(pubKey)
// 	if err != nil {
// 		return "", err
// 	}
// 	pemEncodedPub := pem.EncodeToMemory(&pem.Block{Bytes: encoded})
// 	return string(pemEncodedPub), nil
// }

// // DecodePrivate private key
// func DecodePrivate(pemEncodedPriv string) (*ecdsa.PrivateKey, error) {
// 	blockPriv, _ := pem.Decode([]byte(pemEncodedPriv))
// 	x509EncodedPriv := blockPriv.Bytes
// 	privateKey, err := x509.ParseECPrivateKey(x509EncodedPriv)
// 	return privateKey, err
// }

// // DecodePublic public key
// func DecodePublic(pemEncodedPub string) (*ecdsa.PublicKey, error) {
// 	blockPub, _ := pem.Decode([]byte(pemEncodedPub))
// 	x509EncodedPub := blockPub.Bytes
// 	genericPublicKey, err := x509.ParsePKIXPublicKey(x509EncodedPub)
// 	publicKey := genericPublicKey.(*ecdsa.PublicKey)
// 	return publicKey, err
// }

//func main() {

// pubkeyCurve := elliptic.P256()

// //privatekey := new(ecdsa.PrivateKey)
// privatekey, err := ecdsa.GenerateKey(pubkeyCurve, rand.Reader) // this generates a public & private key pair
// addr_x := fmt.Sprintf("%X", privatekey.PublicKey.X)
// addr_y := fmt.Sprintf("%X", privatekey.PublicKey.Y)
// addr := addr_x + addr_y
// if err != nil {
// 	fmt.Println(err)
// 	os.Exit(1)
// }

// pubkey := privatekey.PublicKey

// fmt.Println(addr)
// fmt.Println("Encoded private key ")
// fmt.Println(encoded_priv)
// fmt.Println("Private Key :")
// fmt.Printf("%X \n", *privatekey)
// fmt.Println("Decoded private key")
// fmt.Printf("%X \n", *decoded_priv)

// fmt.Println("Encoded public key ")
// fmt.Println(encoded_pub)

// // Verify
// verifystatus := ecdsa.Verify(&pubkey, signhash, r, s)
// fmt.Println(verifystatus) // should be true
//}
