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
	confPath := flag.String("config", "./default.toml", "toml file to use for configuration")
	cf := DefaultLoadConfig(*confPath)

	return cf
}

func DefaultLoadConfig(path string) *Config {
	c := new(Config)
	if file, err := os.Open(path); err != nil {
		panic(err)
	} else {
		defer file.Close()
		if err := toml.NewDecoder(file).Decode(c); err != nil {
			panic(err)
		}
	}

	return &Config{
		Server:    c.Server,
		Log:       c.Log,
		Mongo:     c.Mongo,
		ENV:       c.ENV,
		SecretKey: c.SecretKey,
	}
}

func checkEnvVariables(keys ...string) {
	for _, key := range keys {
		if os.Getenv(key) == "" {
			panic(".env data :" + key)
		}
	}
}
