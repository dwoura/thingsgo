package services

import (
	"IOT/models"
	uuid "IOT/utils"
	"errors"

	"IOT/initialize/psql"

	"gorm.io/gorm"
)

type CustomerService struct {
}

// Paginate 分页获取customer数据
func (*CustomerService) Paginate(name string, offset int, pageSize int) ([]models.Customer, int64) {
	var customers []models.Customer
	var count int64
	if name != "" {
		result := psql.Mydb.Model(&models.Customer{}).Where("name LIKE ?", "%"+name+"%").Limit(pageSize).Offset(offset).Find(&customers)
		psql.Mydb.Model(&models.Customer{}).Where("name LIKE ?", "%"+name+"%").Count(&count)
		if result.Error != nil {
			errors.Is(result.Error, gorm.ErrRecordNotFound)
		}
		if len(customers) == 0 {
			customers = []models.Customer{}
		}
		return customers, count
	} else {
		result := psql.Mydb.Model(&models.Customer{}).Limit(pageSize).Offset(offset).Find(&customers)
		psql.Mydb.Model(&models.Customer{}).Count(&count)
		if result.Error != nil {
			errors.Is(result.Error, gorm.ErrRecordNotFound)
		}
		if len(customers) == 0 {
			customers = []models.Customer{}
		}
		return customers, count
	}
}

// Add新增一条customer数据
func (*CustomerService) Add(title string, email string) (bool, string) {
	var uuid = uuid.GetUuid()
	customer := models.Customer{
		ID:    uuid,
		Title: title,
		Email: email,
	}
	result := psql.Mydb.Create(&customer)
	if result.Error != nil {
		errors.Is(result.Error, gorm.ErrRecordNotFound)
		return false, ""
	}
	return true, uuid
}

// 根据ID编辑一条customer数据
func (*CustomerService) Edit(id string, title string, email string, additional_info string, address string, address2 string, city string, country string, phone string, zip string) bool {
	result := psql.Mydb.Model(&models.Customer{}).Where("id = ?", id).Updates(map[string]interface{}{
		"title":           title,
		"email":           email,
		"additional_info": additional_info,
		"address":         address,
		"address2":        address2,
		"city":            city,
		"country":         country,
		"phone":           phone,
		"zip":             zip,
	})
	if result.Error != nil {
		errors.Is(result.Error, gorm.ErrRecordNotFound)
		return false
	}
	return true
}

// 根据ID删除一条customer数据
func (*CustomerService) Delete(id string) bool {
	result := psql.Mydb.Where("id = ?", id).Delete(&models.Customer{})
	if result.Error != nil {
		errors.Is(result.Error, gorm.ErrRecordNotFound)
		return false
	}
	return true
}

// GetCustomerByTitle 根据title获取一条customer数据
func (*CustomerService) GetCustomerByTitle(title string) (*models.Customer, int64) {
	var customer models.Customer
	result := psql.Mydb.Where("title = ?", title).First(&customer)
	if result.Error != nil {
		errors.Is(result.Error, gorm.ErrRecordNotFound)
	}
	return &customer, result.RowsAffected
}

// GetSameCustomerByTitle 根据title获取一条同名称的customer数据
func (*CustomerService) GetSameCustomerByTitle(title string, id string) (*models.Customer, int64) {
	var customer models.Customer
	result := psql.Mydb.Where("title = ? AND id <> ?", title, id).First(&customer)
	if result.Error != nil {
		errors.Is(result.Error, gorm.ErrRecordNotFound)
	}
	return &customer, result.RowsAffected
}

// GetCustomerByEmail 根据email获取一条customer数据
func (*CustomerService) GetUserByEmail(email string) (*models.Customer, int64) {
	var customer models.Customer
	result := psql.Mydb.Where("email = ?", email).First(&customer)
	if result.Error != nil {
		errors.Is(result.Error, gorm.ErrRecordNotFound)
	}
	return &customer, result.RowsAffected
}

// GetSameCustomerByEmail 根据email获取一条同名称的customer数据
func (*CustomerService) GetSameCustomerByEmail(email string, id string) (*models.Customer, int64) {
	var customer models.Customer
	result := psql.Mydb.Where("email = ? AND id <> ?", email, id).First(&customer)
	if result.Error != nil {
		errors.Is(result.Error, gorm.ErrRecordNotFound)
	}
	return &customer, result.RowsAffected
}

// GetCustomerById 根据id获取一条customer数据
func (*CustomerService) GetCustomerById(id string) (*models.Customer, int64) {
	var customer models.Customer
	result := psql.Mydb.Where("id = ?", id).First(&customer)
	if result.Error != nil {
		errors.Is(result.Error, gorm.ErrRecordNotFound)
	}
	return &customer, result.RowsAffected
}
