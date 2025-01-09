package config

import (
	"GoServer/config/structure"
)

// Go에서 Reflection은 Public만 Decode 할 수 있다.
// Config = ReflectionConfig 동일하게 유지해준다.
type Config struct {
	server    structure.Server
	log       structure.Log
	env       structure.ENV
	secretKey structure.SecretKey
	mongo     structure.MongoMap
	mysql     structure.MySQLMap
	redis     structure.RedisMap
}

type ReflectionConfig struct {
	Server    structure.Server
	Log       structure.Log
	ENV       structure.ENV
	SecretKey structure.SecretKey
	Mongo     structure.MongoMap
	MySQL     structure.MySQLMap
	Redis     structure.RedisMap
}

func (c *Config) IsDevelop() bool {
	target := c.GetTarget()
	return target == "local" || target == "develop"
}

func (c *Config) GetTarget() string {
	return c.env.Target
}

func (c *Config) SetTarget(target string) {
	c.env.Target = target
}

func (c *Config) GetServer() structure.Server {
	return c.server
}

func (c *Config) SetServer(server structure.Server) {
	c.server = server
}

func (c *Config) GetMongo(targetDB string) structure.MongoConfig {
	envTarget := c.GetTarget()

	mgInfoMap, envExists := c.mongo[envTarget]
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

	c.mongo[envTarget][targetDB] = config
}

func (c *Config) GetMySQL() structure.MySQLConfig {
	envTg := c.GetTarget()

	mySQLConfig, envExists := c.mysql[envTg]
	if !envExists {
		return structure.MySQLConfig{}
	}

	return mySQLConfig
}

func (c *Config) SetMySQL(config structure.MySQLConfig) {
	envTg := c.GetTarget()
	c.mysql[envTg] = config
}

func (c *Config) SetRedis(config structure.RedisConfig) {
	envTg := c.GetTarget()
	c.redis[envTg] = config
}

func (c *Config) GetRedis() structure.RedisConfig {
	envTg := c.GetTarget()
	redis, envExists := c.redis[envTg]
	if !envExists {
		return structure.RedisConfig{}
	}

	return redis
}

func (c *Config) GetLog() structure.Log {
	return c.log
}

func (c *Config) SetLog(log structure.Log) {
	c.log = log
}

func (c *Config) GetSecretKey() structure.SecretKey {
	return c.secretKey
}

func (c *Config) SetSecretKey(key structure.SecretKey) {
	c.secretKey = key
}
