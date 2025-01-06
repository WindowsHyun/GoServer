package database

import (
	"GoServer/config/define"
	"GoServer/config/structure"
)

// "CollectionName": {DatabaseLocation, "DatabaseName", "CollectionName", []string{"HashKey"}, define.IndexTypeSingle},
var CollectionInfos = map[string]structure.CollectionInfo{
	"UserInfo": {define.DBCommon, "User", "info", []string{"email"}, define.IndexTypeSingle},
}
