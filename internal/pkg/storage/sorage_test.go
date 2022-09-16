package storage

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func entryExpiredMock(t time.Time, ttl time.Duration) bool {
	return true
}

func TestStoreLoad(t *testing.T) {
	entry := NewEntry("key1", []byte("value1"), 5)
	Store(entry)

	result, err := Load("key1")
	if err != nil {
		t.Error(err)
	}

	if result != entry {
		t.Error("result != value")
	}
}

func TestDelete(t *testing.T) {
	entry := NewEntry("key1", []byte("value1"), 5)
	Store(entry)

	Delete("key1")
	_, err := Load("key1")
	if !errors.Is(err, NotFoundError) {
		t.Error(err)
	}
}

func TestExpired(t *testing.T) {
	entry := NewEntry("key1", []byte("value1"), 5)
	Store(entry)
	entry.expiredFunc = entryExpiredMock

	_, err := Load("key1")
	if !errors.Is(err, NotFoundError) {
		t.Error(err)
	}
}

func TestClear(t *testing.T) {
	entry := NewEntry("key1", []byte("value1"), 5)
	Store(entry)
	entry.expiredFunc = entryExpiredMock
	s := getStorage("key1")
	s.Clear()

	_, err := Load("key1")
	if !errors.Is(err, NotFoundError) {
		t.Error(err)
	}
}

func BenchmarkHash(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		getHash(fmt.Sprintf("key_%d", i))
	}
}

func BenchmarkStore(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		entry := NewEntry(fmt.Sprintf("key_%d", i), []byte(fmt.Sprintf("value_%d", i)), 5)
		Store(entry)
	}
}

func BenchmarkLoad(b *testing.B) {
	for i := 0; i < b.N; i++ {
		entry := NewEntry(fmt.Sprintf("key_%d", i), []byte(fmt.Sprintf("value_%d", i)), 5)
		Store(entry)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Load(fmt.Sprintf("key_%d", i))
	}
}

func makeEntries(count, N int) [][]*Entry {
	var bulks [][]*Entry
	for i := 0; i < N; i++ {
		var bulk []*Entry
		for j := 0; j < count; j++ {
			entry := NewEntry(fmt.Sprintf("key_%d_%d", i, j), []byte(fmt.Sprintf("value_%d_%d", i, j)), 5)
			bulk = append(bulk, entry)
		}
		bulks = append(bulks, bulk)
	}

	return bulks
}

func BenchmarkBulkStore(b *testing.B) {
	var count = []int{1000, 10000, 100000}
	for _, c := range count {
		b.Run(fmt.Sprintf("count_%d", c), func(b *testing.B) {
			bulks := makeEntries(c, b.N)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				BulkStore(bulks[i])
			}
		})
	}
}

func BenchmarkBulkStore_Parallel(b *testing.B) {
	var count = []int{1000, 10000, 100000}
	for _, c := range count {
		b.Run(fmt.Sprintf("count_%d", c), func(b *testing.B) {
			bulks := makeEntries(c, b.N)
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				i := 0
				for pb.Next() {
					BulkStore(bulks[i])
					i++
				}
			})
		})
	}
}
