package urlmapper

import (
	"math/rand"
	"time"
)

// This gives us a total of 62^6 = 56,800,235,584 possible keys
// Excluding the custom keys, which means there could be more
const (
	charset     = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	randKeySize = 6
)

type BasicURLMapper struct {
	mapper map[string]string
}

func NewBasicURLMapper() *BasicURLMapper {
	return &BasicURLMapper{
		mapper: make(map[string]string),
	}
}

func (bum *BasicURLMapper) newKey() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]byte, randKeySize)
	for {
		for i := range result {
			result[i] = charset[r.Intn(len(charset))]
		}
		key := string(result)
		if _, ok := bum.mapper[key]; !ok {
			return key
		}
	}
}

func (bum *BasicURLMapper) Add(url string) string {
	key := bum.newKey()

	bum.mapper[key] = url
	return key
}

func (bum *BasicURLMapper) AddCustom(key, url string) bool {
	if _, ok := bum.mapper[key]; ok {
		return false
	}
	bum.mapper[key] = url
	return true
}

func (bum *BasicURLMapper) Get(key string) (string, bool) {
	url, ok := bum.mapper[key]
	return url, ok
}

func (bum *BasicURLMapper) Remove(key string) {
	delete(bum.mapper, key)
}
