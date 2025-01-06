package structure

type Server struct {
	Mode    string
	Port    string
	Version string
}

type Log struct {
	Level     string
	Fpath     string
	MaxSize   int
	MaxAge    int
	MaxBackup int
	Compress  bool
}

type ENV struct {
	Target string
}

type SecretKey struct {
	APPSecretKey    string
	APISecretKey    string
	CommonSecretKey string
}
