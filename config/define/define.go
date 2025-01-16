package define

type contextKey string

var ContextLoggerPrint contextKey = "LoggerPrint"
var ContextUserRealIP contextKey = "UserRealIP"

const (
	IndexTypeSingle = iota
	IndexTypeCompound
)

// Mongo TargetDB
const (
	MongoApp    = "app"
	MongoApi    = "api"
	MongoCommon = "common"
)

// Redis DB
const (
	RedisUserDB    = "0"
	RedisRankingDB = "1"
	RedisItemDB    = "2"
	RedisGuildDB   = "3"
	RedisLogDB     = "4"
	RedisEventDB   = "5"
	RedisCacheDB   = "6"
	RedisTempDB    = "7"
	RedisBackupDB  = "8"
	RedisSessionDB = "9"
)

const (
	FolderPermission = 0755
	FilePermission   = 0660
	PanicLine        = "-------------------------------------------------------------------------------"
)
