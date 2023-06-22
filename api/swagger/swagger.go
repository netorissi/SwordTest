package swagger

import (
	"net/http"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/netorissi/SwordTest/config"
	docs "github.com/netorissi/SwordTest/docs"
	"github.com/netorissi/SwordTest/shared"
)

func New(groupApi *echo.Group) {
	docs.SwaggerInfo.Title = config.Global.Swagger.Title
	docs.SwaggerInfo.Description = config.Global.Swagger.Description
	docs.SwaggerInfo.Version = config.Global.Swagger.Version
	docs.SwaggerInfo.Host = config.Global.Swagger.Host + config.Global.Server.Port
	docs.SwaggerInfo.BasePath = config.Global.Server.BasePath
	docs.SwaggerInfo.Schemes = []string{"http"}

	groupApi.GET("", func(c echo.Context) error {
		return c.Redirect(http.StatusFound, config.Global.Server.BasePath+"/swagger/index.html")
	})

	groupApi.GET("/*", echoSwagger.WrapHandler)

	shared.Logger.Info("Swagger started")
}
