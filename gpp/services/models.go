package services

type NewNodePost struct {
	UserType  string `json:"user_type"`
	IpAddress string `json:"ip_address"`
}
type NewNodeResponse struct {
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
}

type AnnounceTravelPost struct {
	FromLat string `json:"from_lat"`
	FromLon string `json:"from_lon"`
	ToLat   string `json:"to_lat"`
	ToLon   string `json:"to_lon"`
	Time    string `json:"time"`
}

type SyncPost struct {
	BlockchainData string `json:"blockchain_data"`
}
