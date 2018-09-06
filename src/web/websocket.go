package web

import (
	"game"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/websocket"
	"log"
	"net/http"
)

var SocketApplication *iris.Application = nil
var Socket *websocket.Server = nil

var ConnectionPeers = make(map[websocket.Connection]string)

func StartWebSocketServer()  {

	SocketApplication = iris.New()

	SocketApplication.Logger().SetLevel("debug")

	SocketApplication.Use(recover.New())
	SocketApplication.Use(logger.New())

	Socket = websocket.New(websocket.Config{
		CheckOrigin: checkOrigin,
	})

	Socket.OnConnection(onConnect)

	SocketApplication.Get("/echo", Socket.Handler())

	SocketApplication.Run(iris.Addr(":8282"), iris.WithoutServerError(iris.ErrServerClosed), iris.WithOptimizations)
}

func checkOrigin(r *http.Request) bool {
	return true
}

func onConnect(connection websocket.Connection) {
	// Validated toke param
	token := connection.Context().URLParam("token")
	roomNumber := connection.Context().URLParam("roomNumber")
	if token == "" || roomNumber == "" {
		connection.Disconnect()
		return
	}

	log.Println(connection.ID() + " connected socket")

	// Global connections container
	ConnectionPeers[connection] = connection.Context().URLParam("token")
	// Get Room
	room, ok := game.Rooms[roomNumber]
	if ok {
		// Add player to room
		player := &game.Player{
			Id:             token,
			Name:           "Owner",
			Connection:     connection,
			Owner:          true,
		}
		room.AddPlayer(player)
		// Message listener
		connection.OnMessage(room.OnMessage)
		// Disconnect listener
		connection.OnDisconnect(func() {
			// Remove connection
			delete(ConnectionPeers, connection)
			log.Println(connection.ID(), "dis connected")
		})
	} else {
		connection.Disconnect()
	}
}