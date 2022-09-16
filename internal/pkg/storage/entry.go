package storage

import "time"

type ExpiredFunction func(t time.Time, ttl time.Duration) bool

type Ttl interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

type Entry struct {
	Key         string
	Value       []byte
	time        time.Time
	ttl         time.Duration
	expiredFunc ExpiredFunction
}

func hasExpired(t time.Time, ttl time.Duration) bool {
	return t.Add(ttl).Before(time.Now().UTC())
}

func NewEntry[T Ttl](key string, value []byte, ttl T) *Entry {
	return &Entry{
		Key:         key,
		Value:       value,
		time:        time.Now().UTC(),
		ttl:         time.Duration(ttl) * time.Second,
		expiredFunc: hasExpired,
	}
}

func (e *Entry) HasExpired() bool {
	return e.expiredFunc(e.time, e.ttl)
}
