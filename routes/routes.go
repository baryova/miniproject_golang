package routes

import (
	"miniproject_golang/constants"
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
	users.GET("", controllers.GetUsersController, mid.JWT([]byte(constants.JWT_SECRET)))
	users.POST("", controllers.CreateUserController)
	users.POST("/login", controllers.LoginUserController)

	//User_Fav Routes
	userFav := e.Group("/favorite")
	userFav.GET("/list", controllers.GetUserFavorite, mid.JWT([]byte(constants.JWT_SECRET)))
	userFav.POST("/:id", controllers.DoFavorite, mid.JWT([]byte(constants.JWT_SECRET))) //do and undo

	//Movie Routes
	movies := e.Group("/movies")
	movies.GET("/now-playing", controllers.GetNowPlayingMovies)
	movies.GET("/search", controllers.SearchMovies)
	movies.GET("/:movie_id", controllers.GetDetailMovie)

	return e
}
