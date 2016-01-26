package memcache551

import (
	"encoding/json"
	"github.com/bradfitz/gomemcache/memcache"
)

type Memcache struct {
	config *Config
	client *memcache.Client
	sid    string
}

type Config struct {
	Host    string `json:"host"`
	Prefix  string `json:"prefix"`
	Expires int32  `json:"expires"`
}

func New(config *Config, sid string) *Memcache {
	m := &Memcache{
		config: config,
		sid:    sid,
		client: memcache.New(config.Host),
	}

	return m
}

func (m *Memcache) Set(name string, value interface{}) {
	val, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}

	err = m.client.Set(&memcache.Item{
		Key:        m.config.Prefix + ":" + m.sid + ":" + name,
		Value:      val,
		Expiration: m.config.Expires,
	})
	if err != nil {
		panic(err)
	}

}
