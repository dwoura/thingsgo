package services

import (
	tp_cron "IOT/initialize/cron"
	"IOT/initialize/psql"
	"IOT/models"
	uuid "IOT/utils"
	valid "IOT/validate"
	"errors"

	"github.com/beego/beego/v2/core/logs"
	"github.com/robfig/cron/v3"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

type TpAutomationConditionService struct {
	//可搜索字段
	SearchField []string
	//可作为条件的字段
	WhereField []string
	//可做为时间范围查询的字段
	TimeField []string
}

func (*TpAutomationConditionService) GetTpAutomationConditionDetail(tp_automation_condition_id string) (models.TpAutomationCondition, error) {
	var tp_automation_condition models.TpAutomationCondition
	result := psql.Mydb.First(&tp_automation_condition, "id = ?", tp_automation_condition_id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return tp_automation_condition, nil
		} else {
			return tp_automation_condition, result.Error
		}

	}
	return tp_automation_condition, nil
}

// 获取列表
func (*TpAutomationConditionService) GetTpAutomationConditionList(PaginationValidate valid.TpAutomationConditionPaginationValidate) (bool, []models.TpAutomationCondition, int64) {
	var TpAutomationConditions []models.TpAutomationCondition
	offset := (PaginationValidate.CurrentPage - 1) * PaginationValidate.PerPage
	sqlWhere := "1=1"
	if PaginationValidate.Id != "" {
		sqlWhere += " and id = '" + PaginationValidate.Id + "'"
	}
	var count int64
	psql.Mydb.Model(&models.TpAutomationCondition{}).Where(sqlWhere).Count(&count)
	result := psql.Mydb.Model(&models.TpAutomationCondition{}).Where(sqlWhere).Limit(PaginationValidate.PerPage).Offset(offset).Order("created_at desc").Find(&TpAutomationConditions)
	if result.Error != nil {
		logs.Error(result.Error, gorm.ErrRecordNotFound)
		return false, TpAutomationConditions, 0
	}
	return true, TpAutomationConditions, count
}

// 新增数据
func (*TpAutomationConditionService) AddTpAutomationCondition(tp_automation_condition models.TpAutomationCondition) (models.TpAutomationCondition, error) {
	var uuid = uuid.GetUuid()
	tp_automation_condition.Id = uuid
	result := psql.Mydb.Create(&tp_automation_condition)
	if result.Error != nil {
		logs.Error(result.Error, gorm.ErrRecordNotFound)
		return tp_automation_condition, result.Error
	}
	return tp_automation_condition, nil
}

// 修改数据
func (*TpAutomationConditionService) EditTpAutomationCondition(tp_automation_condition valid.TpAutomationConditionValidate) bool {
	result := psql.Mydb.Model(&models.TpAutomationCondition{}).Where("id = ?", tp_automation_condition.Id).Updates(&tp_automation_condition)
	if result.Error != nil {
		logs.Error(result.Error, gorm.ErrRecordNotFound)
		return false
	}
	return true
}

// 删除数据
func (*TpAutomationConditionService) DeleteById(automationConditionId string) error {
	result := psql.Mydb.Where("id = ?", automationConditionId).Delete(&models.TpAutomationCondition{})
	if result.Error != nil {
		logs.Error(result.Error)
		return result.Error
	}
	return nil
}

// 根据自动化Id删除自动化条件中的定时任务
func (*TpAutomationConditionService) DeleteCronsByAutomationId(automationId string) error {
	var automationConditions []models.TpAutomationCondition
	result := psql.Mydb.Model(&models.TpAutomationCondition{}).Where("automation_id = ? and condition_type = '2' and time_condition_type = '2'", automationId).Find(&automationConditions)
	if result.Error != nil {
		logs.Error(result.Error.Error())
		return result.Error
	}
	for _, automationCondition := range automationConditions {
		cronId := cast.ToInt(automationCondition.V2)
		C := tp_cron.C
		C.Remove(cron.EntryID(cronId))
	}
	return nil
}
