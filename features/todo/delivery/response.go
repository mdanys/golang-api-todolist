package delivery

import (
	"golang/features/todo"
	"time"
)

func SuccessResponse(stat string, msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":  stat,
		"message": msg,
		"data":    data,
	}
}

func FailResponse(stat string, msg string) map[string]interface{} {
	return map[string]interface{}{
		"status":  stat,
		"message": msg,
	}
}

type InsertResponse struct {
	ID              uint      `json:"id" form:"id"`
	Title           string    `json:"title" form:"title"`
	ActivityGroupID uint      `json:"activity_group_id" form:"activity_group_id"`
	IsActive        bool      `json:"is_active" form:"is_active"`
	Priority        string    `json:"priority" form:"priority"`
	UpdatedAt       time.Time `json:"updatedAt" form:"updatedAt"`
	CreatedAt       time.Time `json:"createdAt" form:"createdAt"`
}
type TodoResponse struct {
	ID              uint      `json:"id" form:"id"`
	ActivityGroupID uint      `json:"activity_group_id" form:"activity_group_id"`
	Title           string    `json:"title" form:"title"`
	IsActive        bool      `json:"is_active" form:"is_active"`
	Priority        string    `json:"priority" form:"priority"`
	CreatedAt       time.Time `json:"createdAt" form:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt" form:"updatedAt"`
}

type DeleteResponse struct {
	ID uint `json:"-" form:"-"`
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "all":
		var arr []TodoResponse
		cnv := core.([]todo.Core)
		if len(cnv) == 0 {
			arr = make([]TodoResponse, 0)
		} else {
			for _, val := range cnv {
				arr = append(arr, TodoResponse{
					ID:              val.ID,
					ActivityGroupID: val.ActivityGroupID,
					Title:           val.Title,
					IsActive:        val.IsActive,
					Priority:        val.Priority,
					CreatedAt:       val.CreatedAt,
					UpdatedAt:       val.UpdatedAt,
				})
			}
		}
		res = arr
	case "data":
		cnv := core.(todo.Core)
		res = TodoResponse{
			ID:              cnv.ID,
			ActivityGroupID: cnv.ActivityGroupID,
			Title:           cnv.Title,
			IsActive:        cnv.IsActive,
			Priority:        cnv.Priority,
			CreatedAt:       cnv.CreatedAt,
			UpdatedAt:       cnv.UpdatedAt,
		}
	case "insert":
		cnv := core.(todo.Core)
		res = InsertResponse{
			ID:              cnv.ID,
			Title:           cnv.Title,
			ActivityGroupID: cnv.ActivityGroupID,
			IsActive:        cnv.IsActive,
			Priority:        cnv.Priority,
			UpdatedAt:       cnv.UpdatedAt,
			CreatedAt:       cnv.CreatedAt,
		}
	case "delete":
		cnv := core.(todo.Core)
		res = DeleteResponse{
			ID: cnv.ID,
		}
	}

	return res
}
