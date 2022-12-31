package repository

import (
	"golang/features/todo"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type repoQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) todo.Repository {
	return &repoQuery{db: db}
}

func (rq *repoQuery) ShowAll(query string) ([]todo.Core, error) {
	var resQry []Todo
	if query != "" {
		if err := rq.db.Where("activity_group_id = ?", query).Find(&resQry).Error; err != nil {
			log.Error("error on show all: ", err.Error())
			return []todo.Core{}, err
		}
	} else {
		if err := rq.db.Find(&resQry).Error; err != nil {
			log.Error("error on show all: ", err.Error())
			return []todo.Core{}, err
		}
	}

	res := ToCoreArray(resQry)
	return res, nil
}

func (rq *repoQuery) ShowOne(id uint) (todo.Core, error) {
	var resQry Todo
	if err := rq.db.First(&resQry, "id = ?", id).Error; err != nil {
		log.Error("error on show one: ", err.Error())
		return todo.Core{}, err
	}

	res := ToCore(resQry)
	return res, nil
}

func (rq *repoQuery) Insert(data todo.Core) (todo.Core, error) {
	var cnv Todo = FromCore(data)
	if err := rq.db.Create(&cnv).Error; err != nil {
		log.Error("error on insert: ", err.Error())
		return todo.Core{}, err
	}

	res := ToCore(cnv)
	return res, nil
}

func (rq *repoQuery) Edit(data todo.Core, id uint) (todo.Core, error) {
	var cnv Todo = FromCore(data)
	if err := rq.db.Where("id = ?", id).Updates(&cnv).Error; err != nil {
		log.Error("error on edit: ", err.Error())
		return todo.Core{}, err
	}

	if err := rq.db.First(&cnv, "id = ?", id).Error; err != nil {
		log.Error("error on finding edit: ", err.Error())
		return todo.Core{}, err
	}

	res := ToCore(cnv)
	return res, nil
}

func (rq *repoQuery) Remove(id uint) (todo.Core, error) {
	var data Todo
	if err := rq.db.Delete(&data, "id = ?", id).Error; err != nil {
		log.Error("error on remove: ", err.Error())
		return todo.Core{}, err
	}

	if err := rq.db.Unscoped().First(&data, "id = ?", id).Error; err != nil {
		log.Error("error on finding edit: ", err.Error())
		return todo.Core{}, err
	}

	res := ToCore(data)
	return res, nil
}
