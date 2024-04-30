package services

import (
	"IOT/initialize/psql"
	"IOT/models"
	uuid "IOT/utils"
	valid "IOT/validate"

	"github.com/beego/beego/v2/core/logs"
	"gorm.io/gorm"
)

type TpScriptService struct {
	//可搜索字段
	SearchField []string
	//可作为条件的字段
	WhereField []string
	//可做为时间范围查询的字段
	TimeField []string
}

func (*TpScriptService) GetTpScriptDetail(tp_script_id string) []models.TpScript {
	var tp_script []models.TpScript
	psql.Mydb.First(&tp_script, "id = ?", tp_script_id)
	return tp_script
}

// 获取列表
func (*TpScriptService) GetTpScriptList(PaginationValidate valid.TpScriptPaginationValidate) (bool, []models.TpScript, int64) {
	var TpScripts []models.TpScript
	offset := (PaginationValidate.CurrentPage - 1) * PaginationValidate.PerPage
	sqlWhere := "1=1"
	if PaginationValidate.ProtocolType != "" {
		sqlWhere += " and protocol_type = '" + PaginationValidate.ProtocolType + "'"
	}
	if PaginationValidate.Id != "" {
		sqlWhere += " and id = '" + PaginationValidate.Id + "'"
	}
	if PaginationValidate.DeviceType != "" {
		sqlWhere += " and device_type = '" + PaginationValidate.DeviceType + "'"
	}
	var count int64
	psql.Mydb.Model(&models.TpScript{}).Where(sqlWhere).Count(&count)
	result := psql.Mydb.Model(&models.TpScript{}).Where(sqlWhere).Limit(PaginationValidate.PerPage).Offset(offset).Order("created_at desc").Find(&TpScripts)
	if result.Error != nil {
		logs.Error(result.Error, gorm.ErrRecordNotFound)
		return false, TpScripts, 0
	}
	return true, TpScripts, count
}

// 新增数据
func (*TpScriptService) AddTpScript(tp_script models.TpScript) (models.TpScript, error) {
	var uuid = uuid.GetUuid()
	tp_script.Id = uuid
	result := psql.Mydb.Create(&tp_script)
	if result.Error != nil {
		logs.Error(result.Error, gorm.ErrRecordNotFound)
		return tp_script, result.Error
	}
	return tp_script, nil
}

// 修改数据
func (*TpScriptService) EditTpScript(tp_script valid.TpScriptValidate) bool {
	result := psql.Mydb.Model(&models.TpScript{}).Where("id = ?", tp_script.Id).Updates(&tp_script)
	if result.Error != nil {
		logs.Error(result.Error, gorm.ErrRecordNotFound)
		return false
	}
	return true
}

// 删除数据
func (*TpScriptService) DeleteTpScript(tp_script models.TpScript) error {
	result := psql.Mydb.Delete(&tp_script)
	if result.Error != nil {
		logs.Error(result.Error, gorm.ErrRecordNotFound)
		return result.Error
	}
	return nil
}
