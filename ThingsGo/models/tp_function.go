package models

type TpFunction struct {
	Id           string `json:"id,omitempty" gorm:"primaryKey"`          // ID
	FunctionName string `json:"function_name,omitempty" gorm:"size:99"`  // 功能名称
	Path         string `json:"path,omitempty" gorm:"size:255"`          //
	Name         string `json:"name,omitempty" gorm:"size:255"`          //
	Component    string `json:"component,omitempty" gorm:"size:255"`     //
	Title        string `json:"title,omitempty" gorm:"size:255"`         //
	Icon         string `json:"icon,omitempty" gorm:"size:255"`          //
	Type         string `json:"type,omitempty" gorm:"size:255"`          //
	FunctionCode string `json:"function_code,omitempty" gorm:"size:255"` //
	ParentId     string `json:"parent_id,omitempty" gorm:"size:36"`      //
	Sort         int    `json:"sort,omitempty"`                          //
}

func (TpFunction) TableName() string {
	return "tp_function"
}
