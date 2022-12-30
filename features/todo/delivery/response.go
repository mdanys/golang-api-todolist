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

type TodoResponse struct {
	ID              uint      `json:"id" form:"id"`
	ActivityGroupID uint      `json:"activity_group_id" form:"activity_group_id"`
	Title           string    `json:"title" form:"title"`
	IsActive        bool      `json:"is_active" form:"is_active"`
	Priority        string    `json:"priority" form:"priority"`
	CreatedAt       time.Time `json:"createdAt" form:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt" form:"updatedAt"`
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "all":
		var arr []TodoResponse
		cnv := core.([]todo.Core)
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
	}

	return res
}
