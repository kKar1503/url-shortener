package handlers

import (
	"net/http"

	"github.com/kKar1503/url-shortener/internal/urlmapper"
)

func Redirect(w http.ResponseWriter, r *http.Request, mapper urlmapper.URLMapper) {
	key := r.URL.Path[1:]
	url, ok := mapper.Get(key)
	if !ok {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, url, http.StatusFound)
}
