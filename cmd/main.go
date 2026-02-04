package main

import (
	"fmt"
	"sync"
)

type GitiKV struct {
	mu   sync.Mutex
	Giti map[string]any
}

func NewGitiKV() *GitiKV {
	return &GitiKV{
		Giti: map[string]any{},
	}
}

func (gkv *GitiKV) Read(key string) any {
	gkv.mu.Lock()
	val := gkv.Giti[key]
	gkv.mu.Unlock()

	return val
}

func (gkv *GitiKV) Write(key string, val any) {
	gkv.mu.Lock()
	gkv.Giti[key] = val
	gkv.mu.Unlock()
}

func main() {
	gkv := NewGitiKV()
	gkv.Write("workerID", 1)
	gkv.Write("Price", 1)

	val := gkv.Read("Price")
	fmt.Printf("Price: %v\n", val)
}
