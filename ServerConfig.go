package webconfig

import (
	"fmt"
)

// [ServerConfig]

func NewServerConfig() *ServerConfig { // {{{
	return &ServerConfig{}
} // }}}

type ServerConfig struct {
	BaseConfig

	Host      string
	Addr      string
	PoolSize  int // Size of poll
	QueueSize int // Size of Queue
	LogMode   bool
}

func (this *ServerConfig) Name() string { // {{{
	return CONFIG_SERVER
} // }}}
