package memcache551

type Memcache struct{}

var memcacheInstance *Memcache

func Load() *Memcache {
	if memcacheInstance != nil {
		return memcacheInstance
	}

	memcacheInstance = &Memcache{}

	return memcacheInstance
}
