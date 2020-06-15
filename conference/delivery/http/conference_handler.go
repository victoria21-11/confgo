package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"domain"
)

type ConferenceHandler struct {
	Usecase domain.ConferenceUsecase
}

func Create(e *echo.Echo, usecase domain.ConferenceUsecase) {
	handler := &ConferenceHandler{
		Usecase: usecase,
	}

	e.GET("/conferences/", handler.Index)
	e.GET("/conferences/:id", handler.Show)
	e.POST("/conferences", handler.Store)
	e.PUT("/conferences/:id", handler.Store)
	e.DELETE("/conferences/:id", handler.Delete)
}

func (handler *ConferenceHandler) Index(c echo.Context) error {
	ctx := c.Request().Context()

	res := handler.Usecase.Index(ctx)

	return c.JSON(http.StatusOK, res)
}

func (handler *ConferenceHandler) Show(c echo.Context) error {
	ctx := c.Request().Context()

	idP, _ := strconv.Atoi(c.Param("id"))
	id := uint(idP)

	res := handler.Usecase.Show(ctx, id)

	return c.JSON(http.StatusOK, res)
}

func (handler *ConferenceHandler) Store(c echo.Context) error {
	var model domain.Conference
	c.Bind(&model)
	ctx := c.Request().Context()

	res := handler.Usecase.Store(ctx, &model)

	return c.JSON(http.StatusOK, res)
}

func (handler *ConferenceHandler) Update(c echo.Context) error {
	var model domain.Conference
	c.Bind(&model)
	ctx := c.Request().Context()

	res := handler.Usecase.Update(ctx, &model)

	return c.JSON(http.StatusOK, res)
}

func (handler *ConferenceHandler) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	idP, _ := strconv.Atoi(c.Param("id"))
	id := uint(idP)

	res := handler.Usecase.Delete(ctx, id)

	return c.JSON(http.StatusOK, res)
}