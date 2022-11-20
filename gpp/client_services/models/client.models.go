package models

type CreateUserOutput struct {
	Name          string `json:"user_name"`
	PublicAddress string `json:"public_address"`
	PrivateKey    string `json:"private_key"`
}

type SignInputObj struct {
	Hash          string `json:"message_hash"`
	PublicAddress string `json:"public_address"`
	PrivD         string `json:"priv_d"`
}

type ValidateSignInputObj struct {
	Message_Hash  string `json:"message_hash"`
	PublicAddress string `json:"public_address"`
	PrivD         string `json:"priv_d"`
	Signature     string `json:"signature"`
}
