package handlers

import (
	"net/http"
	"url-shortener/internal/shortener"
)

func HandleShorten(svc *shortener.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rawURL := r.PathValue("url")
		if rawURL == "" {
			http.Error(w, "URL is required", http.StatusBadRequest)
			return
		}
		code, err := svc.Shorten(rawURL)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Write([]byte("Your short code is: " + code + "\n"))
		fullShortURL := "http://" + r.Host + "/" + code
		w.Write([]byte("Your short URL is:" + fullShortURL))
	}
}

func HandleRedirect(svc *shortener.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := r.PathValue("code")
		if code == "" {
			http.Error(w, "short code is required", http.StatusBadRequest)
			return
		}

		longURL, exists := svc.Resolve(code)
		if !exists {
			http.Error(w, "short link not found", http.StatusNotFound)
			return
		}
		http.Redirect(w, r, longURL, http.StatusFound)
	}
}
