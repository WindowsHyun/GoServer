package structure

// <<key - env>>
type MongoMap map[string]MongoInfoMap

// <key - common|api|wallet>>
type MongoInfoMap map[string]MongoConfig

type MongoConfig struct {
	Host string
	User string
	Pass string
}

type CollectionInfo struct {
	DatabaseLocation int
	DatabaseName     string
	CollectionName   string
	HashKey          []string
	IndexType        int
}
