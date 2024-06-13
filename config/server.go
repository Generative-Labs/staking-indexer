package config

import "time"

type ServerConfig struct {
	// The server bind address.
	Host string `long:"serverhost" description:"The api server bind address."`
	Port int    `long:"serverport" description:"The api server bind address."`

	WriteTimeout time.Duration
	ReadTimeout  time.Duration
	IdleTimeout  time.Duration
}

func DefaultServerConfig() *ServerConfig {
	return &ServerConfig{
		Host:         "127.0.0.1",
		Port:         8088,
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Second * 10,
	}
}
