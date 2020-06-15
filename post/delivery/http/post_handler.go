package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"domain"
)

type PostHandler struct {
	Usecase domain.PostUsecase
}

func Create(e *echo.Echo, usecase domain.PostUsecase) {
	handler := &PostHandler{
		Usecase: usecase,
	}

	e.GET("/posts/", handler.Index)
	e.GET("/posts/:id", handler.Show)
	e.POST("/posts", handler.Store)
	e.PUT("/posts/:id", handler.Store)
	e.DELETE("/posts/:id", handler.Delete)
}

func (handler *PostHandler) Index(c echo.Context) error {
	ctx := c.Request().Context()

	res := handler.Usecase.Index(ctx)

	return c.JSON(http.StatusOK, res)
}

func (handler *PostHandler) Show(c echo.Context) error {
	ctx := c.Request().Context()

	idP, _ := strconv.Atoi(c.Param("id"))
	id := uint(idP)

	res := handler.Usecase.Show(ctx, id)

	return c.JSON(http.StatusOK, res)
}

func (handler *PostHandler) Store(c echo.Context) error {
	var model domain.Post
	c.Bind(&model)
	ctx := c.Request().Context()

	res := handler.Usecase.Store(ctx, &model)

	return c.JSON(http.StatusOK, res)
}

func (handler *PostHandler) Update(c echo.Context) error {
	var model domain.Post
	c.Bind(&model)
	ctx := c.Request().Context()

	res := handler.Usecase.Update(ctx, &model)

	return c.JSON(http.StatusOK, res)
}

func (handler *PostHandler) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	idP, _ := strconv.Atoi(c.Param("id"))
	id := uint(idP)

	res := handler.Usecase.Delete(ctx, id)

	return c.JSON(http.StatusOK, res)
}