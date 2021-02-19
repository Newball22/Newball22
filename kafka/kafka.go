package kafka

import (
	"errors"
	"fmt"
	"log"
	client "other/gameserver/redisClient"
	"strings"
	"time"

	"github.com/Shopify/sarama"
	"github.com/go-redis/redis"
	"github.com/muesli/cache2go"
)

var (
	SendKafkaMsgTag   = false  // 是否发送消息至kafka
	kafkaAddress      []string //初始化kafka地址
	kafkaSyncProducer sarama.SyncProducer
	kafkaDataChan     = make(chan sarama.ProducerMessage, 10000)
	ServiceTopicName  string

	runningRedisKey = "KAFKA::RUN::TAG"  //是否挽歌卡夫卡里面写数据
	kafkaLocalCache *cache2go.CacheTable // 本地缓存标识，避免频繁调用redis
	kafkaOpenTag    = "open"             //Redis缓存中停止写入数据的标识

	//本地缓存做白名单
	sendKafkaServiceIdMap = map[int64]bool{
		1111: true,
		2222: true,
		3333: true,
	}
)

func InitKafkaConn(addr string, serviceTopic string) error {
	var (
		err error
	)

	kafkaAddress = strings.Split(addr, ",") //kafka集群时用逗号分隔
	ServiceTopicName = serviceTopic
	kafkaSyncProducer, err = sarama.NewSyncProducer(kafkaAddress, getKafkaConf())

	if err != nil {
		return errors.New(fmt.Sprintf("kafka.DialContext error: [adds: %v], err: [%v]", kafkaAddress, err))
	}

	kafkaLocalCache = cache2go.Cache("kafkaTagCache")

	//临时缓存消息，集中发送消息到kafka
	go syncChanDataToKafka()
	return nil
}

func getKafkaConf() *sarama.Config {
	conf := sarama.NewConfig()
	conf.Producer.Retry.Max = 1
	conf.Producer.RequiredAcks = sarama.WaitForAll
	conf.Producer.Return.Successes = true
	conf.Metadata.Full = true
	conf.Version = sarama.V2_6_0_0
	conf.Producer.Partitioner = sarama.NewRandomPartitioner
	return conf
}

//并行转串化，避免同时创建多个SyncProducer
func syncChanDataToKafka() {
	for data := range kafkaDataChan {
		if _, _, err := kafkaSyncProducer.SendMessage(&data); err != nil {
			if strings.Contains(err.Error(), "write tcp") {
				kafkaSyncProducer, err = sarama.NewSyncProducer(kafkaAddress, getKafkaConf())
				if err != nil { //重试
					_, _, err = kafkaSyncProducer.SendMessage(&data)
				}
			} else if strings.Contains(err.Error(), "timeout") {
				_, _, err = kafkaSyncProducer.SendMessage(&data)
			}

			if err != nil {
				log.Fatalf("syncChanDataToKafka warn: %v", err)
				return
			}
		}
	}
}

func SendDataTOKafka(data sarama.ProducerMessage) {
	kafkaDataChan <- data
}

//关闭kafka
func CloseKafkaConn() error {
	return kafkaSyncProducer.Close()
}

func JudgeSendTkV2Msg(id int64) bool {
	// 配置文件没有配置kafka信息，不发送消息
	if SendKafkaMsgTag == false {
		return false
	}
	// 不是指定的App不发送
	if !sendKafkaServiceIdMap[id] {
		return false
	}

	localCache, err := kafkaLocalCache.Value(runningRedisKey)
	if err == nil && localCache != nil {
		data, ok := localCache.Data().(string)
		if ok {
			return data == kafkaOpenTag
		}
	}

	redisData, err := client.RedisClient.Get(runningRedisKey).Result()
	if err != nil {
		if err == redis.Nil { // Redis 为空
			kafkaLocalCache.Add(runningRedisKey, time.Second*10, redisData)
			return false // 默认不发送
		}
		log.Fatalf("get redis msg error, key: %s, err: %v", runningRedisKey, err)
		return false
	}
	kafkaLocalCache.Add(runningRedisKey, time.Second*10, redisData)
	return redisData == kafkaOpenTag
}
