package dao

import (
	log "github.com/sirupsen/logrus"
	"golang-restful-api-framework/internal/model/do"
	"gorm.io/gorm"
)

type TaskDAO interface {
	FindTaskById(id int) (do.Task, error)
}

type TaskDAOImpl struct {
	db *gorm.DB
}

func InitTaskDAO(db *gorm.DB) TaskDAOImpl {
	//for test
	db.AutoMigrate(&do.Task{})
	return TaskDAOImpl{db: db}
}
func (u TaskDAOImpl) FindTaskById(id int) (do.Task, error) {
	task := do.Task{
		ID: id,
	}
	err := u.db.Preload("Task").First(&task).Error
	if err != nil {
		log.Error("Got and error when find user by id. Error: ", err)
		return do.Task{}, err
	}
	return task, nil
}
