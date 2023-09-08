package dao

import (
	log "github.com/sirupsen/logrus"
	"golang-restful-api-framework/internal/model/do"
	"gorm.io/gorm"
)

type UserDAO interface {
	FindUserById(id int) (do.User, error)
}

type UserDAOImpl struct {
	db *gorm.DB
}

func InitUserDAO(db *gorm.DB) UserDAOImpl {
	//for test
	db.AutoMigrate(&do.User{})
	return UserDAOImpl{db: db}
}
func (u UserDAOImpl) FindUserById(id int) (do.User, error) {
	user := do.User{
		ID: id,
	}
	err := u.db.Preload("Role").First(&user).Error
	if err != nil {
		log.Error("Got and error when find user by id. Error: ", err)
		return do.User{}, err
	}
	return user, nil
}
