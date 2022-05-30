package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/gin-gonic/gin"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// todo 获取客户端IP
		//clientIp := r.Header.Get("X-Forward-For")
		//fmt.Println("clientIp:", clientIp)
		//if !tools.IsContain(meta.SocketWhiteIpList, clientIp) {
		//	return false
		//}
		return true
	},
}

type SocketController interface {
	SocketServer(c *gin.Context)
}

type socketController struct {
}

func NewSocketController() SocketController {
	return &socketController{}
}

func (ctr *socketController) SocketServer(c *gin.Context) {
	// 升级为WebSocket
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatalf("upGrader,err is %+v", err)
		return
	}
	defer ws.Close()

	// 处理socket请求
	for {
		// 从ws中读取数据
		//var reqData map[string]string
		//err = ws.ReadJSON(&reqData)
		//if err != nil {
		//	break
		//}
		//
		//message := reqData["name"] + "_" + reqData["age"]

		//_, resData := negotiate.JSON(http.StatusOK, gin.H{
		//	"data": gin.H{
		//		"message": "recv msg2：" + string(message),
		//	},
		//})
		//// 写入ws数据
		//ws.WriteJSON(resData.Data)

		messageType, message, err := ws.ReadMessage()
		if err != nil {
			break
		}

		resData := "recv data:" + string(message)
		fmt.Println(resData)
		ws.WriteMessage(messageType, []byte(resData))
	}
}
