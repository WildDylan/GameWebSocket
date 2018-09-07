package socket

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/websocket"
	"github.com/weibocom/motan-go"
	"log"
	"net/http"
	"sync"
)

// 全部 websocket 实例
var Socket *websocket.Server = nil
// 全部 connections 映射
var PeersConnection = make(map[string]websocket.Connection)
// 锁
var mutex = new(sync.Mutex)

func Start()  {
	application := iris.New()
	// 使用自动恢复中间件
	application.Use(recover.New())
	// 设置日志级别为 debug
	application.Logger().SetLevel("debug")
	// 初始化全局 socket 实例
	Socket = websocket.New(websocket.Config{
		// 跨域检查
		CheckOrigin: checkOrigin,
		ReadBufferSize: 1024,
		WriteBufferSize: 1024,
	})
	// connection Handler
	Socket.OnConnection(onConnect)
	// endpoint
	application.Get("/echo", Socket.Handler())
	// start rpc
	StartRPC()
	// start socket server at 8282
	application.Run(iris.Addr(":8282"), iris.WithoutServerError(iris.ErrServerClosed), iris.WithOptimizations)
}

// cross origin check
func checkOrigin(r *http.Request) bool {
	return true
}

// connection handler
func onConnect(connection websocket.Connection) {
	// validated toke param
	token := connection.Context().URLParam("token")
	// which room want to join
	roomNumber := connection.Context().URLParam("roomNumber")
	if token == "" || roomNumber == "" {
		connection.Disconnect()
		return
	}

	log.Println(connection.ID() + " connected socket")

	// TODO: 用户上线，RPC -> token 上线

	// add connection to container
	mutex.Lock()
	PeersConnection[token] = connection
	// 绑定用户数据
	connection.SetValue("token", token)
	connection.SetValue("room", roomNumber)
	// 加入房间
	connection.Join(roomNumber)
	mutex.Unlock()

	connection.OnMessage(func(bytes []byte) { onMessage(bytes, connection) })
	connection.OnDisconnect(func() { onDisconnect(connection) })
}

func onMessage(bytes []byte, connection websocket.Connection) {
	// TODO: 收到 socket 消息， RPC -> token 收到消息
}

func onDisconnect(connection websocket.Connection) {
	// 取出绑定的用户数据
	token := connection.GetValueString("token")
	if token != "" {
		// TODO: 找到 token，RCP -> token 下线

		mutex.Lock()
		delete(PeersConnection, token)
		mutex.Unlock()
	}
}

func StartRPC()  {
	msContext := motan.GetMotanServerContext("src/socket/server.yaml")
	msContext.RegisterService(&RPCService{}, "socketService")
	msContext.Start(nil)
	msContext.ServicesAvailable()
}

type RPCService struct {

}

func (rpc *RPCService) BroadcastToRoom(room string, message string) bool {
	println(room, message)
	return true
}

