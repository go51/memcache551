package memcache551_test

import (
	"github.com/go51/memcache551"
	"github.com/go51/secure551"
	"testing"
)

func TestLoad(t *testing.T) {
	config := memcache551.Config{
		Host:    "localhost:11211",
		Prefix:  "gorai",
		Expires: 3600,
	}
	sid := secure551.Hash()

	m1 := memcache551.New(&config, sid)
	m2 := memcache551.New(&config, sid)

	if m1 == nil {
		t.Errorf("インスタンス生成に失敗しました。")
	}
	if m2 == nil {
		t.Errorf("インスタンス生成に失敗しました。")
	}
	if &m1 == &m2 {
		t.Errorf("インスタンス生成に失敗しました。\n[%p], [%p]", &m1, &m2)
	}
}

func BenchmarkLoad(b *testing.B) {
	config := memcache551.Config{
		Host:    "localhost:11211",
		Prefix:  "gorai",
		Expires: 3600,
	}
	sid := secure551.Hash()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = memcache551.New(&config, sid)
	}
}

func TestSet(t *testing.T) {
	config := memcache551.Config{
		Host:    "localhost:11211",
		Prefix:  "gorai",
		Expires: 3600,
	}
	sid := secure551.Hash()

	m := memcache551.New(&config, sid)
	m.Set("test", "test_string")
}

func BenchmarkSet(b *testing.B) {
	config := memcache551.Config{
		Host:    "localhost:11211",
		Prefix:  "gorai",
		Expires: 3600,
	}
	sid := secure551.Hash()
	m := memcache551.New(&config, sid)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Set("test", "test_string")
	}

}
