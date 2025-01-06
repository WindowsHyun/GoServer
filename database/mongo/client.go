// db/mongo/client.go
package mongo

import (
	"GoServer/config"
	"GoServer/config/define"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

var Clients map[string]*mongo.Client

func InitializeMongo(ctx context.Context, config *config.Config) error {
	appSvrCfg := config.GetMongo(define.MongoApp)
	apiSvrCfg := config.GetMongo(define.MongoApi)
	commonSvrCfg := config.GetMongo(define.MongoCommon)

	fields := []struct{ Host, User, Pass string }{appSvrCfg, apiSvrCfg, commonSvrCfg}
	dbRepos := make(map[string]MongoInterface)

	return nil
}

func CloseMongo(ctx context.Context) {
	for _, client := range Clients {
		if err := client.Disconnect(ctx); err != nil {
			fmt.Println("CloseMongo Err:", err)
		}
	}
}
