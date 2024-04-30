package controllers

import (
	gvalid "IOT/initialize/validate"
	"IOT/models"
	"IOT/services"
	"IOT/utils"
	response "IOT/utils"
	valid "IOT/validate"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/beego/beego/v2/core/config/yaml"
	"github.com/beego/beego/v2/core/validation"
	beego "github.com/beego/beego/v2/server/web"
	context2 "github.com/beego/beego/v2/server/web/context"
	"github.com/bitly/go-simplejson"
)

type DashBoardController struct {
	beego.Controller
}

type PaginateDashBoard struct {
	CurrentPage int                `json:"current_page"`
	Data        []models.DashBoard `json:"data"`
	Total       int64              `json:"total"`
	PerPage     int                `json:"per_page"`
}

type PropertyData struct {
	ID             string          `json:"id"`
	AdditionalInfo string          `json:"additional_info"`
	CustomerID     string          `json:"customer_id"`
	Name           string          `json:"name"`
	Label          string          `json:"label"`
	SearchText     string          `json:"search_text"`
	Type           string          `json:"type"`
	ParentID       string          `json:"parent_id"`
	Tier           int64           `json:"tier"`
	BusinessID     string          `json:"business_id"`
	Two            []PropertyData2 `json:"two"`
}

type PropertyData2 struct {
	ID             string          `json:"id"`
	AdditionalInfo string          `json:"additional_info"`
	CustomerID     string          `json:"customer_id"`
	Name           string          `json:"name"`
	Label          string          `json:"label"`
	SearchText     string          `json:"search_text"`
	Type           string          `json:"type"`
	ParentID       string          `json:"parent_id"`
	Tier           int64           `json:"tier"`
	BusinessID     string          `json:"business_id"`
	There          []PropertyData3 `json:"there"`
}

type InserttimeData struct {
	StartTime    string `json:"start_time"`
	EndTime      string `json:"end_time"`
	Theme        int64  `json:"theme"`
	IntervalTime int64  `json:"interval_time"`
	BgTheme      int64  `json:"bg_theme"`
}

type PropertyData3 struct {
	ID             string `json:"id"`
	AdditionalInfo string `json:"additional_info"`
	CustomerID     string `json:"customer_id"`
	Name           string `json:"name"`
	Label          string `json:"label"`
	SearchText     string `json:"search_text"`
	Type           string `json:"type"`
	ParentID       string `json:"parent_id"`
	Tier           int64  `json:"tier"`
	BusinessID     string `json:"business_id"`
}

type GettimeData struct {
	ID                string         `json:"id"`
	Configuration     string         `json:"configuration"`
	AssignedCustomers string         `json:"assigned_customers"`
	SearchText        string         `json:"search_text"`
	Title             string         `json:"title"`
	BusinessID        string         `json:"business_id"`
	Config            InserttimeData `json:"config"`
}

type RealtimeData struct {
	EndTime   string `json:"end_time"`
	StartTime string `json:"start_time"`
}

type DashBoardData struct {
	ID        string                           `json:"id"`
	SliceId   int64                            `json:"slice_id"`
	X         int64                            `json:"x"`
	Y         int64                            `json:"y"`
	W         int64                            `json:"w"`
	H         int64                            `json:"h"`
	Width     int64                            `json:"width"`
	Height    int64                            `json:"height"`
	I         string                           `json:"i"`
	ChartType string                           `json:"chart_type"`
	Title     string                           `json:"title"`
	Fields    []map[string]DashBoardFieldsData `json:"fields"`
}

type DashBoardConfig struct {
	SliceId   int64  `json:"slice_id"`
	X         int64  `json:"x"`
	Y         int64  `json:"y"`
	W         int64  `json:"w"`
	H         int64  `json:"h"`
	Width     int64  `json:"width"`
	Height    int64  `json:"height"`
	I         string `json:"i"`
	ChartType string `json:"chart_type"`
	Title     string `json:"title"`
}

type DashBoardFieldsData struct {
	Name   string `json:"name"`
	Type   int64  `json:"type"`
	Symbol string `json:"symbol"`
}

