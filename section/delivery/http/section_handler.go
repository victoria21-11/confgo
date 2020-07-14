package http

import (
	"net/http"
	"strconv"
	"fmt"

	"github.com/labstack/echo"
	"github.com/go-playground/validator"

	"domain"
)

type SectionHandler struct {
	Usecase domain.SectionUsecase

	validator *validator.Validate
}

func Create(e *echo.Echo, usecase domain.SectionUsecase) {
	handler := &SectionHandler{
		Usecase: usecase,
	}

	e.GET("/sections", handler.Index)
	e.GET("/sections/:id", handler.Show)
	e.POST("/sections", handler.Store)
	e.PUT("/sections/:id", handler.Update)
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

	conferenceId, _ := strconv.Atoi(c.FormValue("conference_id"))
	coordinatorId, _ := strconv.Atoi(c.FormValue("coordinator_id"))
	
	model.ConferenceID = uint(conferenceId)
	model.CoordinatorID = uint(coordinatorId)

	errors, err := handler.validate(model)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, errors)
	}

	ctx := c.Request().Context()

	res, err := handler.Usecase.Store(ctx, &model)

	fmt.Print(err)
	if err != nil {
		res := map[string]string{
			"title": err.Error(),
		}
		return c.JSON(http.StatusUnprocessableEntity, res)
	}

	return c.JSON(http.StatusOK, res)
}

func (handler *SectionHandler) Update(c echo.Context) error {
	var model domain.Section

	c.Bind(&model)

	ctx := c.Request().Context()

	err := handler.Usecase.Update(ctx, &model)

	if err != nil {
		res := map[string]string{
			"title": err.Error(),
		}
		return c.JSON(http.StatusUnprocessableEntity, res)
	}

	return c.JSON(http.StatusOK, err)
}

func (handler *SectionHandler) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	idP, _ := strconv.Atoi(c.Param("id"))
	id := uint(idP)

	res := handler.Usecase.Delete(ctx, id)

	return c.JSON(http.StatusOK, res)
}

func (handler *SectionHandler) validate(model domain.Section) (errors map[string]string, err error) {
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