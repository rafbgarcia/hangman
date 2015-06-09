package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/rafbgarcia/hangman/api"
    "github.com/rafbgarcia/hangman/api/models"
)

const (
    baseURL = "/games/"
)

type GamesController struct {
    *Controller
}

// Actions

func (this GamesController) Index(ginContext *gin.Context) {
    games, err := models.Game{}.FindAll(this.Context.AeContext, this.Player(), func(g *models.Game) {
        g.Url = this.ResourceUrl(baseURL, g.Id)
    })

    if err.Any() {
        this.RespondWith.UnexpectedError(err)
        return
    }

    this.RespondWith.JSON(http.StatusOK, games)
}

func (this GamesController) Show(ginContext *gin.Context) {
    game, err := models.Game{}.Find(this.Context.AeContext, ginContext.Params.ByName("id"), this.Player())

    game.Url = this.ResourceUrl(baseURL, game.Id)

    if err.Any() {
        this.RespondWith.UnexpectedError(err)
        return
    }

    this.RespondWith.JSON(http.StatusOK, game)
}

func (this GamesController) Create(ginContext *gin.Context) {
    game, err := models.Game{}.Create(this.Context.AeContext, this.Player())

    game.Url = this.ResourceUrl(baseURL, game.Id)

    if err.Any() {
        this.RespondWith.UnexpectedError(err)
        return
    }

    ginContext.Header("location", game.Url)
    ginContext.JSON(http.StatusCreated, game)
}

func (this GamesController) Guess(ginContext *gin.Context) {
    err := models.Game{}.Guess(this.Context.AeContext, this.Player(), api.M{
        "id": ginContext.Param("id"),
        "guess": this.Context.Body("char"),
    })

    if err.Any() {
        this.RespondWith.UnexpectedError(err)
        ginContext.AbortWithStatus(http.StatusNotModified)
        return
    }

    ginContext.AbortWithStatus(http.StatusNoContent)
}
