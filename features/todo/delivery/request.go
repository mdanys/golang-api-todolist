package delivery

import "golang/features/todo"

type CreateFormat struct {
	Title           string `json:"title" form:"title"`
	ActivityGroupID uint   `json:"activity_group_id" form:"activity_group_id"`
	IsActive        bool   `json:"is_active" form:"is_active"`
}

type UpdateFormat struct {
	Title    string `json:"title" form:"title"`
	Priority string `json:"priority" form:"priority"`
	IsActive bool   `json:"is_active" form:"is_active"`
	Status   string `json:"status" form:"status"`
}

type DeleteFormat struct {
	Title string `json:"title" form:"title"`
}

func ToCore(i interface{}) todo.Core {
	switch i.(type) {
	case CreateFormat:
		cnv := i.(CreateFormat)
		return todo.Core{
			Title:           cnv.Title,
			ActivityGroupID: cnv.ActivityGroupID,
			IsActive:        cnv.IsActive,
		}
	case UpdateFormat:
		cnv := i.(UpdateFormat)
		return todo.Core{
			Title:    cnv.Title,
			Priority: cnv.Priority,
			IsActive: cnv.IsActive,
		}
	case DeleteFormat:
		cnv := i.(DeleteFormat)
		return todo.Core{
			Title: cnv.Title,
		}
	}

	return todo.Core{}
}
