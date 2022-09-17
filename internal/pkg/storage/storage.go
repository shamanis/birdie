package storage

import (
	"errors"
	"github.com/shamanis/birdie/internal/pkg/logging"
	"sync"
)

const (
	Count = 256 // Don't edit!
)

var (
	logger        = logging.New()
	storages      map[int]*Storage
	NotFoundError = errors.New("not found")
)

func init() {
	storages = make(map[int]*Storage, Count)
	for i := 0; i < Count; i++ {
		storages[i] = &Storage{}
		storages[i].mu = &sync.RWMutex{}
		storages[i].s = make(map[string]*Entry, 1024)
	}

	go Clear()
}

type Storage struct {
	mu *sync.RWMutex
	s  map[string]*Entry
}

func (s *Storage) Store(entry *Entry) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.s[entry.Key] = entry
}

func (s *Storage) Load(key string) (*Entry, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	value, ok := s.s[key]
	if !ok {
		return nil, NotFoundError
	}

	if value.HasExpired() {
		return nil, NotFoundError
	}

	return value, nil
}

func (s *Storage) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.s, key)
}

func (s *Storage) Clear() {
	for k, v := range s.s {
		if v.HasExpired() {
			s.Delete(k)
		}
	}
}

func (s *Storage) AllEntries() []*Entry {
	var a []*Entry
	for _, v := range s.s {
		a = append(a, v)
	}
	return a
}

func getStorage(key string) *Storage {
	h := getHash(key)
	return storages[int(h[0])]
}

// Store Entry
func Store(e *Entry) {
	s := getStorage(e.Key)
	s.Store(e)
}

// Load Entry
func Load(key string) (*Entry, error) {
	s := getStorage(key)
	return s.Load(key)
}

// BulkStore store array Entry
func BulkStore(entries []*Entry) {
	for _, e := range entries {
		Store(e)
	}
}

// Delete Entry
func Delete(key string) {
	s := getStorage(key)
	s.Delete(key)
}
