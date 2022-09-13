package storage

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/shamanis/birdie/internal/pkg/logging"
	"github.com/sirupsen/logrus"
	"sync"
)

const (
	Count = 256 // Don't edit!
)

var (
	storages      map[int]*sync.Map
	logger        = logging.New()
	NotFoundError = errors.New("not found")
	TypeCastError = errors.New("fail cast to []byte")
)

func init() {
	storages = make(map[int]*sync.Map, Count)
	for i := 0; i < Count; i++ {
		storages[i] = &sync.Map{}
	}
}

func getStorageByKey(key string) *sync.Map {
	hasher := sha256.New()
	hasher.Write([]byte(key))
	hash := hasher.Sum(nil)

	return storages[int(hash[0])]
}

func Store(entry *Entry) {
	logger.WithFields(
		logrus.Fields{"key": entry.Key, "value": fmt.Sprintf("%+v", entry)},
	).Debug("Store value")
	getStorageByKey(entry.Key).Store(entry.Key, entry)
}

func Load(key string) (*Entry, error) {
	storage := getStorageByKey(key)
	v, ok := storage.Load(key)
	if !ok {
		logger.WithFields(logrus.Fields{"key": key}).Debug("Not found entry")
		return nil, NotFoundError
	}

	value, ok := v.(*Entry)
	if !ok {
		logger.WithFields(
			logrus.Fields{"key": key, "raw_value": fmt.Sprintf("%+v", v)},
		).Error("Type cast error")
		return nil, TypeCastError
	}

	return value, nil
}
