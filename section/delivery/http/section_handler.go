package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"domain"
)

type SectionHandler struct {
	Usecase domain.SectionUsecase
}

func Create(e *echo.Echo, usecase domain.SectionUsecase) {
	handler := &SectionHandler{
		Usecase: usecase,
	}

	e.GET("/sections/", handler.Index)
	e.GET("/sections/:id", handler.Show)
	e.POST("/sections", handler.Store)
	e.PUT("/sections/:id", handler.Store)
	e.DELETE("/sections/:id", handler.Delete)
}

func (handler *SectionHandler) Index(c echo.Context) error {
	ctx := c.Request().Context()

	res := handler.Usecase.Index(ctx)

	return c.JSON(http.StatusOK, res)
}

func (handler *SectionHandler) Show(c echo.Context) error {
	ctx := c.Request().Context()

	idP, _ := strconv.Atoi(c.Param("id"))
	id := uint(idP)

	res := handler.Usecase.Show(ctx, id)

	return c.JSON(http.StatusOK, res)
}

func (handler *SectionHandler) Store(c echo.Context) error {
	var model domain.Section
	c.Bind(&model)
	ctx := c.Request().Context()

	res := handler.Usecase.Store(ctx, &model)

	return c.JSON(http.StatusOK, res)
}

func (handler *SectionHandler) Update(c echo.Context) error {
	var model domain.Section
	c.Bind(&model)
	ctx := c.Request().Context()

	res := handler.Usecase.Update(ctx, &model)

	return c.JSON(http.StatusOK, res)
}

func (handler *SectionHandler) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	idP, _ := strconv.Atoi(c.Param("id"))
	id := uint(idP)

	res := handler.Usecase.Delete(ctx, id)

	return c.JSON(http.StatusOK, res)
}