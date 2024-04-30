package services

import (
	"IOT/initialize/psql"
	"IOT/models"
	uuid "IOT/utils"
	valid "IOT/validate"
	"errors"

	"gorm.io/gorm"
)

type TpDictService struct {
	//可搜索字段
	SearchField []string
	//可作为条件的字段
	WhereField []string
	//可做为时间范围查询的字段
	TimeField []string
}

// 获取列表
func (*TpDictService) GetTpDictList(PaginationValidate valid.TpDictPaginationValidate) (bool, []models.TpDict, int64) {
	var TpDicts []models.TpDict
	offset := (PaginationValidate.CurrentPage - 1) * PaginationValidate.PerPage
	sqlWhere := "1=1"
	if PaginationValidate.DictCode != "" {
		sqlWhere += " and dict_code = '" + PaginationValidate.DictCode + "'"
	}
	var count int64
	psql.Mydb.Model(&models.TpDict{}).Where(sqlWhere).Count(&count)
	result := psql.Mydb.Model(&models.TpDict{}).Where(sqlWhere).Limit(PaginationValidate.PerPage).Offset(offset).Order("created_at desc").Find(&TpDicts)
	if result.Error != nil {
		errors.Is(result.Error, gorm.ErrRecordNotFound)
		return false, TpDicts, 0
	}
	return true, TpDicts, count
}

// 新增数据
func (*TpDictService) AddTpDict(tp_dict models.TpDict) (bool, models.TpDict) {
	var uuid = uuid.GetUuid()
	tp_dict.ID = uuid
	result := psql.Mydb.Create(&tp_dict)
	if result.Error != nil {
		errors.Is(result.Error, gorm.ErrRecordNotFound)
		return false, tp_dict
	}
	return true, tp_dict
}

// 修改数据
func (*TpDictService) EditTpDict(tp_dict models.TpDict) bool {
	result := psql.Mydb.Updates(&tp_dict)
	//result := psql.Mydb.Save(&tp_dict)
	if result.Error != nil {
		errors.Is(result.Error, gorm.ErrRecordNotFound)
		return false
	}
	return true
}

// 删除数据
func (*TpDictService) DeleteTpDict(tp_dict models.TpDict) bool {
	result := psql.Mydb.Delete(&tp_dict)
	if result.Error != nil {
		errors.Is(result.Error, gorm.ErrRecordNotFound)
		return false
	}
	return true
}

// 条件删除
func (*TpDictService) DeleteRowTpDict(tp_dict models.TpDict) error {
	result := psql.Mydb.Exec("DELETE FROM tp_dict WHERE dict_code = ? and dict_value = ?", tp_dict.DictCode, tp_dict.DictValue)
	return result.Error
}
