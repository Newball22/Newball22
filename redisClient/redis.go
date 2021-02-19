package redisClient

import (
	"errors"
	"log"
	"time"

	"github.com/go-redis/redis"
)

var RedisClient *redis.ClusterClient

func Init(addr, password string, connectTimeout, readTimeout, writeTimeout int) {
	RedisClient = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:              []string{addr},
		DialTimeout:        time.Duration(connectTimeout * 1000000),
		ReadTimeout:        time.Duration(readTimeout * 1000000),
		WriteTimeout:       time.Duration(writeTimeout * 1000000),
		IdleTimeout:        -1,
		IdleCheckFrequency: -1,
		Password:           password,
	})
	_, err := RedisClient.ClusterInfo().Result()
	if err != nil {
		log.Fatalf("error redis config err, err = %s", err)
	}
}

func GetIpKey(ip string) (string, error) {
	if ip == "" {
		return "", errors.New("ip is empty")
	}
	return "aii_" + ip, nil
}
