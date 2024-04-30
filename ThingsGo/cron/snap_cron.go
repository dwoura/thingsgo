package cron

import (
	"IOT/services"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/bitly/go-simplejson"
	"github.com/google/uuid"
	"github.com/minio/minio-go"
	"github.com/robfig/cron/v3"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Token struct {
	Value      string `json:"value"`
	ExpireTime int64  `json:"expireTime"`
}

var AppId = "lc15c414d4baf0483a"
var AppSecret = "9cfedd2b3f1144d38a72a98b3f941f"

// var BaseApi string = "https://openapi.lechange.cn/openapi"
var BaseApi = "https://openapi.lechange.cn:443/openapi"
var GetAccessToken = "/accessToken"

// var GetKitToken = "/getKitToken"
var GetSnapEnhanced = "/setDeviceSnapEnhanced"

var AccessToken string

var AccessTokenExpire int64

//var SuccessCode string

// 存储设备的ID + # + channel 和 kitToken
// map初始化，防止触发panic
//var KitTokenMap = make(map[string]Token)

var MinioClient *minio.Client // 持有minio客户端

func getSign(time int64, nonce string, appSecret string) string {
	// 使用当前时间、nonce、appSecret进行md5加密获取
	data := []byte("time:" + strconv.FormatInt(time, 10) + ",nonce:" + nonce + ",appSecret:" + appSecret)
	fmt.Println("raw data: ", "time:"+strconv.FormatInt(time, 10)+",nonce:"+nonce+",appSecret:"+appSecret)
	hash := md5.New()
	// 字节流进行md5加密
	hash.Write(data)
	// md5 Hash转为string
	md5String := hex.EncodeToString(hash.Sum(nil))
	fmt.Println("md5 hash: ", md5String)
	return md5String

}

func getRequestBaseData() *simplejson.Json {
	/**
	{
	    "system":{
	        "ver":"1.0",
	        "appId":"lcdxxxxxxxxx",
	        "sign":"b7e5bbcc6cc07941725d9ad318883d8e",
	        "time":1599013514,
	        "nonce":"fbf19fc6-17a1-4f73-a967-75eadbc805a2"
	    },
	    "id":"98a7a257-c4e4-4db3-a2d3-d97a3836b87c",
	    "params":{
	    }
	}
	*/
	rawJson := "{\n    \"system\":{\n        \"ver\":\"1.0\",\n        \"appId\":\"\",\n        \"sign\":\"\",\n        \"time\":0,\n        \"nonce\":\"\"\n    },\n    \"id\":\"98a7a257-c4e4-4db3-a2d3-d97a3836b87c\",\n    \"params\":{\n\n    }\n}\n"
	jsonData, err := simplejson.NewJson([]byte(rawJson))
	// 获取当前时间戳
	t := time.Now().UnixNano() / int64(time.Millisecond*1000)
	// 使用UUID生成nonce
	nonce := uuid.New().String()
	id := uuid.New().String()

	sign := getSign(t, nonce, AppSecret)
	jsonData.SetPath([]string{"system", "appId"}, AppId)
	jsonData.SetPath([]string{"system", "sign"}, sign)
	jsonData.SetPath([]string{"system", "time"}, t)
	jsonData.SetPath([]string{"system", "nonce"}, nonce)
	jsonData.SetPath([]string{"id"}, id)

	if err != nil {
		log.Fatalln("json parse fail!", err)
		return nil
	}
	return jsonData

}

func getAccessToken() string {
	// 1.拼接json参数
	jsonData := getRequestBaseData()
	marshalJson, err := json.Marshal(jsonData)
	if err != nil {
		log.Fatalf("json transform to byte fail!", err)
		return "nil"
	}
	// 2.拼接请求url，发送请求
	request, reqErr := http.NewRequest("POST", BaseApi+GetAccessToken, bytes.NewBuffer(marshalJson))
	// 设置请求头
	request.Header.Set("Content-Type", "application/json")
	if reqErr != nil {
		log.Fatalln("create request fail!", reqErr)
		return "nil"
	}
	// 3.获取response，accessToken持久化
	client := http.Client{}
	response, repErr := client.Do(request)
	if repErr != nil {
		log.Fatalln("get response fail!", err)
		return "nil"
	}
	// 解析相应数据获取token
	body, transformErr := io.ReadAll(response.Body)
	if transformErr != nil {
		log.Fatalln("transform response to bytes fail!", transformErr)
		return "nil"
	}
	repJson, resolveErr := simplejson.NewJson(body)
	if resolveErr != nil {
		log.Fatalln("resolve response to json fail!", resolveErr)
		return "nil"
	}
	// 响应中的accessToken
	AccessToken = repJson.GetPath("result", "data", "accessToken").MustString()
	// 响应中的过期时间
	AccessTokenExpire = time.Now().Unix() + repJson.GetPath("result", "data", "expireTime").MustInt64()
	//fmt.Println(AccessTokenExpire)
	if response != nil && response.ContentLength > 0 {
		// 将获取到的accessToken持久化到数据库
		var imouAccessToken services.ImouAccessTokenService
		// 获取最新的token与当前token对比
		last, _ := imouAccessToken.GetLastToken()
		// 初始化查询为空时，插入token
		if last.AccessToken == "" || last.Id == "" || last.ExpireTime == 0 {
			add, _ := imouAccessToken.AddAccessToken(AccessToken, AccessTokenExpire)
			if add {
				return AccessToken
			}
		}
	}
	return AccessToken
}

func getKitTokenMapKey(deviceId, channel string) string {
	return deviceId + "#" + channel
}

// deviceId：8L10C5FPAN06A2B
//func getKitToken() map[string]Token {
//	//token := getAccessToken()
//	// 查询所有的设备并遍历
//	var imouKitTokenService services.ImouKitTokenService
//	kitTokenList := imouKitTokenService.GetImouKitTokenList()
//
//	for _, device := range kitTokenList {
//		// 查询access token
//		var imouAccessTokenService = services.ImouAccessTokenService{}
//		accessToken, _ := imouAccessTokenService.GetLastToken()
//		currentTimeStamp := time.Now().Unix()
//		// accessToken expire
//		if accessToken.ExpireTime < currentTimeStamp {
//			accessToken.AccessToken = getAccessToken()
//		}
//		// 数据库中初始化数据，查出所有设备的token，进行遍历
//		// 1、数据库为空时，直接执行插入操作
//		// 2、数据库不为空时，检查kitToken是否过期，对某个设备过期的token重新获取并替换更新
//		jsonTest := getRequestBaseData()
//		jsonTest.SetPath([]string{"params", "token"}, accessToken.AccessToken)
//		jsonTest.SetPath([]string{"params", "deviceId"}, device.DeviceId)
//		jsonTest.SetPath([]string{"params", "channelId"}, 0)
//		jsonTest.SetPath([]string{"params", "type"}, 0)
//		// 发送request请求，获取kitToken
//		marshalJson, _ := json.Marshal(jsonTest)
//		request, reqErr := http.NewRequest("POST", BaseApi+GetKitToken, bytes.NewBuffer(marshalJson))
//		// 设置请求头
//		request.Header.Set("Content-Type", "application/json")
//		if reqErr != nil {
//			log.Fatalln("create request fail!", reqErr)
//		}
//		client := http.Client{}
//		response, repErr := client.Do(request)
//		if repErr != nil {
//			log.Fatalln("get response fail!", repErr)
//		}
//		// 解析相应数据获取token
//		body, transformErr := io.ReadAll(response.Body)
//		if transformErr != nil {
//			log.Fatalln("transform response to bytes fail!", transformErr)
//		}
//		repJson, resolveErr := simplejson.NewJson(body)
//		if resolveErr != nil {
//			log.Fatalln("resolve response to json fail!", resolveErr)
//		}
//		//	/*
//		//		{
//		//			"id":"123456",
//		//			"result":{
//		//				"code":"0",
//		//				"msg":"操作成功",
//		//				"data":{
//		//					"expireTime": 84323,//过期剩余秒数
//		//					"kitToken":"Kt_e6cf503603b848098376bc2dc1c6f38d" //轻应用授权token，新获取的kitToken有效期为2个小时;
//		//				}
//		//			}
//		//		 }
//		//	*/
//		SuccessCode = repJson.GetPath("result", "code").MustString()
//		// 检查kit token获取响应码
//		if SuccessCode == "0" {
//			kitToken := repJson.GetPath("result", "data", "kitToken").MustString()
//			expireTime := time.Now().Unix() + repJson.GetPath("result", "data", "expireTime").MustInt64()
//			// kit token持久化
//			var imouKitToken models.ImouKitToken
//			imouKitToken.Id = uuid.New().String()
//			imouKitToken.DeviceId = device.DeviceId
//			imouKitToken.KitToken = kitToken
//			update, _ := imouKitTokenService.UpdateKitToken(imouKitToken, kitToken)
//			// 数据库更新成功后，将token数据放入全局map中
//			if update {
//				KitTokenMap[getKitTokenMapKey(device.DeviceId, "0")] = Token{Value: kitToken, ExpireTime: expireTime}
//			}
//		}
//
//	}
//	// 将全部摄像头的token map返回
//	return KitTokenMap
//}

// 抓拍
func setDeviceSnapEnhanced(deviceId string) bool {
	//1.获取accessToken
	var imouAccessTokenService = services.ImouAccessTokenService{}
	accessToken, _ := imouAccessTokenService.GetLastToken()
	currentTimeStamp := time.Now().Unix()
	// accessToken expire
	if accessToken.ExpireTime < currentTimeStamp {
		// 重新获取access token
		accessToken.AccessToken = getAccessToken()
	}
	// 2.获取请求参数json数据
	baseData := getRequestBaseData()
	baseData.SetPath([]string{"params", "deviceId"}, deviceId)
	baseData.SetPath([]string{"params", "channelId"}, 0)
	baseData.SetPath([]string{"params", "token"}, accessToken.AccessToken)
	marshalJson, err := json.Marshal(baseData)
	if err != nil {
		log.Fatalln("json Marshal fail!", err)
	}
	request, reqErr := http.NewRequest("POST", BaseApi+GetSnapEnhanced, bytes.NewBuffer(marshalJson))
	// 设置请求头
	request.Header.Set("Content-Type", "application/json")
	if reqErr != nil {
		log.Fatalln("create request fail!", reqErr)
	}
	client := http.Client{}
	response, repErr := client.Do(request)
	if repErr != nil {
		log.Fatalln("get response fail!", repErr)
	}
	body, transformErr := io.ReadAll(response.Body)
	if transformErr != nil {
		log.Fatalln("transform response to bytes fail!", transformErr)
	}
	repJson, resolveErr := simplejson.NewJson(body)
	if resolveErr != nil {
		log.Fatalln("resolve response to json fail!", resolveErr)
	}
	url := repJson.GetPath("result", "data", "url").MustString()
	// 存储图片
	flag := uploadSnapPictureToMinio(url, deviceId)
	if !flag {
		return false
	}
	return true
}

// Minio client init
func InitMinioClient() *minio.Client {
	// 基本的配置信息
	endpoint := "47.105.41.204:9001"
	accessKeyID := "minioadmin"
	secretAccessKey := "grape12138"

	// 初始化一个minio客户端对象
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, false)
	if err != nil {
		log.Fatalf("init MinioClient faile: %s", err.Error())
	}
	return minioClient
}

