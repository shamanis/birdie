package storage

import "time"

type Entry struct {
	Key   string
	Value []byte
	time  time.Time
	ttl   time.Duration
}

func NewEntry(key string, value []byte, ttl time.Duration) *Entry {
	return &Entry{
		Key:   key,
		Value: value,
		time:  time.Now().UTC(),
		ttl:   ttl,
	}
}

func (e *Entry) HasExpired() bool {
	return e.time.Add(e.ttl).Before(time.Now().UTC())
}

func (e *Entry) Store() {
	getStorageByKey(e.Key).Store(e.Key, e)
}
