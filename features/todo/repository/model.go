package repository

import (
	"golang/features/todo"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	ActivityGroupID uint
	Title           string
	IsActive        bool
	Priority        string
}

func FromCore(c todo.Core) Todo {
	return Todo{
		Model:           gorm.Model{ID: c.ID, CreatedAt: c.CreatedAt, UpdatedAt: c.UpdatedAt},
		ActivityGroupID: c.ActivityGroupID,
		Title:           c.Title,
		IsActive:        c.IsActive,
		Priority:        c.Priority,
	}
}

func ToCore(t Todo) todo.Core {
	return todo.Core{
		ID:              t.ID,
		ActivityGroupID: t.ActivityGroupID,
		Title:           t.Title,
		IsActive:        t.IsActive,
		Priority:        t.Priority,
		CreatedAt:       t.CreatedAt,
		UpdatedAt:       t.UpdatedAt,
	}
}

func ToCoreArray(ta []Todo) []todo.Core {
	var arr []todo.Core
	for _, val := range ta {
		arr = append(arr, todo.Core{
			ID:              val.ID,
			ActivityGroupID: val.ActivityGroupID,
			Title:           val.Title,
			IsActive:        val.IsActive,
			Priority:        val.Priority,
			CreatedAt:       val.CreatedAt,
			UpdatedAt:       val.UpdatedAt,
		})
	}
	return arr
}
