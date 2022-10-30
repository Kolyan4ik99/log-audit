package config

import "fmt"

type Config struct {
	Server *ServerConfig
	Db     *MongoConfig
}

func NewConfig() *Config {
	return &Config{
		Server: &ServerConfig{
			NetProtocol: "tcp",
			Address:     "",
			Port:        "8000",
		},
		Db: &MongoConfig{
			Uri:  "mongodb://localhost:27017",
			User: "admin",
			Pass: "qwerty123",
		},
	}
}

type ServerConfig struct {
	NetProtocol string
	Address     string
	Port        string
}

func (s ServerConfig) GetAddressPort() string {
	return fmt.Sprintf("%s:%s", s.Address, s.Port)
}

type MongoConfig struct {
	Uri  string // "MONGO_URI"
	User string // "MONGO_USER"
	Pass string // "MONGO_PASSWORD"
}
