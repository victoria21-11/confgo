package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"domain"
)

type CoordinatorHandler struct {
	Usecase domain.CoordinatorUsecase
}

func Create(e *echo.Echo, usecase domain.CoordinatorUsecase) {
	handler := &CoordinatorHandler{
		Usecase: usecase,
	}

	e.GET("/coordinators/", handler.Index)
	e.GET("/coordinators/:id", handler.Show)
	e.POST("/coordinators", handler.Store)
	e.PUT("/coordinators/:id", handler.Store)
	e.DELETE("/coordinators/:id", handler.Delete)
}

func (handler *CoordinatorHandler) Index(c echo.Context) error {
	ctx := c.Request().Context()

	res := handler.Usecase.Index(ctx)

	return c.JSON(http.StatusOK, res)
}

func (handler *CoordinatorHandler) Show(c echo.Context) error {
	ctx := c.Request().Context()

	idP, _ := strconv.Atoi(c.Param("id"))
	id := uint(idP)

	res := handler.Usecase.Show(ctx, id)

	return c.JSON(http.StatusOK, res)
}

func (handler *CoordinatorHandler) Store(c echo.Context) error {
	var model domain.Coordinator
	c.Bind(&model)
	ctx := c.Request().Context()

	res := handler.Usecase.Store(ctx, &model)

	return c.JSON(http.StatusOK, res)
}

func (handler *CoordinatorHandler) Update(c echo.Context) error {
	var model domain.Coordinator
	c.Bind(&model)
	ctx := c.Request().Context()

	res := handler.Usecase.Update(ctx, &model)

	return c.JSON(http.StatusOK, res)
}

func (handler *CoordinatorHandler) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	idP, _ := strconv.Atoi(c.Param("id"))
	id := uint(idP)

	res := handler.Usecase.Delete(ctx, id)

	return c.JSON(http.StatusOK, res)
}