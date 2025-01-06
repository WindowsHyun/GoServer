// repository/init.go
package repository

import (
	"GoServer/config"
	"GoServer/db/mongo"
)

type Repositories struct {
	User UserRepository
	App  AppRepository
}

func NewRepositories() Repositories {
	repos := Repositories{}

	for _, dbCfg := range config.DBConfigs {
		client := mongo.Clients[dbCfg.ClientName]
		for repoName, collectionCfg := range dbCfg.Collections {
			collection := client.Database(dbCfg.Database).Collection(collectionCfg.CollectionName)

			switch repoName {
			case "User":
				repos.User = NewMongoUserRepository(collection)
			case "App":
				repos.App = NewMongoAppRepository(collection)
				// Add more cases for additional repositories
			}
		}
	}

	return repos
}
