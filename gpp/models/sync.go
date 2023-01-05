package models

type ReceiveSync struct {
	BlockchainData string `json:"blockchain_data"`
	Type           string `json:"type"`
}

type SendSync struct {
	BlockchainData string `json:"blockchain_data"`
	Type           string `json:"type"`
}
