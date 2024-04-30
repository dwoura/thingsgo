package models

type Asset struct {
	ID             string `json:"id" gorm:"primarykey"`
	AdditionalInfo string `json:"additional_info" gorm:"type:longtext"`
	CustomerID     string `json:"customer_id" gorm:"size:36"` // 客户ID
	Name           string `json:"name"`                       // 名称
	Label          string `json:"label"`                      // 标签
	SearchText     string `json:"search_text"`
	Type           string `json:"type"`                       // 类型
	ParentID       string `json:"parent_id" gorm:"size:36"`   // 父级ID
	Tier           int64  `json:"tier"`                       // 层级
	BusinessID     string `json:"business_id" gorm:"size:36"` // 业务ID
}

type Simple struct {
	ID   string `json:"id" gorm:"primarykey"`
	Name string `json:"name"` // 名称
}

type ExtensionFields struct {
	Key    string `json:"key"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	Type   string `json:"type"`
}
type ExtensionDataMap struct {
	Name  string            `json:"name"`
	Field []ExtensionFields `json:"field"`
}

func (Asset) TableName() string {
	return "asset"
}
