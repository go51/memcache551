package memcache551_test

import (
	"github.com/go51/memcache551"
	"testing"
)

func TestLoad(t *testing.T) {
	config := memcache551.Config{
		Host:    "localhost:11211",
		Prefix:  "gorai",
		Expires: 3600,
	}

	m1 := memcache551.New(&config)
	m2 := memcache551.New(&config)

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
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = memcache551.New(&config)
	}
}
