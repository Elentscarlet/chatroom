package process

import (
	"chat/common"
	"chat/server/model"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
)

func SendMsg(msg common.Data) (err error) {
	var sendmsg common.CommonMsg
	err = json.Unmarshal([]byte(msg.Content), &sendmsg)
	if err != nil {
		fmt.Println("some error when json unmarshal: %v\n", err)
	}
	sourceUser := sendmsg.UserName
	fmt.Println(sourceUser)
	var ResponseMsg common.ResponseMsg
	ResponseMsg.Type = common.TypeResopnse
	ResponseMsg.Data = sourceUser + ":" + sendmsg.Content
	data, err := json.Marshal(ResponseMsg)
	if err != nil {
		fmt.Println("some error when json marshal:%v\n", err)
	}
	for _, info := range model.ClientMap {
		fmt.Println(info.UserName)
		write(data, info.Conn)
	}
	return
}

func write(data []byte, conn *websocket.Conn) {
	var dataLen uint32
	dataLen = uint32(len(data))
	var bytes [4]byte
	binary.BigEndian.PutUint32(bytes[0:4], dataLen)
	err := conn.WriteMessage(websocket.TextMessage, data)
	fmt.Printf("send msg to %v\n", conn.RemoteAddr())
	if err != nil {
		fmt.Printf("send msg error %v\n", err)
		return
	}
	return
}
