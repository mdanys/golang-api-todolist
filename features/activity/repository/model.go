package repository

import (
	"golang/features/activity"

	"gorm.io/gorm"
)

type Data struct {
	gorm.Model
	Title string
	Email string
}

func FromCore(c activity.Core) Data {
	return Data{
		Model: gorm.Model{ID: c.ID, CreatedAt: c.CreatedAt, UpdatedAt: c.UpdatedAt},
		Title: c.Title,
		Email: c.Email,
	}
}

func ToCore(d Data) activity.Core {
	return activity.Core{
		ID:        d.ID,
		Title:     d.Title,
		Email:     d.Email,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}

func ToCoreArray(da []Data) []activity.Core {
	var arr []activity.Core
	for _, val := range da {
		arr = append(arr, activity.Core{
			ID:        val.ID,
			Title:     val.Title,
			Email:     val.Email,
			CreatedAt: val.CreatedAt,
			UpdatedAt: val.UpdatedAt,
		})
	}
	return arr
}
