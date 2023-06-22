package task

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/netorissi/SwordTest/api/middleware"
	"github.com/netorissi/SwordTest/shared"
)

// Task godoc
// @Summary      Complete task
// @Description  Only owner can complete task
// @Tags task
// @Param id path string true "Task ID"
// @Produce  json
// @Success 204
// @Failure 400 {object} shared.Response{error=shared.Error}
// @Failure 401 {object} shared.Response{error=shared.Error}
// @Failure 500 {object} shared.Response{error=shared.Error}
// @Security ApiKeyAuth
// @Router /v1/tasks/{id}/complete [patch]
func (i *impl) complete(c echo.Context) error {
	taskID := c.Param("id")
	if len(taskID) == 0 {
		return shared.NewError(http.StatusBadRequest, shared.ErrInvalidPayload.Error())
	}

	ss, err := middleware.GetSessionContext(c)
	if err != nil {
		return shared.NewError(http.StatusBadRequest, err.Error())
	}

	taskIDFormat, err := strconv.Atoi(taskID)
	if err != nil {
		return shared.NewError(http.StatusBadRequest, shared.ErrInvalidPayload.Error())
	}

	err = i.usecase.Task.Complete(c.Request().Context(), taskIDFormat, ss.UserID)
	if err != nil {
		return shared.NewError(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
