package delivery

import "golang/features/activity"

type CreateFormat struct {
	Title string `json:"title" form:"title" binding:"required"`
	Email string `json:"email" form:"email" binding:"required"`
}

type UpdateFormat struct {
	Title string `json:"title" form:"title" binding:"required"`
}

func ToCore(i interface{}) activity.Core {
	switch i.(type) {
	case CreateFormat:
		cnv := i.(CreateFormat)
		return activity.Core{
			Title: cnv.Title,
			Email: cnv.Email,
		}
	case UpdateFormat:
		cnv := i.(UpdateFormat)
		return activity.Core{
			Title: cnv.Title,
		}
	}

	return activity.Core{}
}
