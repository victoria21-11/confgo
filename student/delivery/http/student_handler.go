package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"domain"
)

type StudentHandler struct {
	Usecase domain.StudentUsecase
}

func Create(e *echo.Echo, usecase domain.StudentUsecase) {
	handler := &StudentHandler{
		Usecase: usecase,
	}

	e.GET("/students/", handler.Index)
	e.GET("/students/:id", handler.Show)
	e.POST("/students", handler.Store)
	e.PUT("/students/:id", handler.Store)
	e.DELETE("/students/:id", handler.Delete)
}

func (handler *StudentHandler) Index(c echo.Context) error {
	ctx := c.Request().Context()

	res := handler.Usecase.Index(ctx)

	return c.JSON(http.StatusOK, res)
}

func (handler *StudentHandler) Show(c echo.Context) error {
	ctx := c.Request().Context()

	idP, _ := strconv.Atoi(c.Param("id"))
	id := uint(idP)

	res := handler.Usecase.Show(ctx, id)

	return c.JSON(http.StatusOK, res)
}

func (handler *StudentHandler) Store(c echo.Context) error {
	var model domain.Student
	c.Bind(&model)
	ctx := c.Request().Context()

	res := handler.Usecase.Store(ctx, &model)

	return c.JSON(http.StatusOK, res)
}

func (handler *StudentHandler) Update(c echo.Context) error {
	var model domain.Student
	c.Bind(&model)
	ctx := c.Request().Context()

	res := handler.Usecase.Update(ctx, &model)

	return c.JSON(http.StatusOK, res)
}

func (handler *StudentHandler) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	idP, _ := strconv.Atoi(c.Param("id"))
	id := uint(idP)

	res := handler.Usecase.Delete(ctx, id)

	return c.JSON(http.StatusOK, res)
}