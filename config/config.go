package config

import (
	"GoServer/config/structure"
)

type Config struct {
	Server    structure.Server
	Mongo     structure.MongoMap
	Log       structure.Log
	ENV       structure.ENV
	SecretKey structure.SecretKey
}

func (c *Config) GetTarget() string {
	return c.ENV.Target
}

func (c *Config) SetTarget(target string) {
	c.ENV.Target = target
}

func (c *Config) GetServer() structure.Server {
	return c.Server
}

func (c *Config) SetServer(server structure.Server) {
	c.Server = server
}

func (c *Config) GetMongo(targetDB string) structure.MongoConfig {
	envTarget := c.GetTarget()

	mgInfoMap, envExists := c.Mongo[envTarget]
	if !envExists {
		return structure.MongoConfig{}
	}

	mongoConfig, dbExists := mgInfoMap[targetDB]
	if !dbExists {
		return structure.MongoConfig{}
	}

	return mongoConfig
}

func (c *Config) SetMongo(targetDB string, config structure.MongoConfig) {
	envTarget := c.GetTarget()

	c.Mongo[envTarget][targetDB] = config
}

func (c *Config) GetLog() structure.Log {
	return c.Log
}

func (c *Config) SetLog(log structure.Log) {
	c.Log = log
}

func (c *Config) GetSecretKey() structure.SecretKey {
	return c.SecretKey
}

func (c *Config) SetSecretKey(key structure.SecretKey) {
	c.SecretKey = key
}
