package models

type ImouAccessToken struct {
	Id          string `json:"id" gorm:"primaryKey,autoIncrement"`
	AccessToken string `json:"access_token,omitempty"`
	ExpireTime  int64  `json:"expire_time,omitempty"`
}

func (ImouAccessToken) TableName() string {
	return "imou_access_token"
}
