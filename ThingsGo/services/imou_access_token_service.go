package services

import (
	"IOT/initialize/psql"
	"IOT/models"
	"github.com/google/uuid"
	"log"
)

type ImouAccessTokenService struct {
	SearchField []string
	//可作为条件的字段
	WhereField []string
	//可做为时间范围查询的字段
	TimeField []string
}

func (*ImouAccessTokenService) AddAccessToken(token string, expireTime int64) (bool, error) {
	// 使用uuid生成主键
	id := uuid.New().String()
	tx := psql.Mydb.Create(&models.ImouAccessToken{Id: id, AccessToken: token, ExpireTime: expireTime})
	if tx.Error != nil {
		return false, tx.Error
	}
	return true, nil
}

func (*ImouAccessTokenService) GetLastToken() (models.ImouAccessToken, error) {
	var lastToken models.ImouAccessToken

	// 记录未找到，打印信息，并持久化最新的access_token
	tx := psql.Mydb.Last(&lastToken)

	if tx.Error != nil {
		log.Println("get access token fail!", tx.Error)
		return models.ImouAccessToken{}, tx.Error
	} else {
		return lastToken, nil
	}
}

func (*ImouAccessTokenService) UpdateTokenById(id string, token string) (bool, error) {

	tx := psql.Mydb.Update("access_token", token).Where("id", id)

	if tx.Error != nil {
		return false, tx.Error
	} else {
		return true, nil
	}

}
