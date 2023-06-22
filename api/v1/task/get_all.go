package task

import (
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/netorissi/SwordTest/model"
	"github.com/netorissi/SwordTest/shared"
)

// Task godoc
// @Summary      Get all tasks
// @Description  Only manager can list all tasks
// @Tags task
// @Produce  json
// @Success 200 {object} shared.Response{data=[]model.Task}
// @Failure 400 {object} shared.Response{error=shared.Error}
// @Failure 401 {object} shared.Response{error=shared.Error}
// @Failure 500 {object} shared.Response{error=shared.Error}
// @Security ApiKeyAuth
// @Router /v1/tasks [get]
func (i *impl) getAll(c echo.Context) error {
	tasks, err := i.usecase.Task.Get(c.Request().Context())
	if err != nil {
		return shared.NewError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, shared.Response{Data: tasks})
}
