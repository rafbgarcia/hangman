package controllers_test

import (
    "testing"
    "encoding/json"
    "net/http/httptest"
    "appengine/aetest"
    "github.com/rafbgarcia/hangman/api/config"
    "net/http"
    "github.com/drborges/goexpect"
    "strings"
)

type M map[string]interface{}

type Response struct {
    code int
    body M
}

func TestPlayersController(t *testing.T) {
    expect := goexpect.New(t)

    res := post("/players")

    expect(res.code).ToBe(http.StatusCreated)
    expect(len(res.body["access_token"].(string)) > 0).ToBe(true)
}


// Extract this

func post(url string) Response {
    aeinstance, _ := aetest.NewInstance(nil)
    defer aeinstance.Close()

    req, _ := aeinstance.NewRequest("POST", url, strings.NewReader(""))
    req.Header.Add("Content-type", "application/json")

    return send(req)
}

func send(req *http.Request) Response {
    record := httptest.NewRecorder()
    config.Router().ServeHTTP(record, req)

    responseBody := M{}
    json.Unmarshal(record.Body.Bytes(), &responseBody)

    log(record.Body.String())

    return Response{
        code: record.Code,
        body: responseBody,
    }
}

func log(something interface{}) {
    c, _ := aetest.NewContext(nil)
    defer c.Close()

    c.Infof(">>> %+v", something)
}
