package services

type NewNodePost struct {
	IpAddress string `json:"ip_address"`
}
type NewNodeResponse struct {
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
}
