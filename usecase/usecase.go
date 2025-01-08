package usecase

import (
	IMongo "GoServer/database/mongo"
	IMySQL "GoServer/database/mysql"
	IRedis "GoServer/database/redis"
)

func InitUsecase(mongo map[string]IMongo.MongoInterface, mysql map[string]IMySQL.MySQLInterface, redis map[string]IRedis.RedisInterface) {
	// Initialize usecase
}