type WidgetIcon struct {
	Key       string `json:"key"`
	Name      string `json:"name"`
	Thumbnail string `json:"thumbnail"`
}

// 视图列表
func (this *DashBoardController) Index() {
	paginateDashBoardValidate := valid.PaginateDashBoard{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &paginateDashBoardValidate)
	if err != nil {
		fmt.Println("参数解析失败", err.Error())
	}
	v := validation.Validation{}
	status, _ := v.Valid(paginateDashBoardValidate)
	if !status {
		for _, err := range v.Errors {
			// 获取字段别称
			alias := gvalid.GetAlias(paginateDashBoardValidate, err.Field)
			message := strings.Replace(err.Message, err.Field, alias, 1)
			response.SuccessWithMessage(1000, message, (*context2.Context)(this.Ctx))
			break
		}
		return
	}
	var DashBoardService services.DashBoardService
	offset := (paginateDashBoardValidate.Page - 1) * paginateDashBoardValidate.Limit
	u, c := DashBoardService.Paginate(paginateDashBoardValidate.Title, offset, paginateDashBoardValidate.Limit)
	d := PaginateDashBoard{
		CurrentPage: paginateDashBoardValidate.Page,
		Data:        u,
		Total:       c,
		PerPage:     paginateDashBoardValidate.Limit,
	}
	response.SuccessWithDetailed(200, "success", d, map[string]string{}, (*context2.Context)(this.Ctx))
	return
}

// 添加视图
func (this *DashBoardController) Add() {
	addDashBoardValidate := valid.AddDashBoard{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &addDashBoardValidate)
	if err != nil {
		fmt.Println("参数解析失败", err.Error())
	}
	v := validation.Validation{}
	status, _ := v.Valid(addDashBoardValidate)
	if !status {
		for _, err := range v.Errors {
			alias := gvalid.GetAlias(addDashBoardValidate, err.Field)
			message := strings.Replace(err.Message, err.Field, alias, 1)
			response.SuccessWithMessage(1000, message, (*context2.Context)(this.Ctx))
			break
		}
		return
	}
	var DashBoardService services.DashBoardService
	f, _ := DashBoardService.Add(addDashBoardValidate.BusinessId, addDashBoardValidate.Title)
	if f {
		response.SuccessWithMessage(200, "新增成功", (*context2.Context)(this.Ctx))
		return
	}
	response.SuccessWithMessage(400, "新增失败", (*context2.Context)(this.Ctx))
	return
}

// 编辑图表
func (this *DashBoardController) Edit() {
	editDashBoardValidate := valid.EditDashBoard{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &editDashBoardValidate)
	if err != nil {
		fmt.Println("参数解析失败", err.Error())
	}
	v := validation.Validation{}
	status, _ := v.Valid(editDashBoardValidate)
	if !status {
		for _, err := range v.Errors {
			alias := gvalid.GetAlias(editDashBoardValidate, err.Field)
			message := strings.Replace(err.Message, err.Field, alias, 1)
			response.SuccessWithMessage(1000, message, (*context2.Context)(this.Ctx))
			break
		}
		return
	}
	var DashBoardService services.DashBoardService
	f := DashBoardService.Edit(editDashBoardValidate.ID, editDashBoardValidate.BusinessID, editDashBoardValidate.Title)
	if f {
		response.SuccessWithMessage(200, "编辑成功", (*context2.Context)(this.Ctx))
		return
	}
	response.SuccessWithMessage(400, "编辑失败", (*context2.Context)(this.Ctx))
	return
}

