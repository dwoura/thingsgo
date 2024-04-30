package models

type ImouKitToken struct {
	Id       string `json:"id" gorm:"primaryKey"`
	DeviceId string `json:"device_id,omitempty"`
	KitToken string `json:"kit_token,omitempty"`
}

func (ImouKitToken) TableName() string {
	return "imou_kit_token"
}
