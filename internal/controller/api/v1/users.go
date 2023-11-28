package v1

import (
	"encoding/json"
	"github.com/good1hare/GolangTemplate/internal/usecase"
	"github.com/good1hare/GolangTemplate/pkg/logger"
	"github.com/labstack/echo/v4"
	"strconv"
)

type userRoutes struct {
	u usecase.User
	l logger.Interface
}

func newUsersRoutes(g *echo.Group, l logger.Interface, u usecase.User) {
	{
		r := &userRoutes{u: u, l: l}

		g := g.Group("/users")
		{
			g.POST("/", r.createUser)
			g.GET("/", r.getUsers)
			g.GET("/:id", r.getUser)
			g.PUT("/:id", r.updateUser)
			g.DELETE("/:id", r.deleteUser)
		}
	}

}

func (r *userRoutes) getUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(404, "user not found")
	}
	user, err := r.u.GetUser(id)
	if err != nil {
		return c.String(404, "user not found")
	}

	response, err := json.Marshal(user)
	if err != nil {
		return c.String(500, "Error converting to JSON")
	}

	return c.JSON(200, string(response))
}

func (r *userRoutes) getUsers(c echo.Context) error {
	return c.String(200, "get users")
}

func (r *userRoutes) createUser(c echo.Context) error {
	return c.String(200, "create user")
}

func (r *userRoutes) updateUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(200, id+"update user")
}

func (r *userRoutes) deleteUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(200, id+"delete user")
}
