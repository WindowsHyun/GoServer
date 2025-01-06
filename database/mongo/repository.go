// repository/init.go
package mongo

import (
	"GoServer/config"
)

type Repositories struct {
	User *UserRepository
	App  *AppRepository
	// Add other repositories as needed
}

func NewRepositories() Repositories {
	repos := Repositories{}

	for _, dbCfg := range config.DBConfigs {
		clientName := dbCfg.ClientName
		client, exists := Clients[clientName] // Ensure Clients is defined and holds the MongoDB clients
		if !exists {
			// Handle the error, client not found
			continue
		}

		db := client.Database(dbCfg.Database)

		for repoName, collectionCfg := range dbCfg.Collections {
			collection := db.Collection(collectionCfg.CollectionName)

			creator, exists := repositoryCreators[repoName]
			if exists {
				repo := creator(collection)
				// Assign the repository to the respective field in 'repos'
				switch repoName {
				case "User":
					repos.User = repo.(*UserRepository)
				case "App":
					repos.App = repo.(*AppRepository)
					// Add more cases as needed
				}
			} else {
				// Handle the case where the repository creator is not found
			}
		}
	}

	return repos
}
