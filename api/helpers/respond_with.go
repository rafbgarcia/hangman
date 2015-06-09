package helpers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/rafbgarcia/hangman/api"
)

type RespondWith struct {
    GinContext *gin.Context
}

func (this *RespondWith) JSON(httpStatus int, obj interface{}) {
    this.GinContext.JSON(httpStatus, obj)
}

func (this *RespondWith) UnexpectedError(err api.Error) {
    if err.Any() {
        this.Error(http.StatusInternalServerError, err)
        return
    }

    this.Error(http.StatusInternalServerError, api.UserError("An unexpected error happened."))
}

func (this *RespondWith) Error(httpStatus int, err api.Error) {
    response := gin.H{
        "error_code": httpStatus,
    }

    if err.Any() {
        if err.IsUserError() {
            response["message"] = err.Text()

        } else if gin.Mode() != gin.ReleaseMode {
            response["debug_message"] = err.Text()
            delete(response, "message")
        }
    }

    this.GinContext.JSON(httpStatus, response)
}
