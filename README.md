Humble implementation of hangman in go.


#### Get your access token

```bash
curl -X POST http://sup-hangman.appspot.com/players
```


#### Start a game

```bash
curl -X POST \
     -H "Authorization: Basic <ACCESS_TOKEN>" \
     http://sup-hangman.appspot.com/games
```


#### Resume of all games

```bash
curl -X GET \
     -H "Authorization: Basic <ACCESS_TOKEN>" \
     http://sup-hangman.appspot.com/games
```


#### Resume of a game

```bash
curl -X GET \
     -H "Authorization: Basic <ACCESS_TOKEN>" \
     http://sup-hangman.appspot.com/games/:id
```


#### Guess a letter

```bash
curl -X PUT \
     -H "Authorization: Basic <ACCESS_TOKEN>" \
     -H "Content-type: application/json" \
     -d '{"char":"n"}'` \
     http://sup-hangman.appspot.com/games/:id/guess
```
