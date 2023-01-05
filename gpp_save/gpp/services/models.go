package services

type AnnounceTravelPost struct {
	FromLat    string `json:"from_lat"`
	FromLon    string `json:"from_lon"`
	ToLat      string `json:"to_lat"`
	ToLon      string `json:"to_lon"`
	Time       string `json:"time"`
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
}

type SyncPost struct {
	BlockchainData string `json:"blockchain_data"`
}

type PublicInfoPost struct {
	IpAddress  string `json:"ip_address"`
	Name       string `json:"name"`
	Driver     bool   `json:"driver"`
	License    string `json:"license"`
	Email      string `json:"email"`
	Occupation string `json:"occupation"`
	Hobbies    string `json:"hobbies"`
	Skills     string `json:"skills"`
	Interests  string `json:"interests"`
	Others     string `json:"others"`
}

type MessagePost struct {
	Type                  string `json:"type"`
	PublicKey             string `json:"public_key"`
	PrivateKeyOrSignature string `json:"key"`
	ToPublicKey           string `json:"to_public_key"`
	Message               string `json:"message"`
}

type SendBidPost struct {
	AnnouncementId string `json:"announcement_id"`
	BidAmount      string `json:"bid_amount"`
	PublicKey      string `json:"public_key"`
	PrivateKey     string `json:"private_key"`
}
