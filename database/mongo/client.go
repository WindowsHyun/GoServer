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

var mongoClients = make(map[string]*mongo.Client)

type MongoRepository struct {
	Client         *mongo.Client
	Collection     *mongo.Collection
	DatabaseName   string
	CollectionName string
	HashKey        []string
	IndexType      int
}

func Initialize(ctx context.Context, config *config.Config) (map[string]MongoInterface, error) {
	svrCgf := config.GetMongo()

	fields := []struct{ Host, User, Pass string }{
		{svrCgf.Host, svrCgf.User, svrCgf.Pass},
	}

	dbRepos := make(map[string]MongoInterface)
	for _, cfg := range fields {
		ctx := context.Background()
		if cfg.Host == "" || cfg.User == "" || cfg.Pass == "" {
			continue
		}

		credential := options.Credential{
			Username: cfg.User,
			Password: cfg.Pass,
		}

		for key, colInfo := range database.MongoCollectionInfos {
			client, err := initializeMongoClient(ctx, colInfo.DatabaseName, cfg.Host, credential)
			if err != nil {
				return nil, errors.Wrap(err, "initialize Mongo Client")
			}

			repo, err := CreateDBRepository(ctx, client, colInfo.DatabaseName, colInfo.CollectionName, colInfo.HashKey, colInfo.IndexType)
			if err != nil {
				return nil, errors.Wrap(err, "mongo repo failed")
			}
			dbRepos[key] = repo
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

func initializeMongoClient(ctx context.Context, dbName, uri string, credential options.Credential) (*mongo.Client, error) {
	key := dbName
	if client, exists := mongoClients[key]; exists {
		return client, nil
	}

	clientOptions := options.Client().ApplyURI(uri).SetAuth(credential)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to initialize MongoDB client")
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, errors.Wrap(err, "MongoDB connection failed")
	}

	mongoClients[key] = client
	return client, nil
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
	}

	if _, err := repo.Collection.Indexes().CreateMany(ctx, indexModels); err != nil {
		return fmt.Errorf("create indexes: %w", err)
	}
	return nil
}

func Close(ctx context.Context) {
	for dbName, client := range mongoClients {
		if err := client.Disconnect(ctx); err != nil {
			fmt.Printf("Error closing MongoDB client (%s): %v\n", dbName, err)
		}
	}
	mongoClients = make(map[string]*mongo.Client) // 맵 초기화
}
