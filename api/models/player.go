package models

import (
    "fmt"
    "crypto/rand"
    "appengine"
    "appengine/datastore"
    "github.com/rafbgarcia/hangman/api"
)

const (
    kindPlayers = "Players"
)

type Player struct {
    AccessToken string `json:"access_token"`
}

func (this Player) New(accessToken string) *Player {
    return &Player{
        AccessToken: accessToken,
    }
}

func (this Player) Create(c appengine.Context) (*Player, api.Error) {
    player := this.New(randToken())

    _, err := datastore.Put(c, player.Key(c), player)

    return player, api.DevError(err)
}

func (this Player) Find(c appengine.Context, accessToken string) (*Player, api.Error) {
    player := this.New(accessToken)

    err := datastore.Get(c, player.Key(c), player)

    return player, api.DevError(err)
}

func (this *Player) Key(c appengine.Context) *datastore.Key {
    return datastore.NewKey(c, kindPlayers, this.AccessToken, 0, nil)
}

func randToken() string {
    b := make([]byte, 32);
    rand.Read(b)

    return fmt.Sprintf("%x", string(b))
}
