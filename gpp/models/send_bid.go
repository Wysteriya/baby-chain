package models

type ReceiveBid struct {
	AnnouncementId string `json:announcement_id`
	BidAmount      string `json:bid_amount`
}
