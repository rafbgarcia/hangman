package models_test

import (
    "testing"
    "github.com/drborges/goexpect"
    "github.com/rafbgarcia/hangman/api/models"
)


func TestNewPlayer(t *testing.T) {
    expect := goexpect.New(t)

    accessToken := "123"
    player := models.Player{}.New(accessToken)

    expect(player.AccessToken).ToBe(accessToken)
}
