package model

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var CurrentUserDao *UserDao

type UserDao struct {
	pool *redis.Pool
}

func InitUserDao(pool *redis.Pool) {
	CurrentUserDao = &UserDao{pool: pool}
}

func idIncr(conn redis.Conn) (id int, err error) {
	res, err := conn.Do("incr", "user_id")
	id = int(res.(int64))
	if err != nil {
		fmt.Println("id incr error:%v\n", err)
		return
	}
	return
}

func (this *UserDao) GetByName(userName string) (user User, err error) {
	conn := this.pool.Get()
	defer conn.Close()
	res, err := redis.String(conn.Do("hget", "users", userName))
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(res), &user)
	if err != nil {
		fmt.Printf("Unmarshal user info error: %v\n", err)
		return
	}
	return
}

func (this *UserDao) Register(userName, password string) (user User, err error) {
	conn := this.pool.Get()
	defer conn.Close()
	id, err := idIncr(conn)
	if err != nil {
		return
	}
	user, err = this.GetByName(userName)
	if err == nil {
		fmt.Printf("User already exists!\n")
		err = ErrorUserDoesNotExist
		return
	}
	user = User{
		ID:       id,
		Name:     userName,
		Password: password,
	}
	info, err := json.Marshal(user)
	if err != nil {
		fmt.Printf("json marshal error:%v", err)
		return
	}
	_, err = conn.Do("hset", "users", userName, info)
	if err != nil {
		fmt.Printf("json marshal error:%v", err)
		return
	}
	return
}

func (this *UserDao) Login(userName, password string) (user User, err error) {
	fmt.Println("uname", userName, "pw", password)
	user, err = this.GetByName(userName)
	if err != nil {
		fmt.Printf("get user by id error: %v\n", err)
		return
	}
	if user.Password != password {
		fmt.Printf("password error")
		return
	}
	return

}
