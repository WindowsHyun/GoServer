package database

import (
	"GoServer/config/define"
	"GoServer/config/structure"
)

var MongoCollectionInfos = map[string]structure.MongoCollectionInfo{
	"UserInfo": {
		DatabaseName:   "User",
		CollectionName: "info",
		HashKey:        []string{"email"},
		IndexType:      define.IndexTypeSingle,
	},
	"Menu": {
		DatabaseName:   "Menu",
		CollectionName: "info",
		HashKey:        []string{"name"},
		IndexType:      define.IndexTypeSingle,
	},
	"UserPostBox": {
		DatabaseName:   "User",
		CollectionName: "PostBox",
		HashKey:        []string{},
		IndexType:      define.IndexTypeSingle,
	},
}

var MySQLCollectionInfos = map[string]structure.MySQLCollectionInfo{
	"UserTable": {
		TableName: "user_table",
	},
	"Menu": {
		TableName: "menu",
	},
}
