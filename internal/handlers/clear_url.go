package handlers

import (
	"net/http"

	"github.com/kKar1503/url-shortener/internal/urlmapper"
)

func ClearURL(w http.ResponseWriter, r *http.Request, mapper urlmapper.URLMapper) {
	key := r.URL.Path[1:]
	mapper.Remove(key)
	w.WriteHeader(http.StatusOK)
}
