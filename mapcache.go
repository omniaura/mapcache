package mapcache

import "sync"

type MapCache[K comparable, V any] struct {
	m  map[K]V
	mu sync.RWMutex
}
