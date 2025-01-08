package usecase

import (
	"GoServer/config"
	IMongo "GoServer/database/mongo"
	IMySQL "GoServer/database/mysql"
	IRedis "GoServer/database/redis"
	"fmt"

	"github.com/pkg/errors"
)

type Usecase struct {
	UserUsecase UserUsecase
	AppUsecase  AppUsecase
}

func InitUsecase(config *config.Config, mongo map[string]IMongo.MongoInterface, mysql map[string]IMySQL.MySQLInterface, redis map[string]IRedis.RedisInterface) (*Usecase, error) {
	userUsecase, err := NewUserRepository(mongo)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize userUsecase")
	}
	appUsecase, err := NewAppRepository(mongo, mysql, redis)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize appUsecase")
	}

	return &Usecase{
		UserUsecase: userUsecase,
		AppUsecase:  appUsecase,
	}, nil
}

func IsValidRepoKey[T any](repoMap map[string]T, keys ...string) error {
	for _, key := range keys {
		if _, ok := repoMap[key]; !ok {
			return fmt.Errorf("%s key is invalid", key)
		}
	}
	return nil
}
