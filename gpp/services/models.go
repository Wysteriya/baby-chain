package services

type NewNodePost struct {
	IpAddress string `json:"ip_address"`
}
type NewNodeResponse struct {
	PublicAddress string `json:"public_address"`
	PrivateKey    string `json:"private_key"`
}
