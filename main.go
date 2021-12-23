package main

import (
	"github.com/b4rsch/Kickerplatform/handler"
	"github.com/b4rsch/Kickerplatform/repository"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func setupRouter(repository repository.Repo) *gin.Engine {
	r := gin.Default()

	r.POST("/user/:username/:locationId", func(c *gin.Context) {
		username := c.Param("username")
		locationIdFromParam, err := strconv.ParseInt(c.Param("locationId"), 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, "{}")
		}
		locationId := int(locationIdFromParam)
		err = repository.CreateUser(username, locationId)
		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, `{"error":true,"message":"could not create user"`)
			return
		}
		c.JSON(http.StatusCreated, "{}")
	})

	r.POST("/match/new", func(c *gin.Context) {

		payload := handler.MatchDetails{}
		if err := c.BindJSON(&payload); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, `{"error":true,"message":"bad request"`)
			return
		}

	})

	return r
}

func main() {
	repo := repository.NewRepository()
	r := setupRouter(repo)
	r.Run()
}
