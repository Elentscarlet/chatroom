package process

import (
	"chat/common"
	"chat/server/model"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
)

type UserProcess struct {
	conn *websocket.Conn
}

func register(userName, password string) (user model.User, err error) {
	user, err = model.CurrentUserDao.Register(userName, password)
	return
}

func login(userName, password string) (user model.User, err error) {
	user, err = model.CurrentUserDao.Login(userName, password)
	return
}

func (proc *UserProcess) reponse(Type string, code int, data string, err error) {
	var rspMsg common.ResponseMsg
	rspMsg.Code = code
	rspMsg.Type = Type
	rspMsg.Data = data
	rspData, err := json.Marshal(rspMsg)
	if err != nil {
		fmt.Printf("error in response:%v\n", err)
	}
	write(rspData, proc.conn)
}

func (proc *UserProcess) Register(msg common.Data) (err error) {
	var registerData common.RegisterMsg
	var code int
	err = json.Unmarshal([]byte(msg.Content), &registerData)
	_, err = register(registerData.UserName, registerData.Password)
	switch err {
	case nil:
		code = common.RegisterSucceed
		fmt.Println("sucess")
	default:
		code = common.HasExited
		fmt.Println("failed")
	}
	proc.reponse(common.TypeResopnse, code, "", err)
	return
}

func (proc *UserProcess) Login(msg common.Data) (err error) {
	var loginData common.LoginMsg
	var code int
	err = json.Unmarshal([]byte(msg.Content), &loginData)
	user, err := login(loginData.UserName, loginData.Password)
	switch err {
	case nil:
		code = common.LoginSucceed
		fmt.Println("sucess")
		clientConn := model.ClientConn{}
		clientConn.Save(user.ID, user.Name, proc.conn)
	default:
		code = common.LoginError
		fmt.Println("failed")
	}
	proc.reponse(common.TypeLoginReply, code, "", err)
	return
}
