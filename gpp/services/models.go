package services

type NewNodePost struct {
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

type PublicInfoPost struct {
	IpAddress  string `json:"ip_address"`
	Name       string `json:"name"`
	Driver     bool   `json:"driver"`
	Lisense    string `json:"lisense"`
	Email      string `json:"email"`
	Occupation string `json:"occupation"`
	Hobbies    string `json:"hobbies"`
	Skills     string `json:"skills"`
	Interests  string `json:"interests"`
	Others     string `json:"others"`
}
