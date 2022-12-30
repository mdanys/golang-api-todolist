package delivery

import (
	"fmt"
	"golang/features/activity"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type activityHandler struct {
	srv activity.Service
}

func New(e *echo.Echo, srv activity.Service) {
	handler := activityHandler{srv: srv}
	e.GET("/activity-groups", handler.GetAll())
	e.GET("/activity-groups/:id", handler.GetOne())
	e.POST("/activity-groups", handler.Create())
	e.PATCH("/activity-groups/:id", handler.Update())
	e.DELETE("/activity-groups/:id", handler.Delete())
}

func (ah *activityHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := ah.srv.GetAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse("internal server error", "there is a problem on server"))
		}

		return c.JSON(http.StatusOK, SuccessResponse("Success", "Success", ToResponse(res, "all")))
	}
}

func (ah *activityHandler) GetOne() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		cnv, err := strconv.Atoi(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, FailResponse("Not Found", fmt.Sprintf("Activity with ID %d Not Found", cnv)))
		}

		res, err := ah.srv.GetOne(uint(cnv))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse("internal server error", "there is a problem on server"))
		}

		return c.JSON(http.StatusOK, SuccessResponse("Success", "Success", ToResponse(res, "data")))
	}
}

func (ah *activityHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input CreateFormat
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("Bad Request", "title cannot be null"))
		}

		cnv := ToCore(input)
		res, err := ah.srv.Create(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse("internal server error", "there is a problem on server"))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("Success", "Success", ToResponse(res, "data")))
	}
}

func (ah *activityHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input UpdateFormat
		if err := c.Bind(input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("Bad Request", "title cannot be null"))
		}

		id := c.Param("id")
		cnv, err := strconv.Atoi(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, FailResponse("Not Found", fmt.Sprintf("Activity with ID %d Not Found", cnv)))
		}

		conv := ToCore(input)
		res, err := ah.srv.Update(conv, uint(cnv))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse("internal server error", "there is a problem on server"))
		}

		return c.JSON(http.StatusOK, SuccessResponse("Success", "Success", ToResponse(res, "data")))
	}
}

func (ah *activityHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		cnv, err := strconv.Atoi(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, FailResponse("Not Found", fmt.Sprintf("Activity with ID %d Not Found", cnv)))
		}

		er := ah.srv.Delete(uint(cnv))
		if er != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse("internal server error", "there is a problem on server"))
		}

		return c.JSON(http.StatusOK, SuccessResponse("Success", "Success", ToResponse(er, "data")))
	}
}
