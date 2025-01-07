package mongo

import (
	// Replace with your actual config package
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m *MongoRepository) GetClient() *mongo.Client {
	if m.Client == nil {
		return nil
	}
	return m.Client
}

func (m *MongoRepository) GetCollection() *mongo.Collection {
	if m.Collection == nil {
		m.Collection = m.Client.Database(m.DatabaseName).Collection(m.CollectionName)
	}
	return m.Collection
}

func (m *MongoRepository) IsExist(ctx context.Context, query interface{}) (bool, error) {
	var result bson.M
	err := m.GetCollection().FindOne(ctx, query).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (m *MongoRepository) Insert(ctx context.Context, doc interface{}) error {
	_, err := m.GetCollection().InsertOne(ctx, doc)
	return err
}

func (m *MongoRepository) Delete(ctx context.Context, query interface{}) error {
	_, err := m.GetCollection().DeleteMany(ctx, query)
	return err
}

func (m *MongoRepository) Update(ctx context.Context, query interface{}, doc interface{}) error {
	_, err := m.GetCollection().UpdateMany(ctx, query, bson.M{"$set": doc})
	return err
}

func (m *MongoRepository) UpdateField(ctx context.Context, query interface{}, field string, value interface{}) (bool, error) {
	update := bson.M{"$set": bson.M{field: value}}
	result, err := m.GetCollection().UpdateMany(ctx, query, update)
	if err != nil {
		return false, err
	}
	return result.ModifiedCount > 0, nil
}

func (m *MongoRepository) Upsert(ctx context.Context, query interface{}, doc interface{}) error {
	_, err := m.GetCollection().UpdateOne(ctx, query, bson.M{"$set": doc}, options.Update().SetUpsert(true))
	return err
}

func (m *MongoRepository) GetAllData(ctx context.Context, query interface{}, result interface{}) error {
	cursor, err := m.GetCollection().Find(ctx, query)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)
	return cursor.All(ctx, result)
}

func (m *MongoRepository) GetData(ctx context.Context, query interface{}, field string, result interface{}) error {
	projection := bson.M{field: 1, "_id": 0}
	return m.GetProjectionData(ctx, query, projection, field, result)
}

func (m *MongoRepository) GetProjectionData(ctx context.Context, query interface{}, projection interface{}, field string, result interface{}) error {
	cursor, err := m.GetCollection().Find(ctx, query, options.Find().SetProjection(projection))
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)
	return cursor.All(ctx, result)
}

func (m *MongoRepository) GetMultiplePartitionData(ctx context.Context, query interface{}, skip int, limit int, result interface{}) error {
	opts := options.Find()
	if skip > 0 {
		opts.SetSkip(int64(skip))
	}
	if limit > 0 {
		opts.SetLimit(int64(limit))
	}
	cursor, err := m.GetCollection().Find(ctx, query, opts)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)
	return cursor.All(ctx, result)
}

func (m *MongoRepository) CreateTTLIndex(ctx context.Context, fieldName string, expireAfterSeconds int32) error {
	index := mongo.IndexModel{
		Keys:    bson.M{fieldName: 1},
		Options: options.Index().SetExpireAfterSeconds(expireAfterSeconds),
	}
	_, err := m.GetCollection().Indexes().CreateOne(ctx, index)
	return err
}

func (m *MongoRepository) DropTTLIndex(ctx context.Context, fieldName string) error {
	indexName := fieldName + "_1"
	_, err := m.GetCollection().Indexes().DropOne(ctx, indexName)
	return err
}

func (m *MongoRepository) FindOne(ctx context.Context, query interface{}, projection interface{}, result interface{}) error {
	var opts *options.FindOneOptions
	if projection != nil {
		opts = &options.FindOneOptions{Projection: projection}
	}
	err := m.GetCollection().FindOne(ctx, query, opts).Decode(result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New("document not found")
		}
		return err
	}
	return nil
}

func (m *MongoRepository) UpdateMany(ctx context.Context, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	return m.GetCollection().UpdateMany(ctx, filter, update)
}

func (m *MongoRepository) Find(ctx context.Context, filter interface{}) (*mongo.Cursor, error) {
	return m.GetCollection().Find(ctx, filter)
}

func (m *MongoRepository) IncreaseField(ctx context.Context, query interface{}, field string, value interface{}) (int64, error) {
	update := bson.M{"$inc": bson.M{field: value}}
	result, err := m.GetCollection().UpdateOne(ctx, query, update)
	if err != nil {
		return 0, err
	}
	return result.ModifiedCount, nil
}

func (m *MongoRepository) DocCountFilter(ctx context.Context, filter interface{}) (int64, error) {
	count, err := m.GetCollection().CountDocuments(ctx, filter)
	return count, err
}

func (m *MongoRepository) Aggregate(ctx context.Context, pipeline mongo.Pipeline) ([]bson.M, error) {
	cursor, err := m.GetCollection().Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	return results, nil
}

func (m *MongoRepository) AddArrayField(ctx context.Context, query interface{}, field string, value interface{}) error {
	update := bson.M{"$push": bson.M{field: value}}
	_, err := m.GetCollection().UpdateMany(ctx, query, update)
	return err
}

func (m *MongoRepository) SetArrayField(ctx context.Context, query interface{}, field string, value interface{}) error {
	update := bson.M{"$set": bson.M{field: value}}
	_, err := m.GetCollection().UpdateMany(ctx, query, update)
	return err
}

func (m *MongoRepository) DelArrayField(ctx context.Context, query interface{}, field string, value interface{}) error {
	update := bson.M{"$pull": bson.M{field: value}}
	_, err := m.GetCollection().UpdateMany(ctx, query, update)
	return err
}

func (m *MongoRepository) DropCollection(ctx context.Context) error {
	return m.GetCollection().Drop(ctx)
}
