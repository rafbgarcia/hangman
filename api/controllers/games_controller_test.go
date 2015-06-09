package controllers_test

import (
    "testing"
    "appengine"
    "github.com/rafbgarcia/hangman/api/models"
    "fmt"
    "net/http"
    "github.com/drborges/goexpect"
    "net/url"
)

func TestCreateGame(t *testing.T) {
    expect := goexpect.New(t)

    res := makeRequestFromFn(POST, M{}, func(c appengine.Context, req *http.Request) {
        player, _ := models.Player{}.Create(c)

        req.Header.Add("Authorization", fmt.Sprint("Basic ", player.AccessToken))

        url, _ := url.Parse(fmt.Sprint("/games"))
        req.URL = url
    })

    expect(res.code).ToBe(http.StatusCreated)
}

func TestGuessALetter(t *testing.T) {
    expect := goexpect.New(t)

    res := makeRequestFromFn(PUT, M{"char":"a"}, func(c appengine.Context, req *http.Request) {
        player, _ := models.Player{}.Create(c)
        game, _ := models.Game{}.Create(c, player)

        req.Header.Add("Authorization", fmt.Sprint("Basic ", player.AccessToken))

        url, _ := url.Parse(fmt.Sprint("/games/", game.Id, "/guess"))
        req.URL = url
    })

    expect(res.code).ToBe(http.StatusNoContent)
}

func TestListAllGames(t *testing.T) {
    expect := goexpect.New(t)

    res := makeRequestFromFn(GET, M{}, func(c appengine.Context, req *http.Request) {
        player, _ := models.Player{}.Create(c)
        models.Game{}.Create(c, player)

        req.Header.Add("Authorization", fmt.Sprint("Basic ", player.AccessToken))

        url, _ := url.Parse("/games")
        req.URL = url
    })

    expect(res.code).ToBe(http.StatusOK)
}

func TestListAGame(t *testing.T) {
    expect := goexpect.New(t)

    res := makeRequestFromFn(GET, M{}, func(c appengine.Context, req *http.Request) {
        player, _ := models.Player{}.Create(c)
        game, _ := models.Game{}.Create(c, player)

        req.Header.Add("Authorization", fmt.Sprint("Basic ", player.AccessToken))

        url, _ := url.Parse(fmt.Sprint("/games/", game.Id))
        req.URL = url
    })

    expect(res.code).ToBe(http.StatusOK)
}
