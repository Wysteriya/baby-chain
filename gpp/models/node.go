package models

import "baby-chain/tools"

type ReceiveNode struct {
	tools.Data
}

type SendNode struct {
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
}