// 删除图表
func (this *DashBoardController) Delete() {
	deleteDashBoardValidate := valid.DeleteDashBoard{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &deleteDashBoardValidate)
	if err != nil {
		fmt.Println("参数解析失败", err.Error())
	}
	v := validation.Validation{}
	status, _ := v.Valid(deleteDashBoardValidate)
	if !status {
		for _, err := range v.Errors {
			alias := gvalid.GetAlias(deleteDashBoardValidate, err.Field)
			message := strings.Replace(err.Message, err.Field, alias, 1)
			response.SuccessWithMessage(1000, message, (*context2.Context)(this.Ctx))
			break
		}
		return
	}
	var DashBoardService services.DashBoardService
	f := DashBoardService.Delete(deleteDashBoardValidate.ID)
	if f {
		response.SuccessWithMessage(200, "删除成功", (*context2.Context)(this.Ctx))
		return
	}
	response.SuccessWithMessage(400, "删除失败", (*context2.Context)(this.Ctx))
	return
}

// 业务数据
func (this *DashBoardController) Business() {
	var BusinessService services.BusinessService
	d, _ := BusinessService.All()
	response.SuccessWithDetailed(200, "success", d, map[string]string{}, (*context2.Context)(this.Ctx))
	return
}

func (this *DashBoardController) List() {
	listDashBoardValidate := valid.ListDashBoard{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &listDashBoardValidate)
	if err != nil {
		fmt.Println("参数解析失败", err.Error())
	}
	v := validation.Validation{}
	status, _ := v.Valid(listDashBoardValidate)
	if !status {
		for _, err := range v.Errors {
			alias := gvalid.GetAlias(listDashBoardValidate, err.Field)
			message := strings.Replace(err.Message, err.Field, alias, 1)
			response.SuccessWithMessage(1000, message, (*context2.Context)(this.Ctx))
			break
		}
		return
	}
	var WidgetService services.WidgetService
	d, _ := WidgetService.GetWidgetDashboardId(listDashBoardValidate.DashBoardID)
	response.SuccessWithDetailed(200, "success", d, map[string]string{}, (*context2.Context)(this.Ctx))
	return
}

// 设备数据
func (this *DashBoardController) Property() {
	propertyAssetValidate := valid.PropertyAsset{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &propertyAssetValidate)
	if err != nil {
		fmt.Println("参数解析失败", err.Error())
	}
	v := validation.Validation{}
	status, _ := v.Valid(propertyAssetValidate)
	if !status {
		for _, err := range v.Errors {
			alias := gvalid.GetAlias(propertyAssetValidate, err.Field)
			message := strings.Replace(err.Message, err.Field, alias, 1)
			response.SuccessWithMessage(1000, message, (*context2.Context)(this.Ctx))
			break
		}
		return
	}
	var AssetService services.AssetService
	var propertyData []PropertyData
	a1, c1 := AssetService.GetAssetData(propertyAssetValidate.BusinessID)
	if c1 > 0 {
		for _, av1 := range a1 {
			// 第二层
			a2, c2 := AssetService.GetAssetsByParentID(av1.ID)
			var propertyData2 []PropertyData2
			if c2 > 0 {
				for _, av2 := range a2 {
					// 第三层
					a3, c3 := AssetService.GetAssetsByParentID(av2.ID)
					var propertyData3 []PropertyData3
					if c3 > 0 {
						for _, av3 := range a3 {
							ai3 := PropertyData3{
								ID:             av3.ID,
								AdditionalInfo: av3.AdditionalInfo,
								CustomerID:     av3.CustomerID,
								Name:           av3.Name,
								Label:          av3.Label,
								SearchText:     av3.SearchText,
								Type:           av3.Type,
								ParentID:       av3.ParentID,
								Tier:           av3.Tier,
								BusinessID:     av3.BusinessID,
							}
							propertyData3 = append(propertyData3, ai3)
						}
					}
					if len(propertyData3) == 0 {
						propertyData3 = []PropertyData3{}
					}
					ai2 := PropertyData2{
						ID:             av2.ID,
						AdditionalInfo: av2.AdditionalInfo,
						CustomerID:     av2.CustomerID,
						Name:           av2.Name,
						Label:          av2.Label,
						SearchText:     av2.SearchText,
						Type:           av2.Type,
						ParentID:       av2.ParentID,
						Tier:           av2.Tier,
						BusinessID:     av2.BusinessID,
						There:          propertyData3,
					}
					propertyData2 = append(propertyData2, ai2)
				}
			}
			if len(propertyData2) == 0 {
				propertyData2 = []PropertyData2{}
			}
			ai1 := PropertyData{
				ID:             av1.ID,
				AdditionalInfo: av1.AdditionalInfo,
				CustomerID:     av1.CustomerID,
				Name:           av1.Name,
				Label:          av1.Label,
				SearchText:     av1.SearchText,
				Type:           av1.Type,
				ParentID:       av1.ParentID,
				Tier:           av1.Tier,
				BusinessID:     av1.BusinessID,
				Two:            propertyData2,
			}
			propertyData = append(propertyData, ai1)
		}
	}
	if len(propertyData) == 0 {
		propertyData = []PropertyData{}
	}
	response.SuccessWithDetailed(200, "success", propertyData, map[string]string{}, (*context2.Context)(this.Ctx))
	return
}

