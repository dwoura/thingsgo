package cron

import (
	"IOT/initialize/psql"
	"IOT/models"
	"IOT/services"
	"fmt"
	"time"

	"github.com/beego/beego/v2/core/logs"
	"github.com/robfig/cron/v3"

	tp_cron "IOT/initialize/cron"
)

//var C *cron.Cron

func init() {
	fmt.Println("定时任务初始化开始")
	onceCron()
	automationCron()
	fmt.Println("定时任务初始化完成")
}

func automationCron() {
	C := tp_cron.C
	//C = cron.New()
	var automationConditions []models.TpAutomationCondition
	result := psql.Mydb.Table("tp_automation").
		Select("tp_automation_condition.*").
		Joins("left join tp_automation_condition on tp_automation.id = tp_automation_condition.automation_id").
		Where("tp_automation.enabled = '1' and tp_automation_condition.condition_type = '2' and tp_automation_condition.time_condition_type = '2'").
		Order("tp_automation.priority asc").
		Find(&automationConditions)
	if result.Error != nil {
		logs.Error(result.Error.Error())
		logs.Error("定时任务初始化失败！")
	}
	for _, automationCondition := range automationConditions {
		services.AutomationCron(automationCondition)
	}
	C.Start()
}

func onceCron() {
	//c = cron.New(cron.WithSeconds())
	crontab := cron.New()
	spec := "0/1 * * * *" //每分钟一次
	task := func() {
		fmt.Println("检查单次定时任务开始")
		format := "2006/01/02 15:04:05"
		now, _ := time.Parse(format, time.Now().Format(format))
		var automationConditions []models.TpAutomationCondition
		result := psql.Mydb.Model(&models.TpAutomationCondition{}).Where("condition_type = '2' and time_condition_type = '1' and v1 != '' and v1 < ?", now).Find(&automationConditions)
		if result.Error != nil {
			logs.Error(result.Error.Error())
			return
		}
		for _, automationCondition := range automationConditions {
			// 触发，记录日志
			var logMessage string = "触发" + automationCondition.V1 + "的定时任务;"
			var sutomationLogService services.TpAutomationLogService
			var automationLog models.TpAutomationLog
			automationLog.AutomationId = automationCondition.AutomationId
			automationLog.ProcessDescription = logMessage
			automationLog.TriggerTime = time.Now().Format("2006/01/02 15:04:05")
			automationLog.ProcessResult = "2"
			automationLog, err := sutomationLogService.AddTpAutomationLog(automationLog)
			if err != nil {
				logs.Error(err.Error())
			} else {
				var automationLogMap = make(map[string]interface{})
				automationLogMap["Id"] = automationLog.Id
				var conditionsService services.ConditionsService
				msg, err := conditionsService.ExecuteAutomationAction(automationCondition.AutomationId, automationLog.Id)
				if err != nil {
					//执行失败，记录日志
					logs.Error(err.Error())
					automationLogMap["ProcessDescription"] = logMessage + err.Error()
				} else {
					//执行成功，记录日志
					logs.Info(logMessage)
					automationLogMap["ProcessDescription"] = logMessage + msg
					automationLogMap["ProcessResult"] = "1"
				}
				logs.Warn(automationLogMap)
				err = sutomationLogService.UpdateTpAutomationLog(automationLogMap)
				if err != nil {
					logs.Error(err.Error())
				}
			}
			//删除条件
			var automationConditionService services.TpAutomationConditionService
			err = automationConditionService.DeleteById(automationCondition.Id)
			if err != nil {
				logs.Error(err)
			}
		}
		fmt.Println("检查单次定时任务结束")
	}
	crontab.AddFunc(spec, task)
	crontab.Start()
}
