package handlers

import (
	"net/http"

	"github.com/kKar1503/url-shortener/internal/urlmapper"
)

func CreateURL(w http.ResponseWriter, r *http.Request, mapper urlmapper.URLMapper) {
	url := r.FormValue("url")
	if url == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	customKey := r.FormValue("custom")
	if customKey != "" {
		if !mapper.AddCustom(customKey, url) {
			http.Error(w, "Custom key already exists", http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(customKey)); err != nil {
			http.Error(w, "Error writing response", http.StatusInternalServerError)
			return
		}
		return
	}

	key := mapper.Add(url)
	if _, err := w.Write([]byte(key)); err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
	}
}
