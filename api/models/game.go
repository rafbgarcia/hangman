package models

import (
    "strconv"
    "appengine/datastore"

    "github.com/Pallinder/go-randomdata"
    "github.com/rafbgarcia/hangman/api"
    "strings"
    "regexp"
    "appengine"
)

const (
    GameDefaultTriesLeft = 11
    kindGames = "Games"
)

var (
    GameStatus = struct {Success, Busy, Failed string}{"success", "busy", "failed"}
)

type Game struct {
    Url string `json:"url" datastore:"-"`
    Id int64 `json:"id" datastore:"-"`

    Word string `json:"word"`
    TriesLeft int `json:"tries_left"`
    Status string `json:"status"`
    LettersRemaining string `json:"letters_remaining"`
}

func (this Game) New() *Game {
    word := strings.ToLower(randomdata.FirstName(randomdata.RandomGender))

    return &Game{
        Status: GameStatus.Busy,
        TriesLeft: GameDefaultTriesLeft,
        Word: word,
        LettersRemaining: word,
    }
}

func (this Game) Create(c appengine.Context, player *Player) (*Game, api.Error) {
    game := Game{}.New()

    incompleteKey := datastore.NewIncompleteKey(c, kindGames, player.Key(c))

    key, err := datastore.Put(c, incompleteKey, game)

    game.Id = key.IntID()

    return game, api.DevError(err)
}

func (this Game) FindAll(c appengine.Context, player *Player, f func(*Game)) ([]*Game, api.Error) {
    games := []*Game{}

    keys, err := datastore.NewQuery(kindGames).Ancestor(player.Key(c)).GetAll(c, &games)

    for i, game := range games {
        game.Id = keys[i].IntID()
        if f != nil {
            f(game)
        }
    }

    return games, api.DevError(err)
}

func (this Game) Find(c appengine.Context, id string, player *Player) (*Game, api.Error) {
    game := &Game{}

    key := keyFromId(c, id, player)

    err := datastore.Get(c, key, game)

    game.Id = key.IntID()

    return game, api.DevError(err)
}

func (this Game) Guess(c appengine.Context, player *Player, p api.M) api.Error {
    game, err := this.Find(c, p["id"], player)

    if err.Any() {
        return err
    }

    if game.IsOver() {
        return api.UserError("Game ended wit status: " + this.Status)
    }

    if !game.ValidGuess(p["guess"]) {
        return api.UserError("Sorry, try guessing one lowercase letter per time.")
    }

    game.GuessALetter(p["guess"])

    _, error := datastore.Put(c, keyFromId(c, p["id"], player), game)

    return api.DevError(error)
}

func (this *Game) ValidGuess(guess string) bool {
    hasJustOneLetter, _ := regexp.MatchString("^[a-z]$", guess);
    return hasJustOneLetter
}

func (this *Game) GuessALetter(guess string) {
    if strings.Contains(this.LettersRemaining, guess) {
        this.LettersRemaining = strings.Replace(this.LettersRemaining, guess, "", -1)

        if len(this.LettersRemaining) == 0 {
            this.Status = GameStatus.Success
        }
    } else {
        this.TriesLeft--

        if this.TriesLeft == 0 {
            this.Status = GameStatus.Failed
        }
    }
}

func (this *Game) IsOver() bool {
    return this.Status != GameStatus.Busy
}

func keyFromId(c appengine.Context, id string, player *Player) *datastore.Key {
    val, _ := strconv.ParseInt(id, 10, 64)
    return datastore.NewKey(c, kindGames, "", val, player.Key(c))
}
