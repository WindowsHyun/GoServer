package config

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pelletier/go-toml/v2"
)

func InitConfig() *Config {
	// ENV
	if envErr := godotenv.Load(); envErr != nil {
		checkEnvVariables("TARGET")
	}
	checkEnvVariables("TARGET")

	envTarget := os.Getenv("TARGET")
	config := SettingConfig(envTarget)
	config.SetTarget(envTarget)
	log.Printf("Target : %s", config.GetTarget())
	return config
}

func SettingConfig(target string) *Config {
	confPath := flag.String("config", "./config/default.toml", "toml file to use for configuration")
	cf := DefaultLoadConfig(*confPath)

	return cf
}

func DefaultLoadConfig(path string) *Config {
	c := new(ReflectionConfig)
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if err := toml.NewDecoder(file).Decode(c); err != nil {
		panic(err)
	}

	return &Config{
		server:    c.Server,
		log:       c.Log,
		env:       c.ENV,
		secretKey: c.SecretKey,
		mongo:     c.Mongo,
		mysql:     c.MySQL,
		redis:     c.Redis,
	}
}

func checkEnvVariables(keys ...string) {
	for _, key := range keys {
		if os.Getenv(key) == "" {
			panic(".env data :" + key)
		}
	}
}
