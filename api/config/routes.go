package config

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/rafbgarcia/hangman/api/controllers"
)

func Router() http.Handler {
    router := gin.Default()

    controller := &controllers.Controller{}

    router.Use(controller.Init)


    playersController := controllers.PlayersController{controller}
    players := router.Group("/players")
    {
        players.POST("", playersController.Create)
    }


    gamesController := controllers.GamesController{controller}
    games := router.Group("/games", playersController.Authenticate)
    {
        games.GET("", gamesController.Index)
        games.POST("", gamesController.Create)
        games.GET("/:id", gamesController.Show)
        games.PUT("/:id/guess", gamesController.Guess)
    }


    return router
}
