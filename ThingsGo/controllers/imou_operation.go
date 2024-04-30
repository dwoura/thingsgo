package controllers

import (
	"IOT/services"
	response "IOT/utils"
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	context2 "github.com/beego/beego/v2/server/web/context"
)

type GetSnapUrlData struct {
	BucketName string `json:"bucketName"`
	DeviceId   string `json:"deviceId"`
	StartTime  string `json:"startTime"`
	EndTime    string `json:"endTime"`
}

type DoSnapData struct {
	DeviceId string `json:"deviceId"`
}

type MonitorController struct {
	beego.Controller
}

func (this *MonitorController) GetAccessToken() {
	var monitorService services.MonitorService
	monitorService.GetAccessToken()
}

// 测试：执行抓拍
// 例如：8L10C5FPAN06A2B
func (this *MonitorController) DoSnap() {
	var monitorService services.MonitorService
	// parse参数
	var doSnapData DoSnapData
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &doSnapData)
	if err != nil {
		fmt.Println("参数解析失败", err.Error())
		response.SuccessWithDetailed(1000, "success", "参数解析失败", map[string]string{}, (*context2.Context)(this.Ctx))
		return
	}
	err = monitorService.DoSnap(doSnapData.DeviceId)
	if err != nil {
		fmt.Println("抓拍失败", err.Error())
		response.SuccessWithDetailed(1000, "success", "抓拍失败", map[string]string{}, (*context2.Context)(this.Ctx))
		return
	}
	response.SuccessWithDetailed(200, "success", "抓拍成功", map[string]string{}, (*context2.Context)(this.Ctx))
}

// 获取对象url
func (this *MonitorController) GetSnapUrl() {
	var monitorService services.MonitorService
	// parse参数
	var getSnapUrlData GetSnapUrlData
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &getSnapUrlData)
	if err != nil {
		fmt.Println("参数解析失败", err.Error())
	}

	urlList, err := monitorService.GetSnapUrl(getSnapUrlData.BucketName, getSnapUrlData.DeviceId, getSnapUrlData.StartTime, getSnapUrlData.EndTime)
	if err != nil {
		fmt.Println("获取对象url失败", err.Error())
		response.SuccessWithDetailed(1000, "success", "获取失败", map[string]string{}, (*context2.Context)(this.Ctx))
		return
	}
	d := map[string]interface{}{
		"urlList": urlList,
	}
	response.SuccessWithDetailed(200, "success", d, map[string]string{}, (*context2.Context)(this.Ctx))
	this.Data["json"] = urlList
}

// 获取已有设备
func (this *MonitorController) GetAllDeviceId() {
	var monitorService services.MonitorService
	var deviceIdList []string
	list, err := monitorService.GetAllImouKitTokenList()
	if err != nil {
		fmt.Println("获取设备失败", err.Error())
		response.SuccessWithDetailed(1000, "success", "获取失败", map[string]string{}, (*context2.Context)(this.Ctx))
		return
	}
	for _, v := range list {
		deviceIdList = append(deviceIdList, v.DeviceId)
	}
	d := map[string]interface{}{
		"deviceIdList": deviceIdList,
	}
	response.SuccessWithDetailed(200, "success", d, map[string]string{}, (*context2.Context)(this.Ctx))
	this.Data["json"] = deviceIdList
}
