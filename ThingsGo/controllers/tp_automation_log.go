package controllers

import (
	gvalid "IOT/initialize/validate"
	"IOT/services"
	"IOT/utils"
	valid "IOT/validate"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/beego/beego/v2/core/validation"
	beego "github.com/beego/beego/v2/server/web"
	context2 "github.com/beego/beego/v2/server/web/context"
)

type TpAutomationLogController struct {
	beego.Controller
}

// 列表
func (TpAutomationLogController *TpAutomationLogController) List() {
	PaginationValidate := valid.TpAutomationLogPaginationValidate{}
	err := json.Unmarshal(TpAutomationLogController.Ctx.Input.RequestBody, &PaginationValidate)
	if err != nil {
		fmt.Println("参数解析失败", err.Error())
	}
	v := validation.Validation{}
	status, _ := v.Valid(PaginationValidate)
	if !status {
		for _, err := range v.Errors {
			// 获取字段别称
			alias := gvalid.GetAlias(PaginationValidate, err.Field)
			message := strings.Replace(err.Message, err.Field, alias, 1)
			utils.SuccessWithMessage(1000, message, (*context2.Context)(TpAutomationLogController.Ctx))
			break
		}
		return
	}
	var TpAutomationLogService services.TpAutomationLogService
	d, t, err := TpAutomationLogService.GetTpAutomationLogList(PaginationValidate)
	if err != nil {
		utils.SuccessWithMessage(1000, err.Error(), (*context2.Context)(TpAutomationLogController.Ctx))
		return
	}
	dd := valid.RspTpAutomationLogPaginationValidate{
		CurrentPage: PaginationValidate.CurrentPage,
		Data:        d,
		Total:       t,
		PerPage:     PaginationValidate.PerPage,
	}
	utils.SuccessWithDetailed(200, "success", dd, map[string]string{}, (*context2.Context)(TpAutomationLogController.Ctx))
}
