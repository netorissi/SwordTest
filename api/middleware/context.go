package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/netorissi/SwordTest/model"
	"github.com/netorissi/SwordTest/shared"
)

var contextSessionKey = "ctx_session"

func GetSessionContext(c echo.Context) (*model.Session, error) {
	sessionCtx := c.Get(contextSessionKey)

	sess, ok := sessionCtx.(*model.Session)
	if !ok {
		return nil, shared.NewError(http.StatusUnauthorized, "access_token is required.")
	}

	return sess, nil
}