func (this *DashBoardController) Device() {
	deviceDashBoardValidate := valid.DeviceDashBoard{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &deviceDashBoardValidate)
	if err != nil {
		fmt.Println("参数解析失败", err.Error())
	}
	v := validation.Validation{}
	status, _ := v.Valid(deviceDashBoardValidate)
	if !status {
		for _, err := range v.Errors {
			alias := gvalid.GetAlias(deviceDashBoardValidate, err.Field)
			message := strings.Replace(err.Message, err.Field, alias, 1)
			response.SuccessWithMessage(1000, message, (*context2.Context)(this.Ctx))
			break
		}
		return
	}
	var DeviceService services.DeviceService
	d, _ := DeviceService.GetDevicesByAssetID(deviceDashBoardValidate.AssetID)
	response.SuccessWithDetailed(200, "success", d, map[string]string{}, (*context2.Context)(this.Ctx))
	return
}

// 输入时间
func (this *DashBoardController) Inserttime() {
	inserttimeDashBoardValidate := valid.InserttimeDashBoard{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &inserttimeDashBoardValidate)
	if err != nil {
		fmt.Println("参数解析失败", err.Error())
	}
	v := validation.Validation{}
	status, _ := v.Valid(inserttimeDashBoardValidate)
	if !status {
		for _, err := range v.Errors {
			alias := gvalid.GetAlias(inserttimeDashBoardValidate, err.Field)
			message := strings.Replace(err.Message, err.Field, alias, 1)
			response.SuccessWithMessage(1000, message, (*context2.Context)(this.Ctx))
			break
		}
		return
	}
	di := InserttimeData{
		StartTime:    inserttimeDashBoardValidate.StartTime,
		EndTime:      inserttimeDashBoardValidate.EndTime,
		Theme:        inserttimeDashBoardValidate.Theme,
		IntervalTime: inserttimeDashBoardValidate.IntervalTime,
		BgTheme:      inserttimeDashBoardValidate.BgTheme,
	}
	dcJson, _ := json.Marshal(di)
	config := string(dcJson)
	var DashBoardService services.DashBoardService
	_, ac := DashBoardService.GetDashBoardById(inserttimeDashBoardValidate.ID)
	if ac > 0 {
		//更新
		ri, rf := DashBoardService.ConfigurationEdit(inserttimeDashBoardValidate.ID, config)
		if rf {
			response.SuccessWithDetailed(200, "success", ri, map[string]string{}, (*context2.Context)(this.Ctx))
			return
		} else {
			response.SuccessWithMessage(400, "error", (*context2.Context)(this.Ctx))
			return
		}
	} else {
		// 插入
		ri, rf := DashBoardService.ConfigurationAdd(config)
		if rf {
			response.SuccessWithDetailed(200, "success", ri, map[string]string{}, (*context2.Context)(this.Ctx))
			return
		} else {
			response.SuccessWithMessage(400, "error", (*context2.Context)(this.Ctx))
			return
		}
	}
}

