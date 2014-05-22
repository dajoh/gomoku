package appmain

import (
	"appengine"
	"appengine/channel"
	"appengine/datastore"
	"errors"
	"fmt"
)

const BoardSize = 8

type Game struct {
	PlayersReady    int
	ChannelsCreated int
	Turn            int
	Board           []int
	PlayerTokens    []string
}

func NewGame(ctx appengine.Context) (string, error) {
	game := new(Game)
	game.Board = make([]int, BoardSize*BoardSize)
	game.PlayerTokens = make([]string, 2)

	for i := 0; i < BoardSize*BoardSize; i++ {
		game.Board[i] = -1
	}

	key, err := datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "Game", nil), game)
	if err != nil {
		return "", err
	}

	return key.Encode(), err
}

func JoinGame(ctx appengine.Context, key string) (string, error) {
	game := new(Game)
	token := ""

	err := datastore.RunInTransaction(ctx, func(ctx appengine.Context) error {
		decKey, err := datastore.DecodeKey(key)
		if err != nil {
			return errors.New("Invalid game key.")
		}

		err = datastore.Get(ctx, decKey, game)
		if err != nil {
			return errors.New("Session expired.")
		}

		token, err = channel.Create(ctx, fmt.Sprintf("%v:%v", key, game.ChannelsCreated))
		if err != nil {
			return err
		}

		if game.ChannelsCreated != 2 {
			game.PlayerTokens[game.ChannelsCreated] = token
			game.ChannelsCreated++
		} else {
			return errors.New("All players already ingame.")
		}

		_, err = datastore.Put(ctx, decKey, game)
		return err
	}, nil)

	if err != nil {
		return "", err
	} else {
		return token, nil
	}
}

func PlaceStone(ctx appengine.Context, key, token string, x, y int) error {
	game := new(Game)
	playerId := 0

	err := datastore.RunInTransaction(ctx, func(ctx appengine.Context) error {
		decKey, err := datastore.DecodeKey(key)
		if err != nil {
			return errors.New("Invalid game key.")
		}

		err = datastore.Get(ctx, decKey, game)
		if err != nil {
			return errors.New("Session expired.")
		}

		if game.PlayerTokens[0] == token {
			playerId = 0
		} else if game.PlayerTokens[1] == token {
			playerId = 1
		} else {
			return errors.New("Invalid player token.")
		}

		if game.Turn != playerId {
			return errors.New("Not your turn.")
		}

		if game.Board[y*BoardSize+x] != -1 {
			return errors.New("Tile occupied.")
		} else {
			game.Turn = (playerId + 1) % 2
			game.Board[y*BoardSize+x] = playerId
		}

		_, err = datastore.Put(ctx, decKey, game)
		return err
	}, nil)

	if err != nil {
		ctx.Warningf("Place error: %v\n", err.Error())
		return err
	}

	channel.Send(ctx, key+":0", fmt.Sprintf("move:%v:%v:%v", playerId, x, y))
	channel.Send(ctx, key+":1", fmt.Sprintf("move:%v:%v:%v", playerId, x, y))

	return nil
}

func ConnectToGame(ctx appengine.Context, key string) {
	game := new(Game)

	err := datastore.RunInTransaction(ctx, func(ctx appengine.Context) error {
		decKey, err := datastore.DecodeKey(key)
		if err != nil {
			return err
		}

		err = datastore.Get(ctx, decKey, game)
		if err != nil {
			return err
		} else {
			game.PlayersReady++
		}

		_, err = datastore.Put(ctx, decKey, game)
		return err
	}, nil)

	if err != nil {
		panic(err)
	}

	if game.PlayersReady == 2 {
		channel.Send(ctx, fmt.Sprintf("%v:0", key), "begin")
		channel.Send(ctx, fmt.Sprintf("%v:1", key), "wait")
	}
}

func DisconnectFromGame(ctx appengine.Context, key, playerId string) {
	datastore.RunInTransaction(ctx, func(ctx appengine.Context) error {
		decKey, err := datastore.DecodeKey(key)
		if err == nil {
			return err
		}
		return datastore.Delete(ctx, decKey)
	}, nil)

	if playerId == "0" {
		channel.Send(ctx, fmt.Sprintf("%v:1", key), "partner_dc")
	} else {
		channel.Send(ctx, fmt.Sprintf("%v:0", key), "partner_dc")
	}
}
