package models

type ImouPicture struct {
	Id         string `json:"id" gorm:"primaryKey,autoIncrement"`
	DeviceId   string `json:"device_id,omitempty"`
	Url        string `json:"url,omitempty"`
	PicName    string `json:"pic_name,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

func (ImouPicture) TableName() string {
	return "imou_picture"
}
