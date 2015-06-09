package app

import (
    "net/http"

	"github.com/rafbgarcia/hangman/api/config"
)

func init() {
	http.Handle("/", config.Router())
}
