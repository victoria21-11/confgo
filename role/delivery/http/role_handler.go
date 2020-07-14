package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/go-playground/validator"

	"domain"
)

type RoleHandler struct {
	Usecase domain.RoleUsecase

	validator *validator.Validate
}

func Create(e *echo.Echo, usecase domain.RoleUsecase) {
	handler := &RoleHandler{
		Usecase: usecase,
	}

	e.GET("/roles", handler.Index)
	e.GET("/roles/:id", handler.Show)
	e.POST("/roles", handler.Store)
	e.PUT("/roles/:id", handler.Update)
	e.DELETE("/roles/:id", handler.Delete)
}

func (handler *RoleHandler) Index(c echo.Context) error {
	ctx := c.Request().Context()

	res := handler.Usecase.Index(ctx)

	return c.JSON(http.StatusOK, res)
}

func (handler *RoleHandler) Show(c echo.Context) error {
	ctx := c.Request().Context()

	idP, _ := strconv.Atoi(c.Param("id"))
	id := uint(idP)

	res := handler.Usecase.Show(ctx, id)

	return c.JSON(http.StatusOK, res)
}

func (handler *RoleHandler) Store(c echo.Context) error {
	var model domain.Role

	c.Bind(&model)

	ctx := c.Request().Context()

	errors, err := handler.validate(model)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, errors)
	}

	res := handler.Usecase.Store(ctx, &model)

	return c.JSON(http.StatusOK, res)
}

func (handler *RoleHandler) validate(model domain.Role) (errors map[string]string, err error) {
	errors = map[string]string{}

	handler.validator = validator.New()

	err = handler.validator.Struct(model)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = err.ActualTag()
		}
	}

	return errors, err
}

func (handler *RoleHandler) Update(c echo.Context) error {
	var model domain.Role

	c.Bind(&model)

	errors, err := handler.validate(model)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, errors)
	}

	ctx := c.Request().Context()

	res := handler.Usecase.Update(ctx, &model)

	return c.JSON(http.StatusOK, res)
}

func (handler *RoleHandler) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	idP, _ := strconv.Atoi(c.Param("id"))
	id := uint(idP)

	res := handler.Usecase.Delete(ctx, id)

	return c.JSON(http.StatusOK, res)
}