package process

import (
	"chat/server/model"
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

func (proc *UserProcess) Register(userName, password, passwordConfirm string) (err error) {
	return err
}

func (proc *UserProcess) Login(userName, password string) (err error) {
	return err
}
