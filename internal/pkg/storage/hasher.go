package storage

import (
	"hash/adler32"
)

func getHash(key string) []byte {
	h := adler32.New()
	_, err := h.Write([]byte(key))

	if err != nil {
		panic(err)
	}

	return h.Sum(nil)
}
