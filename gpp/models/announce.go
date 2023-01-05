package models

type ReceiveAnnouncement struct {
	FromLat    string `json:"from_lat"`
	FromLon    string `json:"from_lon"`
	ToLat      string `json:"to_lat"`
	ToLon      string `json:"to_lon"`
	Time       string `json:"time"`
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
}

type SendAnnouncement struct {
	AnnounceBlockData string
}
