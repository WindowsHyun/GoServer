package mongo

import (
	"GoServer/config"
	"GoServer/config/database"
	"GoServer/config/define"
	"context"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Clients []*mongo.Client

type MongoRepository struct {
	Client         *mongo.Client
	Collection     *mongo.Collection
	DatabaseName   string
	CollectionName string
	HashKey        []string
	IndexType      int
}

func InitializeMongo(ctx context.Context, config *config.Config) (map[string]MongoInterface, error) {
	appSvrCfg := config.GetMongo(define.MongoApp)
	apiSvrCfg := config.GetMongo(define.MongoApi)
	commonSvrCfg := config.GetMongo(define.MongoCommon)

	fields := []struct{ Host, User, Pass string }{appSvrCfg, apiSvrCfg, commonSvrCfg}
	dbRepos := make(map[string]MongoInterface)

	for fieldKey, cfg := range fields {
		ctx := context.Background()
		if cfg.Host == "" || cfg.User == "" || cfg.Pass == "" {
			continue
		}

		credential := options.Credential{
			Username: cfg.User,
			Password: cfg.Pass,
		}

		clientOptions := options.Client().ApplyURI(cfg.Host).SetAuth(credential)
		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			return nil, errors.Wrap(err, "Failed mongo connect")
		}

		err = client.Ping(ctx, nil)
		if err != nil {
			return nil, errors.Unwrap(err)
		}
		Clients = append(Clients, client)

		for key, colInfo := range database.CollectionInfos {
			if colInfo.DatabaseLocation == fieldKey {
				repo, err := CreateDBRepository(ctx, client, colInfo.DatabaseName, colInfo.CollectionName, colInfo.HashKey, colInfo.IndexType)
				if err != nil {
					return nil, errors.Wrap(err, "mongo repo failed")
				}
				dbRepos[key] = repo
			}
		}
	}

	return dbRepos, nil
}

func CreateDBRepository(ctx context.Context, client *mongo.Client, databaseName string, collectionName string, hashKey []string, indexType int) (*MongoRepository, error) {
	repo := &MongoRepository{
		Client:         client,
		DatabaseName:   databaseName,
		CollectionName: collectionName,
		HashKey:        hashKey,
		IndexType:      indexType,
		Collection:     client.Database(databaseName).Collection(collectionName),
	}

	if repo.Collection == nil {
		return nil, errors.New("Failed to create collection")
	}

	if err := CreateIndex(ctx, repo); err != nil {
		return nil, err
	}
	return repo, nil
}

func CreateIndex(ctx context.Context, repo *MongoRepository) error {
	var indexModels []mongo.IndexModel

	switch repo.IndexType {
	case define.IndexTypeCompound:
		indexKeys := bson.D{}
		for _, key := range repo.HashKey {
			indexKeys = append(indexKeys, bson.E{Key: key, Value: 1})
		}
		indexModels = append(indexModels, mongo.IndexModel{
			Keys:    indexKeys,
			Options: options.Index().SetUnique(true),
		})

	case define.IndexTypeSingle:
		for _, key := range repo.HashKey {
			indexModels = append(indexModels, mongo.IndexModel{
				Keys:    bson.D{{Key: key, Value: 1}},
				Options: options.Index().SetUnique(!strings.Contains(key, "$**")),
			})
		}
	default:
		return errors.New("invalid index type:" + fmt.Sprint(repo.IndexType))
	}

	if len(indexModels) == 0 {
		return nil
	}

	// 기존 인덱스 존재 여부 확인
	cursor, err := repo.Collection.Indexes().List(ctx)
	if err != nil {
		return fmt.Errorf("list indexes: %w", err)
	}
	defer cursor.Close(ctx)

	if cursor.Next(ctx) {
		if _, err := repo.Collection.Indexes().DropAll(ctx); err != nil {
			return fmt.Errorf("drop existing indexes: %w", err)
		}
		fmt.Println("Existing indexes dropped.")
	} else {
		fmt.Println("No existing indexes found. Skipping drop operation.")
	}

	if _, err := repo.Collection.Indexes().CreateMany(ctx, indexModels); err != nil {
		return fmt.Errorf("create indexes: %w", err)
	}
	return nil
}

func CloseMongo(ctx context.Context) {
	for _, client := range Clients {
		if err := client.Disconnect(ctx); err != nil {
			fmt.Println("CloseMongo Err:", err)
		}
	}
	Clients = nil
}
