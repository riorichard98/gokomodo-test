package config

type DefaultConfig struct {
	Apps Apps `json:"apps"`
}

type Apps struct {
	Name     string `json:"name"`
	Version  string `json:"version"`
	Address  string `json:"address"`
	HttpPort string `json:"httpPort"`
}

type MongoConfig struct {
	URI         string `json:"uri"`
	MaxPoolSize uint64 `json:"max_pool_size"`
	MinPoolSize uint64 `json:"min_pool_size"`
	Timeout     int    `json:"timeout"`
}

type DB struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
	// Schema      string `json:"schema"`
	Port        int `json:"port"`
	Timeout     int `json:"timeout"`
	MaxPoolSize int `json:"maxPool"`
	MinPoolSize int `json:"minPoolSize"`
	// DebugMode   bool   `json:"debugMode"`
}
