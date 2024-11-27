package urlmapper_test

import (
	"testing"

	"github.com/kKar1503/url-shortener/internal/urlmapper"
)

type NamedMappers struct {
	name   string
	mapper func() urlmapper.URLMapper
}

func getImplementedMappers() []NamedMappers {
	return []NamedMappers{
	}
}

func TestAdd(t *testing.T) {
	namedMappers := getImplementedMappers()
	for _, namedMapper := range namedMappers {
		mapper := namedMapper.mapper()
		key := mapper.Add("http://www.google.com")
		if key == "" {
			t.Errorf("%s: Expected key to be non-empty", namedMapper.name)
		}

		url, ok := mapper.Get(key)
		if !ok {
			t.Errorf("%s: Expected key to be found", namedMapper.name)
		}

		if url != "http://www.google.com" {
			t.Errorf("%s: Expected url to be http://www.google.com", namedMapper.name)
		}
	}
}

func TestAddDuplicate(t *testing.T) {
	namedMappers := getImplementedMappers()
	for _, namedMapper := range namedMappers {
		mapper := namedMapper.mapper()
		key1 := mapper.Add("http://www.google.com")
		key2 := mapper.Add("http://www.google.com")

		if key1 == key2 {
			t.Errorf("%s: Expected keys to be different", namedMapper.name)
		}

		url, ok := mapper.Get(key1)
		if !ok {
			t.Errorf("%s: Expected key to be found", namedMapper.name)
		}

		if url != "http://www.google.com" {
			t.Errorf("%s: Expected url to be http://www.google.com", namedMapper.name)
		}

		url, ok = mapper.Get(key2)
		if !ok {
			t.Errorf("%s: Expected key to be found", namedMapper.name)
		}

		if url != "http://www.google.com" {
			t.Errorf("%s: Expected url to be http://www.google.com", namedMapper.name)
		}
	}
}

func TestAddCustom(t *testing.T) {
	namedMappers := getImplementedMappers()
	for _, namedMapper := range namedMappers {
		mapper := namedMapper.mapper()
		if !mapper.AddCustom("abc", "http://www.google.com") {
			t.Errorf("%s: Expected key to be added", namedMapper.name)
		}

		if mapper.AddCustom("abc", "http://www.google.com") {
			t.Errorf("%s: Expected key to not be added", namedMapper.name)
		}

		url, ok := mapper.Get("abc")
		if !ok {
			t.Errorf("%s: Expected key to be found", namedMapper.name)
		}

		if url != "http://www.google.com" {
			t.Errorf("%s: Expected url to be http://www.google.com", namedMapper.name)
		}
	}
}

func TestAddCustomCollision(t *testing.T) {
	namedMappers := getImplementedMappers()
	for _, namedMapper := range namedMappers {
		mapper := namedMapper.mapper()
		key := mapper.Add("http://www.google.com")

		if mapper.AddCustom(key, "http://www.google.com") {
			t.Errorf("%s: Expected key to not be added", namedMapper.name)
		}
	}
}

func TestGet(t *testing.T) {
	namedMappers := getImplementedMappers()
	for _, namedMapper := range namedMappers {
		mapper := namedMapper.mapper()
		key := mapper.Add("http://www.google.com")

		url, ok := mapper.Get(key)
		if !ok {
			t.Errorf("%s: Expected key to be found", namedMapper.name)
		}

		if url != "http://www.google.com" {
			t.Errorf("%s: Expected url to be http://www.google.com", namedMapper.name)
		}
	}
}

func TestGetNotFound(t *testing.T) {
	namedMappers := getImplementedMappers()
	for _, namedMapper := range namedMappers {
		mapper := namedMapper.mapper()
		_, ok := mapper.Get("abc")
		if ok {
			t.Errorf("%s: Expected key to not be found", namedMapper.name)
		}
	}
}

func TestRemove(t *testing.T) {
	namedMappers := getImplementedMappers()
	for _, namedMapper := range namedMappers {
		mapper := namedMapper.mapper()
		key := mapper.Add("http://www.google.com")

		mapper.Remove(key)
		_, ok := mapper.Get(key)
		if ok {
			t.Errorf("%s: Expected key to not be found", namedMapper.name)
		}
	}
}
