package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/kKar1503/url-shortener/internal/handlers"
	"github.com/kKar1503/url-shortener/internal/urlmapper"
)

func main() {
	port := os.Getenv("PORT")

	urlMapper := urlmapper.NewBasicURLMapper()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.Redirect(w, r, urlMapper)
		case http.MethodPost:
			handlers.CreateURL(w, r, urlMapper)
		case http.MethodDelete:
			handlers.ClearURL(w, r, urlMapper)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Printf("URL Shortener is running on :%s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil && err != http.ErrServerClosed {
		fmt.Printf("Error: %v\n", err)
	}
}
