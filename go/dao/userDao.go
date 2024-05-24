package dao

import (
	"packagestudy/core"
	"packagestudy/do"
)

func InsertUser(user do.User) error {
	tx := core.DB.Create(&user)
	if tx.Error != nil {
		core.LOG.Println("insert user in do fail")
		return tx.Error
	}
	return nil
}

func GetUserByUsername(userName string) *do.User {
	var user *do.User
	tx := core.DB.Model(&do.User{}).Where("username=?", userName).Find(&user)
	if tx.Error != nil {
		core.LOG.Println("Query user by username fail")
	}
	return user
}
