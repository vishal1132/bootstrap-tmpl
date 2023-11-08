package http

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"go.uber.org/zap"
	"{{ module_name }}/observability/log"
)

const (
	defaultPort = 8080
)

func GetRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)
	router.Use(render.SetContentType(render.ContentTypeJSON))
	router.Use(SetApplicationTypeJSON)
	router.Use(CORSMiddleware)
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println(r.Header)
			log.Debug(r.Context(), "middleware", zap.String("path", r.URL.Path), zap.String("host", extractHost(r)), zap.String("url.host", r.URL.Host))
			next.ServeHTTP(w, r.WithContext(r.Context()))
		})
	})
	return router
}

func extractHost(r *http.Request) (host string) {
	host = r.Header.Get("X-Forwarded-Host")
	if host != "" {
		return
	}

	// RFC 7239
	host = r.Header.Get("Forwarded")
	_, _, host = parseForwarded(host)
	if host != "" {
		return
	}

	// if all else fails fall back to request host
	host = r.Host
	return
}

func parseForwarded(forwarded string) (addr, proto, host string) {
	if forwarded == "" {
		return
	}
	for _, forwardedPair := range strings.Split(forwarded, ";") {
		if tv := strings.SplitN(forwardedPair, "=", 2); len(tv) == 2 {
			token, value := tv[0], tv[1]
			token = strings.TrimSpace(token)
			value = strings.TrimSpace(strings.Trim(value, `"`))
			switch strings.ToLower(token) {
			case "for":
				addr = value
			case "proto":
				proto = value
			case "host":
				host = value
			}

		}
	}
	return
}

func InitServer(r *chi.Mux, writeTimeout, readTimeout int64) *http.Server {
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", defaultPort),
		Handler:      r,
		WriteTimeout: time.Duration(writeTimeout) * time.Millisecond,
		ReadTimeout:  time.Duration(readTimeout) * time.Millisecond,
	}

	return server
}

func SetApplicationTypeJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8;")
		next.ServeHTTP(w, r)
	})
}

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			return
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		next.ServeHTTP(w, r)
	})
}
