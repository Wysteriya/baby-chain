package services

import (
	"blockchain/wallet"
	"crypto/ecdsa"
	"crypto/elliptic"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ubiq/go-ubiq/common/hexutil"
	"gpp/client_services/models"
	"net/http"
)

// CreateUser - sends the public address and private key representation
func CreateUser(ctx *gin.Context) {
	var name string = ctx.Param("name")
	returnObj := new(models.CreateUserOutput)
	public_address, priv_d := wallet.GeneratePublicAddressAndKey()

	returnObj.Name = name
	returnObj.PublicAddress = public_address
	returnObj.PrivateKey = priv_d

	ctx.IndentedJSON(http.StatusOK, returnObj)
}

func Sign(ctx *gin.Context) {
	var inputObj models.SignInputObj
	if err := ctx.BindJSON(&inputObj); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	priv_d := inputObj.PrivD
	public_address := inputObj.PublicAddress
	message_hash := inputObj.Hash
	private_key := make_keys(public_address, priv_d)
	signature, err := wallet.SignMessage(private_key, []byte(message_hash))
	if err != nil {
		fmt.Println(err.Error())
	}
	ctx.JSON(http.StatusOK, gin.H{"signature": hexutil.Encode(signature), "message_hash": string(message_hash)})
}

func IsSignValid(ctx *gin.Context) {
	var inputObj *models.ValidateSignInputObj
	if err := ctx.BindJSON(&inputObj); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	priv_d := inputObj.PrivD
	public_address := inputObj.PublicAddress
	message_hash := inputObj.Message_Hash
	signature, _ := hexutil.Decode(inputObj.Signature)
	private_key := make_keys(public_address, priv_d)
	isValid := wallet.VerifySignature(&(private_key.PublicKey), []byte(message_hash), signature)
	ctx.JSON(http.StatusOK, gin.H{"is_sign_valid": isValid})
}
func RegisterClientRoutes(rg *gin.RouterGroup) {
	clientroute := rg.Group("/clientservice")

	clientroute.GET("/getkeys:name", CreateUser)
	clientroute.POST("/signmessage", Sign)
	clientroute.POST("/validatesignature", IsSignValid)

}

//utility functions

//make_keys :convert public address and private key representation to generate private key

func make_keys(public_address string, priv_d string) *ecdsa.PrivateKey {
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
