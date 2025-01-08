package redis

import (
	"GoServer/config"
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

var Clients []*redis.Client

type RedisRepository struct {
	Client *redis.Client
}

func Initialize(ctx context.Context, config *config.Config) (map[string]RedisInterface, error) {
	svrCfg := config.GetRedis()

	fields := []struct{ Host, Port, Pass string }{
		{svrCfg.Host, svrCfg.Port, svrCfg.Pass},
	}

	dbRepos := make(map[string]RedisInterface)

	for fieldKey, cfg := range fields {
		if cfg.Host == "" || cfg.Port == "" {
			continue
		}

		for i := 0; i < 10; i++ {
			opt := &redis.Options{
				Addr:     cfg.Host + ":" + cfg.Port,
				Password: cfg.Pass,
				DB:       i,
			}

			client := redis.NewClient(opt)
			Clients = append(Clients, client)
			if i == 0 {
				// Ping 체크는 1회만 진행한다.
				_, err := client.Ping(ctx).Result()
				if err != nil {
					return nil, fmt.Errorf("failed to ping redis %d: %v", fieldKey, err)
				}
			}
			dbRepos[fmt.Sprintf("%d", i)] = &RedisRepository{Client: client}
		}
	}

	return dbRepos, nil
}

func Close(ctx context.Context) {
	for _, client := range Clients {
		if err := client.Close(); err != nil {
			log.Printf("CloseRedis Err: %v", err)
		}
	}
	Clients = nil
}
