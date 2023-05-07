package routes

import (
	"miniproject_golang/controllers"
	"miniproject_golang/middleware"

	"github.com/labstack/echo"
	mid "github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	middleware.LogMiddleware(e)
	e.Pre(mid.RemoveTrailingSlash())

	//User Routes
	users := e.Group("/users")
	users.GET("", controllers.GetUsersController)
	users.POST("", controllers.CreateUserController)
	users.POST("/login", controllers.LoginUserController)

	return e
}
