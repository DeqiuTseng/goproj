package dao

import (
	"packagestudy/core"
	"packagestudy/model"
)

func InsertUser(user model.User) error {
	tx := core.DB.Create(&user)
	if tx.Error != nil {
		core.LOG.Println("insert user in do fail")
		return tx.Error
	}
	return nil
}

func GetUserByUsername(userName string) *model.User {
	var user *model.User
	tx := core.DB.Model(&model.User{}).Where("username=?", userName).Find(&user)
	if tx.Error != nil {
		core.LOG.Println("Query user by username fail")
	}
	return user
}
