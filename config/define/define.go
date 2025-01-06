package define

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
