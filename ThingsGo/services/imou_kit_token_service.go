package services

import (
	"IOT/initialize/psql"
	"IOT/models"
	"github.com/google/uuid"
	"log"
)

type ImouKitTokenService struct {
	SearchField []string
	//可作为条件的字段
	WhereField []string
	//可做为时间范围查询的字段
	TimeField []string
}

func (*ImouKitTokenService) GetImouKitTokenList() []models.ImouKitToken {

	var kitTokens []models.ImouKitToken

	tx := psql.Mydb.Find(&kitTokens)
	if tx.Error != nil {
		log.Fatalf("get kit token list fail!", tx.Error.Error())
		return nil
	}
	return kitTokens
}

func (*ImouKitTokenService) UpdateKitToken(imouKitToken models.ImouKitToken, token string) (bool, error) {
	tx := psql.Mydb.Model(&imouKitToken).Update("kit_token", token)
	if tx.Error != nil {
		log.Fatalf("update kit token fail!", tx.Error)
		return false, tx.Error
	}
	return true, nil
}

func (*ImouKitTokenService) AddDevice(deviceId string) (bool, error) {
	id := uuid.New().String()
	kitToken := models.ImouKitToken{Id: id, DeviceId: deviceId}
	tx := psql.Mydb.Create(kitToken)
	if tx.Error != nil {
		log.Fatalf("add deivce fail!", tx.Error)
		return false, tx.Error
	}
	return true, nil
}
