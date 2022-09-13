package storage

import (
	"fmt"
	"testing"
	"time"
)

func TestStore(t *testing.T) {
	entry := NewEntry("key1", []byte("value1"), 5*time.Second)

	Store(entry)

	result, err := Load("key1")
	if err != nil {
		t.Error(err)
	}

	if result != entry {
		t.Error("result != value")
	}
}

func BenchmarkStore(b *testing.B) {
	for i := 0; i < b.N; i++ {
		entry := NewEntry(fmt.Sprintf("key_%d", i), []byte(fmt.Sprintf("value_%d", i)), 5*time.Second)
		Store(entry)
	}
}

func BenchmarkLoad(b *testing.B) {
	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key_%d", i)
		Load(key)
	}
}
