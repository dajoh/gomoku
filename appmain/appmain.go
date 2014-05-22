package appmain

import (
	"appengine"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

var mainTpl = template.Must(template.ParseFiles("templates/main.tpl"))
var gameTpl = template.Must(template.ParseFiles("templates/game.tpl"))
var errorTpl = template.Must(template.ParseFiles("templates/error.tpl"))

func init() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/game", gameHandler)
	http.HandleFunc("/place", placeHandler)
	http.HandleFunc("/new_game", newGameHandler)
	http.HandleFunc("/_ah/channel/connected/", connectHandler)
	http.HandleFunc("/_ah/channel/disconnected/", disconnectHandler)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	mainTpl.Execute(w, nil)
}

func gameHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	token, err := JoinGame(ctx, r.FormValue("g"))

	if err != nil {
		errorTpl.Execute(w, err.Error())
	} else {
		gameTpl.Execute(w, token)
	}
}

func placeHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	game := r.FormValue("g")
	token := r.FormValue("p")
	x, err1 := strconv.Atoi(r.FormValue("x"))
	y, err2 := strconv.Atoi(r.FormValue("y"))

	if err1 != nil || err2 != nil {
		errorTpl.Execute(w, "Invalid X/Y")
		return
	}

	err := PlaceStone(ctx, game, token, x, y)
	if err != nil {
		errorTpl.Execute(w, err.Error())
	}
}

func newGameHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	key, err := NewGame(ctx)

	if err != nil {
		errorTpl.Execute(w, err.Error())
	} else {
		http.Redirect(w, r, "/game?g="+key, 302)
	}
}

func connectHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	key := strings.Split(r.FormValue("from"), ":")[0]
	ConnectToGame(ctx, key)
}

func disconnectHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	key := strings.Split(r.FormValue("from"), ":")[0]
	playerId := strings.Split(r.FormValue("from"), ":")[1]
	DisconnectFromGame(ctx, key, playerId)
}
