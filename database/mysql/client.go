package mysql

import (
	"GoServer/config"
	"GoServer/config/database"
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

var Databases []*sql.DB

type MySQLRepository struct {
	DB           *sql.DB
	DatabaseName string
	TableName    string
}

func Initialize(ctx context.Context, config *config.Config) (map[string]MySQLInterface, error) {
	svrCgf := config.GetMySQL()
	fields := []struct{ Host, Port, User, Pass, DB string }{
		{svrCgf.Host, svrCgf.Port, svrCgf.User, svrCgf.Pass, svrCgf.DB},
	}

	dbRepos := make(map[string]MySQLInterface)

	for fieldKey, cfg := range fields {
		if cfg.Host == "" || cfg.User == "" || cfg.Pass == "" || cfg.DB == "" {
			continue
		}

		dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.User, cfg.Pass, cfg.Host, cfg.Port, cfg.DB)
		db, err := sql.Open("mysql", dataSourceName)
		if err != nil {
			log.Fatalf("Error connecting to %s database: %v", strconv.Itoa(fieldKey), err)
			return nil, err
		}
		Databases = append(Databases, db)

		for key, tableInfo := range database.MySQLCollectionInfos {
			repo, err := CreateDBRepository(db, svrCgf.DB, tableInfo.TableName)
			if err != nil {
				return nil, err
			}
			dbRepos[key] = repo
		}
	}

	return dbRepos, nil
}

func CreateDBRepository(db *sql.DB, databaseName, tableName string) (*MySQLRepository, error) {
	repo := &MySQLRepository{
		DB:           db,
		DatabaseName: databaseName,
		TableName:    tableName,
	}
	return repo, nil
}

func Close(ctx context.Context) {
	for _, db := range Databases {
		if err := db.Close(); err != nil {
			log.Println("CloseMySQL Err:", err)
		}
	}
	Databases = nil
}
