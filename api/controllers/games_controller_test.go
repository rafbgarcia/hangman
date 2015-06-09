package controllers_test

import (
    "testing"
    "appengine"
    "appengine/aetest"
    "strings"
    "github.com/rafbgarcia/hangman/api/models"
    "fmt"
    "net/http"
    "github.com/drborges/goexpect"
    "encoding/json"
)

func TestCreateGame(t *testing.T) {
    expect := goexpect.New(t)

    aeinstance, _ := aetest.NewInstance(nil)
    defer aeinstance.Close()


    // Given
    req, _ := aeinstance.NewRequest("POST", "/games", strings.NewReader(""))
    req.Header.Add("Content-type", "application/json")

    c := appengine.NewContext(req)
    player, _ := models.Player{}.Create(c)

    req.Header.Add("Authorization", fmt.Sprint("Basic ", player.AccessToken))


    // When
    res := send(req)


    // Then
    expect(res.code).ToBe(http.StatusCreated)
}

func TestGuessALetter(t *testing.T) {
    expect := goexpect.New(t)

    aeinstance, _ := aetest.NewInstance(nil)
    defer aeinstance.Close()


    // Given
    bytes, _ := json.Marshal(M{
        "char": "a",
    })

    req, _ := aeinstance.NewRequest("POST", "/games/:id/guess", strings.NewReader(string(bytes)))
    req.Header.Add("Content-type", "application/json")

    c := appengine.NewContext(req)
    player, _ := models.Player{}.Create(c)
    game, _ := models.Game{}.Create(c, player)


    req, _ = aeinstance.NewRequest("PUT", fmt.Sprint("/games/", game.Id, "/guess"), strings.NewReader(string(bytes)))
    req.Header.Add("Authorization", fmt.Sprint("Basic ", player.AccessToken))


    // When
    res := send(req)


    // Then
    expect(res.code).ToBe(http.StatusNoContent)
}

func TestListAllGames(t *testing.T) {
    expect := goexpect.New(t)

    aeinstance, _ := aetest.NewInstance(nil)
    defer aeinstance.Close()


    // Given
    req, _ := aeinstance.NewRequest("GET", "/games", strings.NewReader(""))
    req.Header.Add("Content-type", "application/json")

    c := appengine.NewContext(req)
    player, _ := models.Player{}.Create(c)
    models.Game{}.Create(c, player)


    req, _ = aeinstance.NewRequest("GET", "/games", strings.NewReader(""))
    req.Header.Add("Authorization", fmt.Sprint("Basic ", player.AccessToken))


    // When
    res := send(req)


    // Then
    expect(res.code).ToBe(http.StatusOK)
}

func TestListAGame(t *testing.T) {
    expect := goexpect.New(t)

    aeinstance, _ := aetest.NewInstance(nil)
    defer aeinstance.Close()


    // Given
    req, _ := aeinstance.NewRequest("GET", "/games/:id", strings.NewReader(""))
    req.Header.Add("Content-type", "application/json")

    c := appengine.NewContext(req)
    player, _ := models.Player{}.Create(c)
    game, _ := models.Game{}.Create(c, player)


    req, _ = aeinstance.NewRequest("GET", fmt.Sprint("/games/", game.Id), strings.NewReader(""))
    req.Header.Add("Authorization", fmt.Sprint("Basic ", player.AccessToken))


    // When
    res := send(req)


    // Then
    expect(res.code).ToBe(http.StatusOK)
}
