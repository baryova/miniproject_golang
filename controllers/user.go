package controllers

import (
	"fmt"
	"log"
	"miniproject_golang/lib/database"
	"miniproject_golang/middleware"
	"miniproject_golang/models"
	"net/http"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

func GetUsersController(c echo.Context) error {
	users, err := database.GetUser()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success get all user",
		Data:    users,
	})
}

func CreateUserController(c echo.Context) error {

	user := models.User{}
	c.Bind(&user)

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}

	user.Password = string(hashPassword)

	user, err = database.CreateUser(user)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success create user",
		Data:    user,
	})
}

func LoginUserController(c echo.Context) error {
	userInput := models.User{}
	c.Bind(&userInput)

	user, err := database.LoginUser(userInput)

	log.Println(user.Password)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "user not found or wrong password",
		})
	}

	user.Token, err = middleware.CreateToken(int(user.ID), user.Username)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "fail login",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success login",
		Data:    user,
	})
}

func UpdateUserPassword(c echo.Context) error {
	var newPassword string

	userInput := models.User{}

	userId, err := middleware.GetUserIdByToken(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	user, err := database.GetUsernameById(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	fmt.Println("aaa : ", user)

	c.Bind(&newPassword)

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}

	userInput.Username = user.Username
	userInput.Password = string(hashPassword)

	newUser, err := database.UpdateUserPassword(userInput)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success change password",
		Data:    newUser,
	})
}

func DeleteUser(c echo.Context) error {

	userId, err := middleware.GetUserIdByToken(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	user, err := database.GetUsernameById(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = database.DeleteUser(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success delete user",
		Data:    nil,
	})
}
