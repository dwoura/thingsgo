package services

import (
	"IOT/models"
	uuid "IOT/utils"
	"errors"

	"IOT/initialize/psql"

	"gorm.io/gorm"
)

type LogoService struct {
}

// 获取logo配置
func (*LogoService) GetLogo() models.Logo {
	var Logos []models.Logo
	var Logo models.Logo
	result := psql.Mydb.Model(&models.Logo{}).Limit(1).Find(&Logos)
	if result.Error != nil {
		errors.Is(result.Error, gorm.ErrRecordNotFound)
	}
	if len(Logos) == 0 {
		Logo = models.Logo{}
	} else {
		Logo = Logos[0]
	}
	return Logo
}

// Add新增一条Logo数据
func (*LogoService) Add(logo models.Logo) (string, error) {
	var uuid = uuid.GetUuid()
	logo.Id = uuid
	result := psql.Mydb.Create(&logo)
	if result.Error != nil {
		return "", result.Error
	}
	return uuid, nil
}

// 根据ID编辑一条Logo数据
func (*LogoService) Edit(logo models.Logo) error {
	result := psql.Mydb.Save(&logo)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
