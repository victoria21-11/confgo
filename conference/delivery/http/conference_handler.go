package http

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/go-playground/validator"

	"domain"
)

type ConferenceHandler struct {
	Usecase domain.ConferenceUsecase

	validator *validator.Validate
}

func Create(e *echo.Echo, usecase domain.ConferenceUsecase) {
	handler := &ConferenceHandler{
		Usecase: usecase,
	}

	e.GET("/conferences", handler.Index)
	e.GET("/conferences/:id", handler.Show)
	e.POST("/conferences", handler.Store)
	e.PUT("/conferences/:id", handler.Update)
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

	startDate := c.FormValue("start_date")
	endDate := c.FormValue("end_date")

	model.StartDate, _ = time.Parse("2006-01-02", startDate)
	model.EndDate, _ = time.Parse("2006-01-02", endDate)

	errors, err := handler.validate(model)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, errors)
	}

	res := handler.Usecase.Store(ctx, &model)

	return c.JSON(http.StatusOK, res)
}

func (handler *ConferenceHandler) validate(model domain.Conference) (errors map[string]string, err error) {
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

func (handler *ConferenceHandler) Update(c echo.Context) error {
	var model domain.Conference

	c.Bind(&model)

	startDate := c.FormValue("start_date")
	endDate := c.FormValue("end_date")

	model.StartDate, _ = time.Parse("2006-01-02", startDate)
	model.EndDate, _ = time.Parse("2006-01-02", endDate)

	errors, err := handler.validate(model)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, errors)
	}

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