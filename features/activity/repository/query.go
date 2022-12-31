package repository

import (
	"golang/features/activity"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type repoQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) activity.Repository {
	return &repoQuery{db: db}
}

func (rq *repoQuery) ShowAll() ([]activity.Core, error) {
	var resQry []Activity
	if err := rq.db.Find(&resQry).Error; err != nil {
		log.Error("error on show all: ", err.Error())
		return nil, err
	}

	res := ToCoreArray(resQry)
	return res, nil
}

func (rq *repoQuery) ShowOne(id uint) (activity.Core, error) {
	var resQry Activity
	if err := rq.db.First(&resQry, "id = ?", id).Error; err != nil {
		log.Error("error on show one: ", err.Error())
		return activity.Core{}, err
	}

	res := ToCore(resQry)
	return res, nil
}

func (rq *repoQuery) Insert(data activity.Core) (activity.Core, error) {
	var cnv Activity = FromCore(data)
	if err := rq.db.Create(&cnv).Error; err != nil {
		log.Error("error on insert: ", err.Error())
		return activity.Core{}, err
	}

	res := ToCore(cnv)
	return res, nil
}

func (rq *repoQuery) Edit(data activity.Core, id uint) (activity.Core, error) {
	var cnv Activity = FromCore(data)
	if err := rq.db.Where("id = ?", id).Updates(&cnv).Error; err != nil {
		log.Error("error on edit: ", err.Error())
		return activity.Core{}, err
	}

	if err := rq.db.First(&cnv, "id = ?", id).Error; err != nil {
		log.Error("error on finding edit: ", err.Error())
		return activity.Core{}, err
	}

	res := ToCore(cnv)
	return res, nil
}

func (rq *repoQuery) Remove(id uint) (activity.Core, error) {
	var data Activity
	if err := rq.db.Delete(&data, "id = ?", id).Error; err != nil {
		log.Error("error on remove: ", err.Error())
		return activity.Core{}, err
	}

	if err := rq.db.Unscoped().First(&data, "id = ?", id).Error; err != nil {
		log.Error("error on finding edit: ", err.Error())
		return activity.Core{}, err
	}

	res := ToCore(data)
	return res, nil
}
