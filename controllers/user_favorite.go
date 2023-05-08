package controllers

import (
	"miniproject_golang/lib/database"
	"miniproject_golang/middleware"
	"miniproject_golang/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// Get all fav movies by logged in user
func GetUserFavorite(c echo.Context) error {
	userId, err := middleware.GetUserIdByToken(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	userFav, err := database.GetUserFavorite(userId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success get user favorite lists",
		Data:    userFav,
	})
}

func DoFavorite(c echo.Context) error {
	userFav := models.UserFavorite{}

	userId, err := middleware.GetUserIdByToken(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	//get film_id by param
	filmId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	//bind user_id and film_id
	userFav.UserId = userId
	userFav.FilmId = filmId

	//get the fav status by logged in user and the filmId
	faved, err := database.GetSameUserFavorite(userFav)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if faved == true { //if already fav , then undo the fav
		userFav, err = database.UndoFavorite(userFav)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, models.Response{
			Message: "success un-favorite movie",
			Data:    userFav,
		})

	} else {
		userFav, err = database.DoFavorite(userFav)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, models.Response{
			Message: "success favorite movie",
			Data:    userFav,
		})
	}

}
