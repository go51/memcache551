package memcache551

import (
	"crypto/md5"
	"encoding/json"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/go51/string551"
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
		Key:        generateKey(m.config.Prefix, m.sid, name),
		Value:      val,
		Expiration: m.config.Expires,
	})
	if err != nil {
		panic(err)
	}

}

func (m *Memcache) Get(name string) interface{} {
	item, err := m.client.Get(generateKey(m.config.Prefix, m.sid, name))
	if err != nil {
		return nil
	}

	var obj interface{} = nil

	json.Unmarshal(item.Value, &obj)

	return obj
}

func generateKey(prefix, sid, name string) string {
	hash := md5.New()
	b := string551.StringToBytes(prefix + ":" + sid + ":" + name)
	hash.Write(b)
	return string551.HexBytesToString(hash.Sum(nil))

}
