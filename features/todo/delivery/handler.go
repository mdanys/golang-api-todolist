package delivery

import (
	"fmt"
	"golang/features/todo"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type todoHandler struct {
	srv todo.Service
}

func New(e *echo.Echo, srv todo.Service) {
	handler := todoHandler{srv: srv}
	e.GET("/todo-items", handler.GetAll())
	e.GET("/todo-items/:id", handler.GetOne())
	e.POST("/todo-items", handler.Create())
	e.PATCH("/todo-items/:id", handler.Update())
	e.DELETE("/todo-items/:id", handler.Delete())
}

func (th *todoHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.QueryParam("activity_group_id")

		res, err := th.srv.GetAll(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse("internal server error", "there is a problem on server"))
		}

		return c.JSON(http.StatusOK, SuccessResponse("Success", "Success", ToResponse(res, "all")))
	}
}

func (th *todoHandler) GetOne() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		cnv, _ := strconv.Atoi(id)

		res, err := th.srv.GetOne(uint(cnv))
		if err != nil {
			return c.JSON(http.StatusNotFound, FailResponse("Not Found", fmt.Sprintf("Todo with ID %d Not Found", cnv)))
		}

		return c.JSON(http.StatusOK, SuccessResponse("Success", "Success", ToResponse(res, "data")))
	}
}

func (th *todoHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input CreateFormat
		c.Bind(&input)

		if input.Title == "" {
			return c.JSON(http.StatusBadRequest, FailResponse("Bad Request", "title cannot be null"))
		}

		cnv := ToCore(input)
		res, err := th.srv.Create(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse("internal server error", "there is a problem on server"))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("Success", "Success", ToResponse(res, "insert")))
	}
}

func (th *todoHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input UpdateFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("Bad Request", "title cannot be null"))
		}

		id := c.Param("id")
		cnv, _ := strconv.Atoi(id)

		conv := ToCore(input)
		res, err := th.srv.Update(conv, uint(cnv))
		if err != nil {
			return c.JSON(http.StatusNotFound, FailResponse("Not Found", fmt.Sprintf("Todo with ID %d Not Found", cnv)))
		}

		return c.JSON(http.StatusOK, SuccessResponse("Success", "Success", ToResponse(res, "data")))
	}
}

func (th *todoHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input DeleteFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("Bad Request", "title cannot be null"))
		}

		id := c.Param("id")
		cnv, _ := strconv.Atoi(id)

		res, err := th.srv.Delete(uint(cnv))
		if err != nil {
			return c.JSON(http.StatusNotFound, FailResponse("Not Found", fmt.Sprintf("Todo with ID %d Not Found", cnv)))
		}

		return c.JSON(http.StatusOK, SuccessResponse("Success", "Success", ToResponse(res, "delete")))
	}
}
