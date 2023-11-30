package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const (
	LIMIT = 100 // data item count
)

type Data struct {
	Id   int
	Name string
}

type DbStore struct {
	db         map[int]*Data
	cache      map[int]*Data
	queryCount int
	cacheHit   int
}

func NewDbStore() *DbStore {
	db := make(map[int]*Data)
	for i := 0; i < LIMIT; i++ {
		db[i+1] = &Data{
			Id:   i + 1,
			Name: fmt.Sprintf("data_%d", i+1),
		}
	}
	return &DbStore{
		db:    db,
		cache: make(map[int]*Data),
	}
}

func (s *DbStore) getFromCache(id int) (*Data, bool) {
	data, ok := s.cache[id]
	return data, ok
}

func (s *DbStore) handleGetData(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	dataId, _ := strconv.Atoi(id)

	// check cache first
	data, ok := s.getFromCache(dataId)
	if ok {
		s.cacheHit++
		json.NewEncoder(w).Encode(data)
		return
	}

	// query db
	data, ok = s.db[dataId]
	if !ok {
		panic("data not found")
	}
	s.queryCount++

	// update cache before returning
	s.cache[dataId] = data
	json.NewEncoder(w).Encode(data)
}

func main() {}
