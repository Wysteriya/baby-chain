package models

type ReceiveNode struct {
	Driver      string `json:"driver"`
	LicenceNum  string `json:"licence-num"`
	VehicleType string `json:"vehicle-type"`

	BloodGroup string `json:"blood-group"`
	Age        string `json:"age"`
	Gender     string `json:"gender"`
}

type SendNode struct {
	PublicKey  string `json:"public-key"`
	PrivateKey string `json:"private-key"`
}
