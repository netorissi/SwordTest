package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/netorissi/SwordTest/model"
	"github.com/netorissi/SwordTest/shared"
	"github.com/netorissi/SwordTest/usecase"
)

var (
	ErrManagerRequired = errors.New("role manager is required")
	ErrTechRequired    = errors.New("role tech is required")
)

type Middleware interface {
	Tech(next echo.HandlerFunc) echo.HandlerFunc
	Manager(next echo.HandlerFunc) echo.HandlerFunc
}

type impl struct {
	usecase usecase.Usecase
}

func New(usecase usecase.Usecase) Middleware {
	return &impl{usecase}
}

func (i *impl) Tech(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		s, err := i.getAndSetSession(c)
		if err != nil {
			return err
		}

		user, err := i.usecase.User.GetByID(c.Request().Context(), s.UserID)
		if err != nil {
			return err
		}

		if user.Role != model.RoleTech {
			return shared.NewError(http.StatusUnauthorized, ErrTechRequired.Error())
		}

		return next(c)
	}
}

func (i *impl) Manager(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		s, err := i.getAndSetSession(c)
		if err != nil {
			return err
		}

		user, err := i.usecase.User.GetByID(c.Request().Context(), s.UserID)
		if err != nil {
			return err
		}

		if user.Role != model.RoleManager {
			return shared.NewError(http.StatusUnauthorized, ErrManagerRequired.Error())
		}

		return next(c)
	}
}

func (i *impl) getAndSetSession(c echo.Context) (*model.Session, error) {
	jwtToken := strings.Split(c.Request().Header.Get("Authorization"), "Bearer ")

	if len(jwtToken) != 2 {
		return nil, shared.NewError(http.StatusUnauthorized, "access_token is required.")
	}

	ss, err := i.usecase.Session.GetByToken(c.Request().Context(), jwtToken[1])
	if err != nil || ss.UserID == 0 {
		return nil, shared.NewError(http.StatusUnauthorized, "your session is invalid.")
	}

	c.Set(contextSessionKey, ss)

	return ss, nil
}
