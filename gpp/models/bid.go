package models

type ReceiveBid struct {
	AnnouncementId string `json:"announcement_id"`
	BidAmount      string `json:"bid_amount"`
	PublicKey      string `json:"public_key"`
	PrivateKey     string `json:"private_key"`
}

type SendBid struct {
	BidBlockData string
}
