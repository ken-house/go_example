package controller

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/go_example/internal/utils/tools"

	"github.com/go_example/internal/meta"

	"github.com/gorilla/websocket"

	"github.com/gin-gonic/gin"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		clientIp := tools.GetClientIp(r)
		socketWhiteIpList := meta.GlobalConfig.Common.SocketWhiteIpList
		if len(socketWhiteIpList) > 0 {
			for _, ip := range socketWhiteIpList {
				pattern := strings.ReplaceAll(strings.ReplaceAll(ip, ".", "\\."), "*", ".*")
				match, _ := regexp.MatchString(pattern, clientIp)
				if match {
					return true
				}
			}
			return false
		}
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
	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatalf("upGrader,err is %+v", err)
		return
	}
	defer conn.Close()

	// 处理socket请求
	for {
		// 从conn中读取数据
		//var reqData map[string]string
		//err = conn.ReadJSON(&reqData)
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
		//// 写入conn数据
		//conn.WriteJSON(resData.Data)

		messageType, message, err := conn.ReadMessage()
		if err != nil {
			break
		}

		resData := "recv data:" + string(message)
		fmt.Println(resData)
		conn.WriteMessage(messageType, []byte(resData))

		// 定义一个chan
		heartChan := make(chan byte)
		go heartBeating(message[:1], heartChan)
		go heartHandler(conn, heartChan, 30)
	}
}

// 如果有消息则写入通道
func heartBeating(msg []byte, heartChan chan byte) {
	for _, v := range msg {
		heartChan <- v
	}
}

// 保活
func heartHandler(conn *websocket.Conn, heartChan chan byte, timeout int) {
	select {
	case <-heartChan:
		conn.SetReadDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
		conn.SetWriteDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
	}
}
