package storage

import "time"

// Clear clearing storage from expired values
func Clear() {
	t := time.NewTicker(600 * time.Second)
	defer t.Stop()
	for range t.C {
		for _, s := range storages {
			s.Clear()
		}
	}
}
