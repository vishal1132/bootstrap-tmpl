package http

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

const (
	DefaultPageSize = 10
	DefaultPage     = 1
	PageParameter   = "page"
	SizeParameter   = "size"
)

func ExtractHost(r *http.Request) (host string) {
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

func Error(w http.ResponseWriter, r *http.Request, statusCode int, err error) {
	w.WriteHeader(statusCode)
	res := map[string]string{
		"__debugstring__": err.Error(),
	}
	render.Respond(w, r, res)
}

func ResponseOK(w http.ResponseWriter, r *http.Request, data interface{}) {
	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, data)
}

func PaginatedResponse(w http.ResponseWriter, r *http.Request, data interface{}) {
	w.WriteHeader(http.StatusOK)
	render.Respond(w, r, data)
}

func ExtractToken(r *http.Request) (string, error) {
	token := r.Header.Get("Authorization")
	if token == "" || len(token) < 7 {
		return "", errors.New("no token found in request")
	}
	token = token[7:]
	return token, nil
}

func GetPathParam(r *http.Request, name string) string {
	return chi.URLParam(r, name)
}
