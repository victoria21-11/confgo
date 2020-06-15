package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"domain"
)

type TagHandler struct {
	Usecase domain.TagUsecase
}

func Create(e *echo.Echo, usecase domain.TagUsecase) {
	handler := &TagHandler{
		Usecase: usecase,
	}

	e.GET("/tags/", handler.Index)
	e.GET("/tags/:id", handler.Show)
	e.POST("/tags", handler.Store)
	e.PUT("/tags/:id", handler.Store)
	e.DELETE("/tags/:id", handler.Delete)
}

func (handler *TagHandler) Index(c echo.Context) error {
	ctx := c.Request().Context()

	res := handler.Usecase.Index(ctx)

	return c.JSON(http.StatusOK, res)
}

func (handler *TagHandler) Show(c echo.Context) error {
	ctx := c.Request().Context()

	idP, _ := strconv.Atoi(c.Param("id"))
	id := uint(idP)

	res := handler.Usecase.Show(ctx, id)

	return c.JSON(http.StatusOK, res)
}

func (handler *TagHandler) Store(c echo.Context) error {
	var model domain.Tag
	c.Bind(&model)
	ctx := c.Request().Context()

	res := handler.Usecase.Store(ctx, &model)

	return c.JSON(http.StatusOK, res)
}

func (handler *TagHandler) Update(c echo.Context) error {
	var model domain.Tag
	c.Bind(&model)
	ctx := c.Request().Context()

	res := handler.Usecase.Update(ctx, &model)

	return c.JSON(http.StatusOK, res)
}

func (handler *TagHandler) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	idP, _ := strconv.Atoi(c.Param("id"))
	id := uint(idP)

	res := handler.Usecase.Delete(ctx, id)

	return c.JSON(http.StatusOK, res)
}