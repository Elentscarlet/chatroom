package main

import (
	"chat/conf"
	"chat/server/model"
	"chat/server/process"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

var pool *redis.Pool

func init() {
	redisInfo := conf.Config.RedisInfo
	initRedisPool(redisInfo.MaxIdle, redisInfo.MaxActive, time.Second*(redisInfo.IdleTimeout), redisInfo.Host)

	model.InitUserDao(pool)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/ws", process.MainHandler)
	serverInfo := conf.Config.ServerInfo
	fmt.Println(serverInfo)
	if err := http.ListenAndServe(serverInfo.Host, router); err != nil {
		fmt.Println("err:", err)
	}
}

func initRedisPool(maxIdle, maxActive int, idleTimeout time.Duration, host string) {
	pool = &redis.Pool{
		// 初始化链接数量
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: idleTimeout,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", host)
		},
	}
}