// 获取时间
func (this *DashBoardController) Gettime() {
	gettimeDashBoardValidate := valid.GettimeDashBoard{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &gettimeDashBoardValidate)
	if err != nil {
		fmt.Println("参数解析失败", err.Error())
	}
	v := validation.Validation{}
	status, _ := v.Valid(gettimeDashBoardValidate)
	if !status {
		for _, err := range v.Errors {
			alias := gvalid.GetAlias(gettimeDashBoardValidate, err.Field)
			message := strings.Replace(err.Message, err.Field, alias, 1)
			response.SuccessWithMessage(1000, message, (*context2.Context)(this.Ctx))
			break
		}
		return
	}
	var DashBoardService services.DashBoardService
	di, dc := DashBoardService.GetDashBoardById(gettimeDashBoardValidate.ID)
	if dc > 0 {
		var config InserttimeData
		err := json.Unmarshal([]byte(di.Configuration), &config)
		if err != nil {
			fmt.Println(err)
		}
		timeTemplate := "2006-01-02T15:04"
		et := time.Now().Unix()
		st := et - 300
		start_time := time.Unix(st, 0).Format(timeTemplate)
		end_time := time.Unix(et, 0).Format(timeTemplate)
		config.StartTime = start_time
		config.EndTime = end_time
		config_str, _ := json.Marshal(config)
		res := GettimeData{
			ID:                di.ID,
			Configuration:     string(config_str),
			AssignedCustomers: di.AssignedCustomers,
			SearchText:        di.SearchText,
			Title:             di.Title,
			BusinessID:        di.BusinessID,
			Config:            config,
		}
		response.SuccessWithDetailed(200, "success", res, map[string]string{}, (*context2.Context)(this.Ctx))
		return
	}
	response.SuccessWithMessage(400, "error", (*context2.Context)(this.Ctx))
	return
}

// 可视化图标
func (this *DashBoardController) Dashboard() {
	dashBoardDashBoardValidate := valid.DashBoardDashBoard{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &dashBoardDashBoardValidate)
	if err != nil {
		fmt.Println("参数解析失败", err.Error())
	}
	v := validation.Validation{}
	status, _ := v.Valid(dashBoardDashBoardValidate)
	if !status {
		for _, err := range v.Errors {
			alias := gvalid.GetAlias(dashBoardDashBoardValidate, err.Field)
			message := strings.Replace(err.Message, err.Field, alias, 1)
			response.SuccessWithMessage(1000, message, (*context2.Context)(this.Ctx))
			break
		}
		return
	}
	var WidgetService services.WidgetService
	var AssetService services.AssetService
	var config services.DashboardConfig
	var fieldDashBoardData []DashBoardData
	wl, wc := WidgetService.GetWidgetDashboardIdAndAssetId(dashBoardDashBoardValidate.DashboardID, dashBoardDashBoardValidate.AssetId)
	if wc > 0 {
		for _, wv := range wl {
			arr := strings.Split(wv.WidgetIdentifier, ":")
			var fields []map[string]DashBoardFieldsData
			if len(arr) > 0 {
				fs := AssetService.Field(arr[0], arr[1])
				if len(fs) > 0 {
					for _, fv := range fs {
						fi := map[string]DashBoardFieldsData{
							fv.Key: {
								Name:   fv.Name,
								Type:   fv.Type,
								Symbol: fv.Symbol,
							},
						}
						fields = append(fields, fi)
					}
				}

			}
			if len(fields) == 0 {
				fields = []map[string]DashBoardFieldsData{}
			}

			err := json.Unmarshal([]byte(wv.Config), &config)
			if err != nil {
				fmt.Println(err)
			}
			var DeviceService services.DeviceService
			theDevice, _ := DeviceService.GetDeviceByID(wv.DeviceID)
			// 赋值
			d := DashBoardData{
				ID:        wv.ID,
				SliceId:   config.SliceId,
				X:         config.X,
				Y:         config.Y,
				W:         config.W,
				H:         config.H,
				Width:     config.Width,
				Height:    config.Height,
				I:         config.I,
				ChartType: config.ChartType,
				Title:     theDevice.Name,
				Fields:    fields,
			}
			fieldDashBoardData = append(fieldDashBoardData, d)
		}
	}
	if len(fieldDashBoardData) == 0 {
		fieldDashBoardData = []DashBoardData{}
	}
	response.SuccessWithDetailed(200, "success", fieldDashBoardData, map[string]string{}, (*context2.Context)(this.Ctx))
	return
}

