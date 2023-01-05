package models

type ReceiveNode struct {
	Driver      string `json:"driver"`
	LicenceNum  string `json:"licence_num"`
	VehicleType string `json:"vehicle_type"`

	BloodGroup string `json:"blood_group"`
	Age        string `json:"age"`
	Gender     string `json:"gender"`
}

type SendNode struct {
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
}
