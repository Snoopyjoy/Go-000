package api

import "net/http"

func OK(w http.ResponseWriter, data []byte) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(200)
	_, _ = w.Write(data)
}

func NotFound(w http.ResponseWriter, data []byte) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(404)
	_, _ = w.Write(data)
}

func PageNotFound(w http.ResponseWriter) {
	NotFound(w, []byte("page not found"))
}

func ServeError(w http.ResponseWriter, data []byte) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(500)
	_, _ = w.Write(data)
}

func ClientError(w http.ResponseWriter, data []byte) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(400)
	_, _ = w.Write(data)
}
