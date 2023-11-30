package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"
)

const (
	reqLimit = 500 // no of request made
)

func TestGetData(t *testing.T) {
	store := NewDbStore()
	h := httptest.NewServer(http.HandlerFunc(store.handleGetData))

	// to count and wait for all goroutines
	wg := &sync.WaitGroup{}

	for i := 0; i < reqLimit; i++ {
		wg.Add(1)
		// make concurrent requests
		go func(i int) {
			id := i%100 + 1
			url := fmt.Sprintf("%s/?id=%d", h.URL, id)

			response, err := http.Get(url)
			if err != nil {
				t.Error(err)
			}

			data := &Data{}
			if err = json.NewDecoder(response.Body).Decode(data); err != nil {
				t.Error(err)
			}
			wg.Done()
		}(i)
		// add sleep to verify request
		time.Sleep(time.Millisecond * 1)
	}
	wg.Wait()

	fmt.Println("total no of request made", reqLimit)
	fmt.Println("db query count", store.queryCount)
	fmt.Println("cache hit count", store.cacheHit)
}
