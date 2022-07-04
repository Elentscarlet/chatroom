package model

import (
	"github.com/gorilla/websocket"
)

type ClientConn struct {
}

type ConnInfo struct {
	Conn     *websocket.Conn
	UserName string
}

var ClientMap map[int]ConnInfo

func init() {
	ClientMap = make(map[int]ConnInfo)
}

func (c ClientConn) Save(userID int, name string, userConn *websocket.Conn) {
	ClientMap[userID] = ConnInfo{userConn, name}
}

func (c ClientConn) Del(userConn *websocket.Conn) {
	for id, info := range ClientMap {
		if userConn == info.Conn {
			delete(ClientMap, id)
		}
	}
}

func (c ClientConn) SearchByName(userName string) (ConnInfo *websocket.Conn, err error) {
	user, err := CurrentUserDao.GetByName(userName)
	if err != nil {
		return
	}
	ConnInfo = ClientMap[user.ID].Conn
	return
}
