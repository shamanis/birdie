package storage

import (
	"crypto/md5"
)

func getHash(key string) []byte {
	h := md5.New()
	_, err := h.Write([]byte(key))

	if err != nil {
		panic(err)
	}

	return h.Sum(nil)
}
