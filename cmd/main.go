package main

import (
	"encoding/binary"
	"fmt"
	"sync"
)

type GitiKV struct {
	mu   sync.Mutex
	Giti map[string][]byte
}

func NewGitiKV() *GitiKV {
	return &GitiKV{
		Giti: map[string][]byte{},
	}
}

func (gkv *GitiKV) Read(key string) ([]byte, bool) {
	gkv.mu.Lock()
	defer gkv.mu.Unlock()
	val, ok := gkv.Giti[key]

	return val, ok
}

func (gkv *GitiKV) Write(key string, val []byte) {
	gkv.mu.Lock()
	gkv.Giti[key] = val
	gkv.mu.Unlock()
}

func encodeInt32(i int32) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(i))
	return b
}

func decodeInt32(b []byte) int32 {
	return int32(binary.BigEndian.Uint32(b))
}

func main() {
	gkv := NewGitiKV()

	gkv.Write("userID", encodeInt32(2))
	val, ok := gkv.Read("userID")
	if !ok {
		panic("key not found")
	}

	fmt.Printf("userID: %d\n", decodeInt32(val))
}
