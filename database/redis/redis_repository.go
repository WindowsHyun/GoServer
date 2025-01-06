package redis

import (
	"GoServer/model"

	"github.com/go-redis/redis/v8"
)

type RedisUserRepository struct {
	client *redis.Client
}

func NewRedisUserRepository(client *redis.Client) *RedisUserRepository {
	return &RedisUserRepository{client: client}
}

func (r *RedisUserRepository) GetUserByID(id int) (*model.User, error) {
	return &model.User{}, nil
}

func (r *RedisUserRepository) SaveUser(user *model.User) error {
	return nil
}
