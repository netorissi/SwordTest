package task

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/netorissi/SwordTest/api/middleware"
	"github.com/netorissi/SwordTest/shared"
)

// Task godoc
// @Summary      Get my tasks
// @Description  Only tech can list my tasks
// @Tags task
// @Produce  json
// @Success 200 {object} shared.Response{data=[]model.Task}
// @Failure 400 {object} shared.Response{error=shared.Error}
// @Failure 401 {object} shared.Response{error=shared.Error}
// @Failure 500 {object} shared.Response{error=shared.Error}
// @Security ApiKeyAuth
// @Router /v1/tasks/me [get]
func (i *impl) me(c echo.Context) error {
	ss, err := middleware.GetSessionContext(c)
	if err != nil {
		return err
	}

	tasks, err := i.usecase.Task.GetByUserID(c.Request().Context(), ss.UserID)
	if err != nil {
		return shared.NewError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, tasks)
}