func getObjectName(snapUrl string) string {
	return strings.Split(snapUrl, "/")[5]
}

func getCurrentTime() string {
	currentTime := time.Now()
	dateFormat := currentTime.Format("2006-01-02")
	timeFormat := currentTime.Format("15-04-05")
	return dateFormat + "-" + timeFormat
}

// 上传抓拍图片到minio
func uploadSnapPictureToMinio(imgUrl string, deviceId string) bool {
	// 获取minio客户端对象
	minioClient := MinioClient
	// 存储桶名称 在minio可视化界面中创建，并将access policy设置为public
	bucketName := "grape"
	// 通过设备ID和日期标记存储的照片
	objectName := deviceId + "/" + deviceId + "-" + getCurrentTime() + ".jpg"
	time.Sleep(time.Second * 2)
	resp, getErr := http.Get(imgUrl)
	fmt.Println("---------------StatusCode-----------------")
	fmt.Println(resp.StatusCode)
	if getErr != nil {
		log.Fatalln("get file fail!", getErr.Error())
		//return false
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalln("close body fail!", err.Error())
		}
	}(resp.Body)
	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		length := int64(len(body))
		file := bytes.NewReader(body)
		if err != nil {
			log.Fatalln("body read fail!", err.Error())
		}
		// 上传文件
		_, putErr := minioClient.PutObject(
			bucketName,
			objectName,
			file,
			length,
			minio.PutObjectOptions{ContentType: "image/jpeg"},
		)
		if putErr != nil {
			log.Fatalf("put file fail! %s", putErr.Error())
		}
		fmt.Println("snap picture upload successfully!")
	}
	return true
}

func task() {
	setDeviceSnapEnhanced("8L10C5FPAN06A2B")
}

func init() {
	logs.Info("每天上午10点，下午2点，4点进行抓拍")
	MinioClient = InitMinioClient()
	c := cron.New()
	spec := "0 10,14,17 * * *"
	_, err := c.AddFunc(spec, task)
	if err != nil {
		log.Fatalln("Failed to add job to cron", err.Error())
	}
	c.Start()
}
