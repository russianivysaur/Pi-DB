package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"pidb/types"
)

type BufferPoolConfig struct {
	PageSize  types.PageSize `json:"page_size"`
	PageCount int            `json:"page_count"`
}

type ServerConfig struct {
	ClientReadBufferSize int `json:"client_read_buffer_size"`
}
type Config struct {
	PoolConf   BufferPoolConfig `json:"buffer_pool"`
	ServerConf ServerConfig     `json:"server_config"`
}

func NewConfig(configFilePath string) *Config {
	var config Config
	configFile, err := os.OpenFile(configFilePath, os.O_RDONLY, 0444)
	if err != nil {
		log.Fatalln(err)
	}
	yamlDecoder := yaml.NewDecoder(configFile)
	err = yamlDecoder.Decode(&config)
	if err != nil {
		log.Fatalln(err)
	}
	return &config
}
