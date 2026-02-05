package store

import "sync"

type Value struct {
}

type GitiKV struct {
	mu   sync.Mutex
	Giti map[string][]byte
}

func NewGitiKV() *GitiKV {
	return &GitiKV{
		Giti: map[string][]byte{},
	}
}
