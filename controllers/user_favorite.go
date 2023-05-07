package controllers

import (
	"fmt"
	"miniproject_golang/lib/database"
	"miniproject_golang/middleware"
	"miniproject_golang/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetUserFavorite(c echo.Context) error {
	userId, err := middleware.GetUserIdByToken(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	users, err := database.GetUserFavorite(userId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success get all user",
		Data:    users,
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

	faved, err := database.GetSameUserFavorite(userFav)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	fmt.Println("ajioewjfois: ", faved)

	if faved == true {
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