// 转换时间
func (this *DashBoardController) Realtime() {
	realtimeDashBoardValidate := valid.RealtimeDashBoard{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &realtimeDashBoardValidate)
	if err != nil {
		fmt.Println("参数解析失败", err.Error())
	}
	v := validation.Validation{}
	status, _ := v.Valid(realtimeDashBoardValidate)
	if !status {
		for _, err := range v.Errors {
			alias := gvalid.GetAlias(realtimeDashBoardValidate, err.Field)
			message := strings.Replace(err.Message, err.Field, alias, 1)
			response.SuccessWithMessage(1000, message, (*context2.Context)(this.Ctx))
			break
		}
		return
	}
	timeTemplate := "2006-01-02 15:04:05"
	tn := time.Now().Unix()
	var ts string
	if realtimeDashBoardValidate.Type == 1 {
		ts = time.Unix(tn-900, 0).Format(timeTemplate)
	} else if realtimeDashBoardValidate.Type == 2 {
		ts = time.Unix(tn-1800, 0).Format(timeTemplate)
	} else if realtimeDashBoardValidate.Type == 3 {
		ts = time.Unix(tn-3600, 0).Format(timeTemplate)
	} else if realtimeDashBoardValidate.Type == 4 {
		ts = time.Unix(tn-3600*3, 0).Format(timeTemplate)
	} else if realtimeDashBoardValidate.Type == 5 {
		ts = time.Unix(tn-3600*6, 0).Format(timeTemplate)
	} else if realtimeDashBoardValidate.Type == 6 {
		ts = time.Unix(tn-3600*12, 0).Format(timeTemplate)
	} else if realtimeDashBoardValidate.Type == 7 {
		ts = time.Unix(tn-3600*24, 0).Format(timeTemplate)
	}
	te := time.Unix(tn, 0).Format(timeTemplate)
	d := RealtimeData{
		EndTime:   te,
		StartTime: ts,
	}
	response.SuccessWithDetailed(200, "success", d, map[string]string{}, (*context2.Context)(this.Ctx))
	return
}

// 跟新可视化组件
func (this *DashBoardController) Updatedashboard() {
	updateDashBoardValidate := valid.UpdateDashBoard{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &updateDashBoardValidate)
	if err != nil {
		fmt.Println("参数解析失败", err.Error())
	}
	v := validation.Validation{}
	status, _ := v.Valid(updateDashBoardValidate)
	if !status {
		for _, err := range v.Errors {
			alias := gvalid.GetAlias(updateDashBoardValidate, err.Field)
			message := strings.Replace(err.Message, err.Field, alias, 1)
			response.SuccessWithMessage(1000, message, (*context2.Context)(this.Ctx))
			break
		}
		return
	}
	var dbc DashBoardConfig
	var WidgetService services.WidgetService
	w, c := WidgetService.GetWidgetById(updateDashBoardValidate.WidgetID)
	if c > 0 {
		res, err := simplejson.NewJson([]byte(updateDashBoardValidate.Config))
		if err != nil {
			fmt.Println("解析出错", err)
		}
		qx := res.Get("x").MustInt64()
		qy := res.Get("y").MustInt64()
		qw := res.Get("w").MustInt64()
		qh := res.Get("h").MustInt64()
		qwidth := res.Get("width").MustInt64()
		qheight := res.Get("height").MustInt64()
		res2, err2 := simplejson.NewJson([]byte(w.Config))
		if err2 != nil {
			fmt.Println("解析出错", err2)
		}
		slice_id := res2.Get("slice_id").MustInt64()
		i := res2.Get("i").MustString()
		chart_type := res2.Get("chart_type").MustString()
		title := res2.Get("title").MustString()
		dbc = DashBoardConfig{
			SliceId:   slice_id,
			X:         qx,
			Y:         qy,
			W:         qw,
			H:         qh,
			Width:     qwidth,
			Height:    qheight,
			I:         i,
			ChartType: chart_type,
			Title:     title,
		}
		dbc_string, _ := json.Marshal(&dbc)
		f := WidgetService.UpdateConfigByWidgetId(updateDashBoardValidate.WidgetID, string(dbc_string))
		if f {
			response.SuccessWithMessage(200, "编辑成功", (*context2.Context)(this.Ctx))
			return
		}
	}
	response.SuccessWithMessage(400, "编辑失败", (*context2.Context)(this.Ctx))
	return
}

