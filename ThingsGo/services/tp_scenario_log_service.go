package services

import (
	"IOT/initialize/psql"
	"IOT/models"
	uuid "IOT/utils"
	valid "IOT/validate"

	"github.com/beego/beego/v2/core/logs"
	"gorm.io/gorm"
)

type TpScenarioLogService struct {
	//可搜索字段
	SearchField []string
	//可作为条件的字段
	WhereField []string
	//可做为时间范围查询的字段
	TimeField []string
}

// 获取列表
func (*TpScenarioLogService) GetTpScenarioLogList(PaginationValidate valid.TpScenarioLogPaginationValidate) ([]models.TpScenarioLog, int64, error) {
	var TpScenarioLogs []models.TpScenarioLog
	offset := (PaginationValidate.CurrentPage - 1) * PaginationValidate.PerPage
	sqlWhere := "1=1"
	var paramList []interface{}
	if PaginationValidate.Id != "" {
		sqlWhere += " and id = ?"
		paramList = append(paramList, PaginationValidate.Id)
	}
	if PaginationValidate.ProcessResult != "" {
		sqlWhere += " and process_result = ?"
		paramList = append(paramList, PaginationValidate.ProcessResult)
	}
	if PaginationValidate.ScenarioStrategyId != "" {
		sqlWhere += " and scenario_strategy_id = ?"
		paramList = append(paramList, PaginationValidate.ScenarioStrategyId)
	}
	var count int64
	psql.Mydb.Model(&models.TpScenarioLog{}).Where(sqlWhere, paramList...).Count(&count)
	result := psql.Mydb.Model(&models.TpScenarioLog{}).Where(sqlWhere, paramList...).Limit(PaginationValidate.PerPage).Offset(offset).Order("trigger_time desc").Find(&TpScenarioLogs)
	if result.Error != nil {
		logs.Error(result.Error, gorm.ErrRecordNotFound)
		return TpScenarioLogs, 0, result.Error
	}
	return TpScenarioLogs, count, nil
}

// 新增数据
func (*TpScenarioLogService) AddTpScenarioLog(scenarioLog models.TpScenarioLog) (models.TpScenarioLog, error) {
	var uuid = uuid.GetUuid()
	scenarioLog.Id = uuid
	result := psql.Mydb.Create(&scenarioLog)
	if result.Error != nil {
		logs.Error(result.Error, gorm.ErrRecordNotFound)
		return scenarioLog, result.Error
	}
	return scenarioLog, nil
}

// 修改数据
func (*TpScenarioLogService) UpdateTpScenarioLog(scenarioLog models.TpScenarioLog) (models.TpScenarioLog, error) {
	result := psql.Mydb.Model(&models.TpScenarioLog{}).Where("id = ?", scenarioLog.Id).Updates(&scenarioLog)
	if result.Error != nil {
		logs.Error(result.Error, gorm.ErrRecordNotFound)
		return scenarioLog, result.Error
	}
	return scenarioLog, nil
}
