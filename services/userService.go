package services

import (
	"errors"
	"prometheus_devops/db"
	"prometheus_devops/model"
)

func GetUser(id string) (user *model.User, int642 int64) {
	user = &model.User{}
	row := db.DB.Where("id = ?", id).Find(&user).RowsAffected

	return user, row
}
func GetAllUser() []*model.User {
	userlist := []*model.User{}
	db.DB.Find(&userlist)
	return userlist
}

func DelteUser(id string) bool {
	user := &model.User{}
	row := db.DB.Where("id = ?", id).Delete(&user).RowsAffected
	if row == 1 {
		return true
	}
	return false
}

func UpdateUser(user *model.User) bool {

	row := db.DB.Updates(&user).RowsAffected
	if row == 1 {
		return true
	}
	return false
}

func GetUserByName(name string) (*model.User, error) {
	user := &model.User{}
	row := db.DB.Where("name = ? ", name).Find(&user).RowsAffected

	if row == 1 {
		return user, nil
	}
	return nil, errors.New("the user not register")
}
