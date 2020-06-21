package main

import "net/http"

type Middleware struct {
	config *Config
}

func NewMiddleware(config *Config) *Middleware {
	return &Middleware{
		config: config,
	}
}

// LogMiddleware для логирования всех входящих запросов
func (m *Middleware) AddLogger() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			m.config.Log.Log.Infof("ip=%s, method=%s, route=%s", r.RemoteAddr, r.Method, r.RequestURI)
			next.ServeHTTP(w, r)
		})
	}
}

// HeadersMiddleware для добавления необходимых заголовков
func (m *Middleware) AddHeaders() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	}

}
