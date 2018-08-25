package game

import "github.com/kataras/iris/websocket"

type Player struct {

	Id 			string                  `json:"id"`
	Name 		string                  `json:"name"`
	Owner		bool                    `json:"owner"`
	Connection 	websocket.Connection

}