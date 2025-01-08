package database

import (
	"GoServer/config/define"
	"GoServer/config/structure"
)

// "CollectionName": {DatabaseLocation, "DatabaseName", "CollectionName", []string{"HashKey"}, define.IndexTypeSingle},
var MongoCollectionInfos = map[string]structure.MongoCollectionInfo{
	"UserInfo": {define.DBCommon, "User", "info", []string{"email"}, define.IndexTypeSingle},
}

var MySQLCollectionInfos = map[string]structure.MySQLCollectionInfo{
	"UserTable": {"user_table"},
	"Menu":      {"menu"},
}
