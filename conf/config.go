package conf

import (
	"github.com/koding/multiconfig"
)

var Config *APIConfig

type APIConfig struct {
	AreaConfig *AreaConfig
}

type AreaConfig struct {
	RedisConfig RedisConfig
	KafkaConfig KafkaConfig
}

type RedisConfig struct {
	Addr           string
	ConnectTimeout int
	ReadTimeout    int
	WriteTimeout   int
	Password       string
}

type KafkaConfig struct {
	Address      string
	ServiceTopic string
}

func init() {
	Config = new(APIConfig)
}

func (ac *APIConfig) LoadConfigFile(confPath string) error {
	m := &multiconfig.TOMLLoader{Path: confPath}
	areaConfig := new(AreaConfig)
	if err := m.Load(areaConfig); err != nil {
		return err
	}
	ac.AreaConfig = areaConfig
	ac.AreaConfig.RedisConfig = areaConfig.RedisConfig
	ac.AreaConfig.KafkaConfig = areaConfig.KafkaConfig

	return nil
}
