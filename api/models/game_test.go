package models_test

import (
    "testing"
    "github.com/drborges/goexpect"
    "github.com/rafbgarcia/hangman/api/models"
    "strings"
)


func TestNewGame(t *testing.T) {
    expect := goexpect.New(t)

    game := models.Game{}.New()

    expect(len(game.Word)).ToNotBe(0)
    expect(game.LettersRemaining).ToBe(game.Word)
    expect(game.Status).ToBe(models.GameStatus.Busy)
    expect(game.TriesLeft).ToBe(models.GameDefaultTriesLeft)
}

func TestGuessALetter(t *testing.T) {
    expect := goexpect.New(t)

    game := models.Game{}.New()


    // Make a correct guess

    existingLetter := string(game.Word[0])
    game.GuessALetter(existingLetter)

    expect(game.TriesLeft).ToBe(11)
    expect(strings.Contains(game.LettersRemaining, existingLetter)).ToBe(false)


    // Incorrect guess

    game.GuessALetter(existingLetter)
    expect(game.TriesLeft).ToBe(10)


    // Failed game

    loopUntil := game.TriesLeft
    for i := 0; i < loopUntil; i++ {
        game.GuessALetter(existingLetter)
    }
    expect(game.TriesLeft).ToBe(0)
    expect(game.Status).ToBe(models.GameStatus.Failed)


    // Success game

    game = models.Game{}.New()

    for i, _ := range game.Word {
        game.GuessALetter(string(game.Word[i]))
    }
    expect(game.Status).ToBe(models.GameStatus.Success)
}

func TestIsOver(t *testing.T) {
    expect := goexpect.New(t)

    game := &models.Game{
        Status: models.GameStatus.Busy,
    }
    expect(game.IsOver()).ToBe(false)

    game.Status = models.GameStatus.Failed
    expect(game.IsOver()).ToBe(true)

    game.Status = models.GameStatus.Success
    expect(game.IsOver()).ToBe(true)
}

func TestValidGuess(t *testing.T) {
    expect := goexpect.New(t)

    game := models.Game{}.New()


    // Invalid guess

    expect(game.ValidGuess("1")).ToBe(false)
    expect(game.ValidGuess("as")).ToBe(false)
    expect(game.ValidGuess("[")).ToBe(false)
    expect(game.ValidGuess("A")).ToBe(false)


    // Valid guess

    expect(game.ValidGuess("a")).ToBe(true)
    expect(game.ValidGuess("z")).ToBe(true)
}