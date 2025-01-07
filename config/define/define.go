package define

type contextKey string

var ContextLoggerPrint contextKey = "LoggerPrint"
var ContextUserRealIP contextKey = "UserRealIP"

const (
	IndexTypeSingle = iota
	IndexTypeCompound
)

const (
	DBApp = iota
	DBApi
	DBCommon
)

// Mongo TargetDB
const (
	MongoApp    = "app"
	MongoApi    = "api"
	MongoCommon = "common"
)

const (
	FolderPermission = 0755
	FilePermission   = 0660
	PanicLine        = "-------------------------------------------------------------------------------"
)
