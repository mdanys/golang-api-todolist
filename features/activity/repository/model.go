package repository

import (
	"golang/features/activity"

	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model
	Title string
	Email string
}

func FromCore(ac activity.Core) Activity {
	return Activity{
		Model: gorm.Model{ID: ac.ID, CreatedAt: ac.CreatedAt, UpdatedAt: ac.UpdatedAt},
		Title: ac.Title,
		Email: ac.Email,
	}
}

func ToCore(a Activity) activity.Core {
	return activity.Core{
		ID:        a.ID,
		Title:     a.Title,
		Email:     a.Email,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}
}

func ToCoreArray(aa []Activity) []activity.Core {
	var arr []activity.Core
	for _, val := range aa {
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
