package v1

import (
	"github.com/good1hare/GolangTemplate/internal/usecase"
	"github.com/good1hare/GolangTemplate/pkg/logger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(e *echo.Echo, l logger.Interface, u usecase.User) {
	g := e.Group("/v1")
	{
		newUsersRoutes(g, l, u)
	}
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "joe" && password == "secret" {
			return true, nil
		}
		return false, nil
	}))

}
