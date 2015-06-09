package api

import (
    "strings"
    "io/ioutil"
    "encoding/json"
    "appengine"
    "github.com/gin-gonic/gin"
)

type M map[string]string

type Context struct {
    AeContext appengine.Context
    Param func(string) string
    postBody map[string]string
    AccessToken string
}

func (this *Context) Init(ginContext *gin.Context) {
    this.AeContext = appengine.NewContext(ginContext.Request)

    bytes, err := ioutil.ReadAll(ginContext.Request.Body)
    if err == nil {
        json.Unmarshal(bytes, &this.postBody)
    }

    this.Param = ginContext.Param

    this.AccessToken = extractAccessToken(ginContext)
}

func (this *Context) Body(key string) string {
    return this.postBody[key]
}

func extractAccessToken(ginContext *gin.Context) string {
    auth := ginContext.Request.Header.Get("Authorization")
    return strings.TrimPrefix(auth, "Basic ")
}
