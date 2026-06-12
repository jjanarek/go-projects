package main

import "sync"

type Store struct {
	mu     sync.RWMutex
	notes  map[int]Note
	nextID int
}

func NewStore() *Store {
	return &Store{
		notes:  make(map[int]Note),
		nextID: 1,
	}
}
