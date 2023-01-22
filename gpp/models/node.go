package models

import (
	"baby-chain/tools/data"
)

type ReceiveNode struct {
	data.Data
}

type SendNode struct {
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
}
