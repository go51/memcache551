package memcache551_test

import (
	"github.com/go51/memcache551"
	"testing"
)

func TestLoad(t *testing.T) {
	m1 := memcache551.Load()
	m2 := memcache551.Load()

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
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = memcache551.Load()
	}
}
