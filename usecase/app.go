package usecase

import (
	"GoServer/config/define"
	IMongo "GoServer/database/mongo"
	IMySQL "GoServer/database/mysql"
	IRedis "GoServer/database/redis"
	"GoServer/model"
	"context"

	"github.com/pkg/errors"
)

type appRepository struct {
	userInfo  IMongo.MongoInterface
	menu      IMongo.MongoInterface
	ranking   IRedis.RedisInterface
	userTable IMySQL.MySQLInterface
}

type AppUsecase interface {
	GetMenu(ctx context.Context) ([]model.Menu, error)
}

func NewAppRepository(mongo map[string]IMongo.MongoInterface, mysql map[string]IMySQL.MySQLInterface, redis map[string]IRedis.RedisInterface) (AppUsecase, error) {
	if err := IsValidRepoKey(mongo, "UserInfo", "Menu"); err != nil {
		return nil, errors.Wrap(err, "Invalid Mongo key")
	}
	if err := IsValidRepoKey(mysql, "UserTable", "Menu"); err != nil {
		return nil, errors.Wrap(err, "Invalid MySQL key")
	}
	if err := IsValidRepoKey(redis, define.RedisRankingDB, define.RedisUserDB); err != nil {
		return nil, errors.Wrap(err, "Invalid Redis key")
	}

	return &appRepository{
		userInfo:  mongo["UserInfo"],
		menu:      mongo["Menu"],
		ranking:   redis[define.RedisRankingDB],
		userTable: mysql["UserTable"],
	}, nil
}

func (a *appRepository) GetMenu(ctx context.Context) ([]model.Menu, error) {
	return nil, nil
}
