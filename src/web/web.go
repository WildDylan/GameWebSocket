package web

import (
	"game"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	. "scheduler"
	. "web/request"
)

var App *iris.Application = nil

func StartWebServer()  {
	App = iris.New()
	App.Logger().SetLevel("debug")

	App.Use(recover.New())
	App.Use(logger.New())

	App.Post("/create", func(context iris.Context) {
		var roomRequest RoomRequest
		if e := context.ReadJSON(&roomRequest); e != nil {
			// Handler error
			return
		}

		// Create a game room
		room := &game.Room{
			Play:       &roomRequest.Play,
			Players:    []*game.Player{},
			Owner:      &roomRequest.Owner,

			Id:         roomRequest.Id,
			Lock:       roomRequest.Lock,
			Password:   roomRequest.Password,
		}

		// Add to global
		game.Rooms[roomRequest.Id] = room

		// Submit a Job to scheduler
		SubmitJob(func(gr *game.Room) {gr.Execute()}, room)
	})

	go App.Run(iris.Addr(":8181"), iris.WithoutServerError(iris.ErrServerClosed), iris.WithOptimizations)
}