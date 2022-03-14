package pool

import (
	"sync"
	"testing"
)

func BenchmarkOurPool(b *testing.B) {
	p := NewPool(func() (itf interface{}) { itf = 1488; return }, func(itf interface{}) {}, 10000000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		o := p.Get()
		p.Put(o)
	}
}

func BenchmarkStandartPool(b *testing.B) {
	p := sync.Pool{ New: func() (itf interface{}) {itf = 1488; return}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		o := p.Get()
		p.Put(o)
	}
}
