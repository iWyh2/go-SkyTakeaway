package websocket

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	errs "go-SkyTakeaway/common/errors"
	"log"
	"net/http"
)

// 创建WebSocket升级器并初始化
var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// 允许跨域访问
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WSServer 创建websocketServer并初始化
var WSServer = Server{
	conns: make([]*websocket.Conn, 0),
}

// WSRouter 注册WebSocket路由
func WSRouter(r *gin.Engine) {
	r.GET("/ws/:id", websocketHandler)
}

// WebSocket处理器
func websocketHandler(ctx *gin.Context) {
	// 升级HTTP连接到WebSocket连接
	conn, err := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		panic(errs.WebSocketUpgradeError)
	}
	clientId := ctx.Param("id")
	log.Printf("client[%v]连接成功", clientId)
	WSServer.conns = append(WSServer.conns, conn)
}

// Server websocketServer
type Server struct {
	conns []*websocket.Conn
}

// SendToAllClients 发送消息给所有客户端
func (s Server) SendToAllClients(jsonMsg any) {
	for i := 0; i < len(s.conns); i++ {
		if err := s.conns[i].WriteJSON(jsonMsg); err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure) ||
				websocket.IsCloseError(err, websocket.CloseGoingAway) ||
				websocket.IsCloseError(err, websocket.CloseAbnormalClosure) ||
				websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) ||
				websocket.IsUnexpectedCloseError(err, websocket.CloseNormalClosure) ||
				websocket.IsUnexpectedCloseError(err, websocket.CloseAbnormalClosure) {
				// 移除无效conn
				s.conns = append(s.conns[:i], s.conns[i+1:]...)
				continue
			}
			panic(errs.WebSocketSendMsgError)
		}
	}
}
