package storage

import (
	"crypto/md5"
)

func getHash(key string) []byte {
	h := md5.New()
	_, err := h.Write([]byte(key))

	if err != nil {
		logger.Panicf("fail hashing key: %s", err)
	}

	return h.Sum(nil)
}
