package mysql

import (
	"context"
	"database/sql"
)

type MySQLInterface interface {
	GetDB() *sql.DB
	IsExist(ctx context.Context, query string, args ...interface{}) (bool, error)
	Insert(ctx context.Context, query string, args ...interface{}) error
	Delete(ctx context.Context, query string, args ...interface{}) error
	Update(ctx context.Context, query string, args ...interface{}) error
	UpdateField(ctx context.Context, table, field, value, where string, args ...interface{}) (int64, error)
	Upsert(ctx context.Context, query string, args ...interface{}) error
	GetAllData(ctx context.Context, query string, dest interface{}, args ...interface{}) error
	GetData(ctx context.Context, query string, dest interface{}, args ...interface{}) error
	GetMultiplePartitionData(ctx context.Context, query, orderBy string, offset, limit int, dest interface{}, args ...interface{}) error
	CreateIndex(ctx context.Context, table, fieldName string) error
	DropIndex(ctx context.Context, table, indexName string) error
	FindOne(ctx context.Context, query string, dest interface{}, args ...interface{}) error
	UpdateMany(ctx context.Context, query string, args ...interface{}) (int64, error)
	Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Count(ctx context.Context, query string, args ...interface{}) (int64, error)
	Transaction(ctx context.Context, fn func(tx *sql.Tx) error) error
	Close(ctx context.Context)
}
