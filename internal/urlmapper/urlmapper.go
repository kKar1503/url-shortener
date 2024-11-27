package urlmapper

type URLMapper interface {
	Add(url string) string
	AddCustom(key, url string) bool
	Get(key string) (string, bool)
	Remove(key string)
}