func (this *DashBoardController) Component() {
	componentDashBoardValidate := valid.ComponentDashBoard{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &componentDashBoardValidate)
	if err != nil {
		fmt.Println("参数解析失败", err.Error())
	}
	v := validation.Validation{}
	status, _ := v.Valid(componentDashBoardValidate)
	if !status {
		for _, err := range v.Errors {
			alias := gvalid.GetAlias(componentDashBoardValidate, err.Field)
			message := strings.Replace(err.Message, err.Field, alias, 1)
			response.SuccessWithMessage(1000, message, (*context2.Context)(this.Ctx))
			break
		}
		return
	}
	var DeviceService services.DeviceService
	var AssetService services.AssetService
	var wi []WidgetIcon
	d, c := DeviceService.GetAllDeviceByID(componentDashBoardValidate.DeviceID)
	if c > 0 {
		for _, dv := range d {
			widgets := AssetService.Widget(dv.Type)
			if len(widgets) > 0 {
				for _, wv := range widgets {
					i := WidgetIcon{
						Key:       dv.Type + ":" + wv.Key,
						Name:      wv.Name,
						Thumbnail: wv.Thumbnail,
					}
					wi = append(wi, i)
				}
			}
		}
	}
	if len(wi) == 0 {
		wi = []WidgetIcon{}
	}
	response.SuccessWithDetailed(200, "success", wi, map[string]string{}, (*context2.Context)(this.Ctx))
	return
}

func (thisController *DashBoardController) PluginList() {
	var DashBoardService services.DashBoardService
	plugList := DashBoardService.GetPlugList()
	response.SuccessWithDetailed(200, "success", plugList, map[string]string{}, (*context2.Context)(thisController.Ctx))
}

// 业务预览的组件列表
func (dashBoardController *DashBoardController) BidComponent() {
	BidComponentValidate := valid.BidComponentValidate{}
	err := json.Unmarshal(dashBoardController.Ctx.Input.RequestBody, &BidComponentValidate)
	if err != nil {
		fmt.Println("参数解析失败", err.Error())
	}
	v := validation.Validation{}
	status, _ := v.Valid(BidComponentValidate)
	if !status {
		for _, err := range v.Errors {
			alias := gvalid.GetAlias(BidComponentValidate, err.Field)
			message := strings.Replace(err.Message, err.Field, alias, 1)
			response.SuccessWithMessage(1000, message, (*context2.Context)(dashBoardController.Ctx))
			break
		}
		return
	}
	var wi []WidgetIcon
	f := utils.FileExist("./extensions/business/config.yaml")
	if f {
		conf, err := yaml.ReadYmlReader("./extensions/business/config.yaml")
		if err != nil {
			fmt.Println(err)
		}
		for _, v := range conf {
			str, _ := v.(map[string]interface{})
			widgets, _ := str["widgets"].(map[string]interface{})
			if len(widgets) > 0 {
				for wk, wv := range widgets {
					item, _ := wv.(map[string]interface{})
					WidgetIcon := WidgetIcon{
						Key:       "business:" + wk,
						Name:      item["name"].(string),
						Thumbnail: "",
					}
					wi = append(wi, WidgetIcon)
				}
			}
		}
	}
	if len(wi) == 0 {
		wi = []WidgetIcon{}
	}
	response.SuccessWithDetailed(200, "success", wi, map[string]string{}, (*context2.Context)(dashBoardController.Ctx))
}
