package process

import (
	"chat/common"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

type Process struct {
	conn *websocket.Conn
}

var wu = &websocket.Upgrader{ReadBufferSize: 512,
	WriteBufferSize: 512, CheckOrigin: func(r *http.Request) bool { return true }}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("message get")
	ws, err := wu.Upgrade(w, r, nil)
	fmt.Println(ws.RemoteAddr().String())
	if err != nil {
		return
	}
	go reader(ws)
}

func reader(conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("conn err")
			return
		}
		var data common.Data
		json.Unmarshal(message, &data)
		fmt.Println("data", data)
		messageProcess(data, conn)
	}
}

func messageProcess(msg common.Data, conn *websocket.Conn) (err error) {
	switch msg.Type {
	case common.TypeLogin:
		fmt.Println("do login")
		u := UserProcess{conn: conn}
		u.Login(msg)
	case common.TypeRegister:
		fmt.Println("do regist")
		u := UserProcess{conn: conn}
		u.Register(msg)
	case common.TypeCommonMsag:
		fmt.Println("send msg")
		SendMsg(msg)
	default:

	}
	return
}
