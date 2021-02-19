package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"os"
	"other/gameserver/conf"
	"other/gameserver/kafka"
	"other/gameserver/model"
	"other/gameserver/redisClient"
	"syscall"
	"time"
)

func main() {
	Exit := make(chan os.Signal)
	confPath := flag.String("c", "conf/test.conf", "configure file path")
	flag.Parse()

	// 读取配置文件
	config := conf.Config
	if err := config.LoadConfigFile(*confPath); err != nil {
		fmt.Println("load config failed: ", err.Error())
		Exit <- syscall.SIGTERM
	}

	//初始化Redis
	ipRedisConf := config.AreaConfig.RedisConfig
	redisClient.Init(ipRedisConf.Addr, ipRedisConf.Password, ipRedisConf.ConnectTimeout, ipRedisConf.ReadTimeout, ipRedisConf.WriteTimeout)
	if len(config.AreaConfig.KafkaConfig.Address) > 0 {
		if err := kafka.InitKafkaConn(config.AreaConfig.KafkaConfig.Address, config.AreaConfig.KafkaConfig.ServiceTopic); err != nil {
			log.Printf("Kafka init failed.err: %v\n", err)
			Exit <- syscall.SIGTERM
		}
		kafka.SendKafkaMsgTag = true
		defer kafka.CloseKafkaConn()
	}

	isOK := kafka.JudgeSendTkV2Msg(1111)
	user := model.NewUserInfo(1, "kafka", "博雅")
	data, _ := json.Marshal(user)
	if isOK {
		go func() {
			kafka.SendDataTOKafka(sarama.ProducerMessage{
				Topic: kafka.ServiceTopicName,
				Value: sarama.StringEncoder(data),
			})
		}()
	}
	time.Sleep(time.Second * 5)
}
