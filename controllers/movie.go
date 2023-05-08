package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"miniproject_golang/models"
	"net/http"

	"github.com/labstack/echo"
)

func GetNowPlayingMovies(c echo.Context) error {

	var bodyResults models.MovieResponse

	url := "https://api.themoviedb.org/3/movie/now_playing?language=en-US&page=1"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiIzYTAxZjc5YmRkY2M5MTU2N2FlMTljMTljZTE4N2U0YSIsInN1YiI6IjY0NTZhMzE2NmM4NDkyMDEyNGM0OWNhMSIsInNjb3BlcyI6WyJhcGlfcmVhZCJdLCJ2ZXJzaW9uIjoxfQ.ND5Av9DjGOeF3hKIYD1mGyIMVKLJwj3OTd7eNsRrZFg")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	err := json.Unmarshal(body, &bodyResults)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success get now playing movies",
		Data:    bodyResults,
	})
}

func SearchMovies(c echo.Context) error {
	var bodyResults models.MovieResponse

	url := "https://api.themoviedb.org/3/search/movie"

	req, _ := http.NewRequest("GET", url, nil)

	q := req.URL.Query()
	query := c.QueryParam("query")

	q.Add("query", query)
	req.URL.RawQuery = q.Encode()

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiIzYTAxZjc5YmRkY2M5MTU2N2FlMTljMTljZTE4N2U0YSIsInN1YiI6IjY0NTZhMzE2NmM4NDkyMDEyNGM0OWNhMSIsInNjb3BlcyI6WyJhcGlfcmVhZCJdLCJ2ZXJzaW9uIjoxfQ.ND5Av9DjGOeF3hKIYD1mGyIMVKLJwj3OTd7eNsRrZFg")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	err := json.Unmarshal(body, &bodyResults)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success search movies",
		Data:    bodyResults,
	})
}

func GetDetailMovie(c echo.Context) error {
	var bodyDetail models.MovieDetailResponse

	url := "https://api.themoviedb.org/3/movie"

	movieId := c.Param("movie_id") //eg. 758323 , 594767, 594767

	fmt.Println("aaaa : ", movieId)

	url += "/" + movieId
	fmt.Println("bbbb : ", url)
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiIzYTAxZjc5YmRkY2M5MTU2N2FlMTljMTljZTE4N2U0YSIsInN1YiI6IjY0NTZhMzE2NmM4NDkyMDEyNGM0OWNhMSIsInNjb3BlcyI6WyJhcGlfcmVhZCJdLCJ2ZXJzaW9uIjoxfQ.ND5Av9DjGOeF3hKIYD1mGyIMVKLJwj3OTd7eNsRrZFg")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	err := json.Unmarshal(body, &bodyDetail)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, models.Response{
		Message: "success get now playing movies",
		Data:    bodyDetail,
	})
}
