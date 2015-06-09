package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/rafbgarcia/hangman/api/models"
    "github.com/rafbgarcia/hangman/api"
)

type PlayersController struct {
    *Controller
}

// Actions

func (this PlayersController) Create(ginContext *gin.Context) {
    player, err := models.Player{}.Create(this.Context.AeContext)

    if err.Any() {
        this.RespondWith.UnexpectedError(err)
        return
    }

    this.RespondWith.JSON(http.StatusCreated, player)
}


// Middlewares

func (this PlayersController) Authenticate(ginContext *gin.Context) {
    _, err := models.Player{}.Find(this.Context.AeContext, this.Context.AccessToken)

    if err.Any() {
        this.RespondWith.Error(http.StatusForbidden, api.UserError("Invalid Access Token"))
        ginContext.Abort()
    }
}