package structure

// <<key - env>>
type MongoMap map[string]MongoInfoMap
type MySQLMap map[string]MySQLConfig
type RedisMap map[string]RedisConfig

// <key - common|api|wallet>>
type MongoInfoMap map[string]MongoConfig

type MongoConfig struct {
	Host string
	User string
	Pass string
}

type MySQLConfig struct {
	Host string
	User string
	Pass string
	Port string
	DB   string
}

type RedisConfig struct {
	Host string
	Port string
	Pass string
}

type MongoCollectionInfo struct {
	DatabaseLocation int
	DatabaseName     string
	CollectionName   string
	HashKey          []string
	IndexType        int
}

type MySQLCollectionInfo struct {
	TableName string
}
