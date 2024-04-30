package controllers

import (
	gvalid "IOT/initialize/validate"
	"IOT/services"
	response "IOT/utils"
	valid "IOT/validate"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/beego/beego/v2/core/validation"
	beego "github.com/beego/beego/v2/server/web"
	context2 "github.com/beego/beego/v2/server/web/context"
)

type ConditionslogController struct {
	beego.Controller
}

type PaginateConditionslog struct {
	CurrentPage int                      `json:"current_page"`
	Data        []map[string]interface{} `json:"data"`
	Total       int64                    `json:"total"`
	PerPage     int                      `json:"per_page"`
}

// 获取控制日志
func (conditionslogController *ConditionslogController) Index() {
	conditionsLogListValidate := valid.ConditionsLogListValidate{}
	err := json.Unmarshal(conditionslogController.Ctx.Input.RequestBody, &conditionsLogListValidate)
	if err != nil {
		fmt.Println("参数解析失败", err.Error())
	}
	v := validation.Validation{}
	status, _ := v.Valid(conditionsLogListValidate)
	if !status {
		for _, err := range v.Errors {
			// 获取字段别称
			alias := gvalid.GetAlias(conditionsLogListValidate, err.Field)
			message := strings.Replace(err.Message, err.Field, alias, 1)
			response.SuccessWithMessage(1000, message, (*context2.Context)(conditionslogController.Ctx))
			break
		}
		return
	}
	var ConditionsLogService services.ConditionsLogService
	w, count := ConditionsLogService.Paginate(conditionsLogListValidate)
	d := PaginateConditionslog{
		CurrentPage: conditionsLogListValidate.Current,
		PerPage:     conditionsLogListValidate.Size,
		Data:        w,
		Total:       count,
	}
	response.SuccessWithDetailed(200, "success", d, map[string]string{}, (*context2.Context)(conditionslogController.Ctx))
}
