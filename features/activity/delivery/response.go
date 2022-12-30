package delivery

import (
	"golang/features/activity"
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

type ActivityResponse struct {
	ID        uint      `json:"id" form:"id"`
	Title     string    `json:"title" form:"title"`
	Email     string    `json:"email" form:"email"`
	CreatedAt time.Time `json:"createdAt" form:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" form:"updatedAt"`
}

func ToResponse(core interface{}, code string) interface{} {
	var res interface{}
	switch code {
	case "all":
		var arr []ActivityResponse
		cnv := core.([]activity.Core)
		for _, val := range cnv {
			arr = append(arr, ActivityResponse{
				ID:        val.ID,
				Title:     val.Title,
				Email:     val.Email,
				CreatedAt: val.CreatedAt,
				UpdatedAt: val.UpdatedAt,
			})
		}
		res = arr
	case "data":
		cnv := core.(activity.Core)
		res = ActivityResponse{
			ID:        cnv.ID,
			Title:     cnv.Title,
			Email:     cnv.Email,
			CreatedAt: cnv.CreatedAt,
			UpdatedAt: cnv.UpdatedAt,
		}
	}

	return res
}
