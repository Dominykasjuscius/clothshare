package webconfig

type Config struct {
	Http    HTTPServer    `toml:"http"`
	Mongodb MongoDatabase `toml:"mongodb"`
}

type HTTPServer struct {
	Adress string `toml:"adress"`
}

type MongoDatabase struct {
	URI      string `toml:"uri"`
	Database string `toml:"database"`
	Poolsize int    `toml:"poolsize"`
}
