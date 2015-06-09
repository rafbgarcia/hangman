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
    "appengine"
    "net/url"
)

const (
    notificationBody = "Test notification"
    POST = "POST"
    PUT = "PUT"
    GET = "GET"
)

type M map[string]interface{}

type Response struct {
    code int
    body M
}

func TestPlayersController(t *testing.T) {
    expect := goexpect.New(t)

    res := post("/players", M{})

    expect(res.code).ToBe(http.StatusCreated)
    expect(len(res.body["access_token"].(string)) > 0).ToBe(true)
}


// Extract this

func post(rawUrl string, body M) Response {
    return makeRequestFromFn(POST, body, func(c appengine.Context, req *http.Request)  {
        url, _ := url.Parse(rawUrl)
        req.URL = url
    })
}

func makeRequestFromFn(method string, body M,  fn func(appengine.Context, *http.Request)) Response {
    aeinstance, _ := aetest.NewInstance(nil)
    defer aeinstance.Close()

    bytes, _ := json.Marshal(body)
    reader := strings.NewReader(string(bytes))

    req, _ := aeinstance.NewRequest(method, "/", reader)
    req.Header.Add("Content-type", "application/json")

    c := appengine.NewContext(req)

    fn(c, req)

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
