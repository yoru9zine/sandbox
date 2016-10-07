package hashmap

import (
	"fmt"
	"hash/crc32"
)

type entry struct {
	Key   string
	Value interface{}
}

// HashMap is hash table
type HashMap struct {
	buckets []*LinkedList
	size    int
}

// NewHashMap returns new hash instance
func NewHashMap(n int) *HashMap {
	return &HashMap{
		buckets: make([]*LinkedList, n),
		size:    n,
	}
}

// Set sets key and value
func (h *HashMap) Set(key string, value interface{}) {
	idx := h.index(key)
	bucket := h.buckets[idx]

	if bucket == nil {
		h.buckets[idx] = &LinkedList{Value: &entry{key, value}}
		return
	}

	p := bucket
	n := 0
	for ; p.Next != nil; p = p.Next {
		if p.Value.(*entry).Key == key {
			p.Value = &entry{key, value}
			return
		}
		n++
	}
	if n == 10 {
		h.rehash()
	}
	p.Next = &LinkedList{Value: &entry{key, value}}
}

// Get returns value of key
func (h *HashMap) Get(key string) interface{} {
	bucket := h.buckets[h.index(key)]
	if bucket == nil {
		return nil
	}
	for e := bucket; e != nil; e = e.Next {
		entry := e.Value.(*entry)
		if entry.Key == key {
			return entry.Value
		}
	}
	return nil
}

func (h *HashMap) rehash() {
	fmt.Printf("rehash: %d\n", h.size*2)
	hh := NewHashMap(h.size * 2)
	for _, b := range h.buckets {
		for e := b; e != nil; e = e.Next {
			v := e.Value.(*entry)
			hh.Set(v.Key, v.Value)
		}
	}
	h.buckets = hh.buckets
	h.size = hh.size
}

func (h *HashMap) index(key string) int {
	return int(crc32.Checksum([]byte(key), nil) % uint32(h.size))
}
