package conf

import (
	"encoding/json"
	logf "github.com/sirupsen/logrus"
	"os"
	"time"
)

type config struct {
	ServerInfo serverInfo
	RedisInfo  redisInfo
}

type serverInfo struct {
	Host string
}

type redisInfo struct {
	Host        string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var Config = config{}

//读取redis设置
func init() {
	file, err := os.Open("conf/redis.json")
	defer file.Close()
	if err != nil {
		logf.Error(err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)
	if err != nil {
		logf.Error(err)
	}
}
