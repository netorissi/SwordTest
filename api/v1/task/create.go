package task

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/netorissi/SwordTest/api/middleware"
	"github.com/netorissi/SwordTest/model"
	"github.com/netorissi/SwordTest/shared"
)

// Task godoc
// @Summary      Create task
// @Description  Only tech can create task
// @Tags task
// @Param body body createRequest true "Request"
// @Produce  json
// @Success 201
// @Failure 400 {object} shared.Response{error=shared.Error}
// @Failure 401 {object} shared.Response{error=shared.Error}
// @Failure 500 {object} shared.Response{error=shared.Error}
// @Security ApiKeyAuth
// @Router /v1/tasks [post]
func (i *impl) create(c echo.Context) error {
	ss, err := middleware.GetSessionContext(c)
	if err != nil {
		return err
	}

	input := new(createRequest)

	if err := c.Bind(&input); err != nil {
		return shared.NewError(http.StatusBadRequest, shared.ErrInvalidBind.Error())
	}

	if err := c.Validate(input); err != nil {
		return shared.NewError(http.StatusBadRequest, err.Error())
	}

	err = i.usecase.Task.Create(c.Request().Context(), model.Task{
		UserID:  ss.UserID,
		Summary: input.Summary,
	})

	if err != nil {
		return shared.NewError(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusCreated)
}
