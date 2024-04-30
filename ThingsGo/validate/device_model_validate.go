package valid

import "IOT/models"

type DeviceModelValidate struct {
	Id        string `json:"id" alias:"ID" valid:"Required;MaxSize(36)"`
	ModelName string `json:"model_name,omitempty" alias:"插件名称" valid:"MaxSize(99)"`
	Flag      int64  `json:"flag,omitempty" alias:"插件标志" `
	ChartData string `json:"chart_data,omitempty" alias:"插件json" `
	ModelType string `json:"model_type,omitempty" alias:"插件类型" `                 // 插件类型
	Describe  string `json:"describe,omitempty" alias:"描述" valid:"MaxSize(255)"` // 描述
	Version   string `json:"version,omitempty" alias:"版本" valid:"MaxSize(36)"`   // 版本
	Author    string `json:"author,omitempty" alias:"坐着" valid:"MaxSize(36)"`
	Sort      int64  `json:"sort,omitempty" alias:"排序" `
	Issued    int64  `json:"issued,omitempty" alias:"发布" `
	Remark    string `json:"remark,omitempty" alias:"备注" valid:"MaxSize(255)"`
	CreatedAt int64  `json:"created_at,omitempty" alias:"创建时间" `
}

type AddDeviceModelValidate struct {
	ModelName string `json:"model_name,omitempty" alias:"插件名称" valid:"MaxSize(99)"`
	Flag      int64  `json:"flag,omitempty" alias:"插件标志" `
	ChartData string `json:"chart_data,omitempty" alias:"插件json" `
	ModelType string `json:"model_type,omitempty" alias:"插件类型" `                 // 插件类型
	Describe  string `json:"describe,omitempty" alias:"描述" valid:"MaxSize(255)"` // 描述
	Version   string `json:"version,omitempty" alias:"版本" valid:"MaxSize(36)"`   // 版本
	Author    string `json:"author,omitempty" alias:"坐着" valid:"MaxSize(36)"`
	Sort      int64  `json:"sort,omitempty" alias:"排序" `
	Issued    int64  `json:"issued,omitempty" alias:"发布" `
	Remark    string `json:"remark,omitempty" alias:"备注" valid:"MaxSize(255)"`
	CreatedAt int64  `json:"created_at,omitempty" alias:"创建时间" `
}

type DeviceModelPaginationValidate struct {
	CurrentPage int    `json:"current_page"  alias:"当前页" valid:"Required;Min(1)"`
	PerPage     int    `json:"per_page"  alias:"每页页数" valid:"Required;Max(10000)"`
	Issued      int    `json:"issued" alias:"发布状态" valid:"Max(36)"`
	ModelType   string `json:"model_type" alias:"插件类型" valid:"MaxSize(36)"`
	Id          string `json:"id" alias:"id" valid:"MaxSize(36)"`
	Flag        int    `json:"flag" alias:"标志" valid:"Max(36)"`
}

type RspDeviceModelPaginationValidate struct {
	CurrentPage int                  `json:"current_page"  alias:"当前页" valid:"Required;Min(1)"`
	PerPage     int                  `json:"per_page"  alias:"每页页数" valid:"Required;Max(10000)"`
	Data        []models.DeviceModel `json:"data" alias:"返回数据"`
	Total       int64                `json:"total" alias:"总数" valid:"Max(10000)"`
}
