package database

import (
	"GoServer/config/define"
	"GoServer/config/structure"
)

var MongoCollectionInfos = map[string]structure.MongoCollectionInfo{
	"UserInfo": {
		DatabaseLocation: define.DBCommon,
		DatabaseName:     "User",
		CollectionName:   "info",
		HashKey:          []string{"email"},
		IndexType:        define.IndexTypeSingle,
	},
	"Menu": {
		DatabaseLocation: define.DBCommon,
		DatabaseName:     "Menu",
		CollectionName:   "info",
		HashKey:          []string{"name"},
		IndexType:        define.IndexTypeSingle,
	},
	"UserPostBox": {
		DatabaseLocation: define.DBCommon,
		DatabaseName:     "User",
		CollectionName:   "PostBox",
		HashKey:          []string{},
		IndexType:        define.IndexTypeSingle,
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
