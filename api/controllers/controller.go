package controllers

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/rafbgarcia/hangman/api"
    "github.com/rafbgarcia/hangman/api/helpers"
    "github.com/rafbgarcia/hangman/api/models"
)

type Controller struct {
    Context *api.Context
    RespondWith *helpers.RespondWith
    ginContext *gin.Context
}

func (this *Controller) Init(ginContext *gin.Context) {
    this.Context = &api.Context{}
    this.Context.Init(ginContext)

    this.RespondWith = &helpers.RespondWith{
        GinContext: ginContext,
    }

    this.ginContext = ginContext
}

func (this *Controller) ResourceUrl(path ...interface{}) string {
    paths := []interface{}{
        this.ginContext.Request.URL.Scheme,
        "://",
        this.ginContext.Request.Host,
    }

    paths = append(paths, path...)

    return fmt.Sprint(paths...)
}


func (this *Controller) Player() *models.Player {
    return models.Player{}.New(this.Context.AccessToken)
}