package memcache551

type Memcache struct {
	config *Config
}

type Config struct {
	Host    string `json:"host"`
	Prefix  string `json:"prefix"`
	Expires int32  `json:"expires"`
}

func New(config *Config) *Memcache {
	return &Memcache{
		config: config,
	}
}
