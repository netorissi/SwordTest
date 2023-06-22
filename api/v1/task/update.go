package task

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/netorissi/SwordTest/api/middleware"
	"github.com/netorissi/SwordTest/model"
	"github.com/netorissi/SwordTest/shared"
)

// Task godoc
// @Summary      Update task
// @Description  Only owner can update task
// @Tags task
// @Param id path string true "Task ID"
// @Param body body updateRequest true "Request"
// @Produce  json
// @Success 204
// @Failure 400 {object} shared.Response{error=shared.Error}
// @Failure 401 {object} shared.Response{error=shared.Error}
// @Failure 500 {object} shared.Response{error=shared.Error}
// @Security ApiKeyAuth
// @Router /v1/tasks/{id} [put]
func (i *impl) update(c echo.Context) error {
	taskID := c.Param("id")
	if len(taskID) == 0 {
		return shared.NewError(http.StatusBadRequest, shared.ErrInvalidPayload.Error())
	}

	ss, err := middleware.GetSessionContext(c)
	if err != nil {
		return err
	}

	input := new(updateRequest)

	if err := c.Bind(&input); err != nil {
		return shared.NewError(http.StatusBadRequest, shared.ErrInvalidBind.Error())
	}

	if err := c.Validate(input); err != nil {
		return shared.NewError(http.StatusBadRequest, err.Error())
	}

	taskIDFormat, err := strconv.Atoi(taskID)
	if err != nil {
		return shared.NewError(http.StatusBadRequest, shared.ErrInvalidPayload.Error())
	}

	err = i.usecase.Task.Update(c.Request().Context(), model.Task{
		ID:      taskIDFormat,
		UserID:  ss.UserID,
		Summary: input.Summary,
	})

	if err != nil {
		return shared.NewError(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
