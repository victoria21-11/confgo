package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"domain"
)

type UniversityHandler struct {
	Usecase domain.UniversityUsecase
}

func Create(e *echo.Echo, usecase domain.UniversityUsecase) {
	handler := &UniversityHandler{
		Usecase: usecase,
	}

	e.GET("/universities/", handler.Index)
	e.GET("/universities/:id", handler.Show)
	e.POST("/universities", handler.Store)
	e.PUT("/universities/:id", handler.Store)
	e.DELETE("/universities/:id", handler.Delete)
}

func (handler *UniversityHandler) Index(c echo.Context) error {
	ctx := c.Request().Context()

	res := handler.Usecase.Index(ctx)

	return c.JSON(http.StatusOK, res)
}

func (handler *UniversityHandler) Show(c echo.Context) error {
	ctx := c.Request().Context()

	idP, _ := strconv.Atoi(c.Param("id"))
	id := uint(idP)

	res := handler.Usecase.Show(ctx, id)

	return c.JSON(http.StatusOK, res)
}

func (handler *UniversityHandler) Store(c echo.Context) error {
	var model domain.University
	c.Bind(&model)
	ctx := c.Request().Context()

	res := handler.Usecase.Store(ctx, &model)

	return c.JSON(http.StatusOK, res)
}

func (handler *UniversityHandler) Update(c echo.Context) error {
	var model domain.University
	c.Bind(&model)
	ctx := c.Request().Context()

	res := handler.Usecase.Update(ctx, &model)

	return c.JSON(http.StatusOK, res)
}

func (handler *UniversityHandler) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	idP, _ := strconv.Atoi(c.Param("id"))
	id := uint(idP)

	res := handler.Usecase.Delete(ctx, id)

	return c.JSON(http.StatusOK, res)
}