package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"domain"
)

type EventHandler struct {
	Usecase domain.EventUsecase
}

func Create(e *echo.Echo, usecase domain.EventUsecase) {
	handler := &EventHandler{
		Usecase: usecase,
	}

	e.GET("/events/", handler.Index)
	e.GET("/events/:id", handler.Show)
	e.POST("/events", handler.Store)
	e.PUT("/events/:id", handler.Store)
	e.DELETE("/events/:id", handler.Delete)
}

func (handler *EventHandler) Index(c echo.Context) error {
	ctx := c.Request().Context()

	res := handler.Usecase.Index(ctx)

	return c.JSON(http.StatusOK, res)
}

func (handler *EventHandler) Show(c echo.Context) error {
	ctx := c.Request().Context()

	idP, _ := strconv.Atoi(c.Param("id"))
	id := uint(idP)

	res := handler.Usecase.Show(ctx, id)

	return c.JSON(http.StatusOK, res)
}

func (handler *EventHandler) Store(c echo.Context) error {
	var model domain.Event
	c.Bind(&model)
	ctx := c.Request().Context()

	res := handler.Usecase.Store(ctx, &model)

	return c.JSON(http.StatusOK, res)
}

func (handler *EventHandler) Update(c echo.Context) error {
	var model domain.Event
	c.Bind(&model)
	ctx := c.Request().Context()

	res := handler.Usecase.Update(ctx, &model)

	return c.JSON(http.StatusOK, res)
}

func (handler *EventHandler) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	idP, _ := strconv.Atoi(c.Param("id"))
	id := uint(idP)

	res := handler.Usecase.Delete(ctx, id)

	return c.JSON(http.StatusOK, res)
}