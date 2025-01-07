package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoInterface interface {
	GetClient() *mongo.Client
	GetCollection() *mongo.Collection
	IsExist(ctx context.Context, query interface{}) (bool, error)
	Insert(ctx context.Context, doc interface{}) error
	Delete(ctx context.Context, query interface{}) error
	Update(ctx context.Context, query interface{}, doc interface{}) error
	UpdateField(ctx context.Context, query interface{}, field string, value interface{}) (bool, error)
	Upsert(ctx context.Context, query interface{}, doc interface{}) error
	GetAllData(ctx context.Context, query interface{}, result interface{}) error
	GetData(ctx context.Context, query interface{}, field string, result interface{}) error
	GetProjectionData(ctx context.Context, query interface{}, projection interface{}, field string, result interface{}) error
	GetMultiplePartitionData(ctx context.Context, query interface{}, skip int, limit int, result interface{}) error
	CreateTTLIndex(ctx context.Context, fieldName string, expireAfterSeconds int32) error
	DropTTLIndex(ctx context.Context, fieldName string) error
	FindOne(ctx context.Context, query interface{}, projection interface{}, result interface{}) error
	UpdateMany(ctx context.Context, filter interface{}, update interface{}) (*mongo.UpdateResult, error)
	Find(ctx context.Context, filter interface{}) (*mongo.Cursor, error)
	IncreaseField(ctx context.Context, query interface{}, field string, value interface{}) (int64, error)
	DocCountFilter(ctx context.Context, filter interface{}) (int64, error)
	Aggregate(ctx context.Context, pipeLine mongo.Pipeline) ([]bson.M, error)
	AddArrayField(ctx context.Context, query interface{}, field string, value interface{}) error
	SetArrayField(ctx context.Context, query interface{}, field string, value interface{}) error
	DelArrayField(ctx context.Context, query interface{}, field string, value interface{}) error
	DropCollection(ctx context.Context) error
}
