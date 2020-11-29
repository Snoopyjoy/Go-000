package api

import "net/http"

type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

type IndexHandler struct{}

func NewIndexHandler() Handler {
	return &IndexHandler{}
}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/" || req.URL.Path == "" {
		OK(w, []byte("hello index"))
		return
	}
	PageNotFound(w)
}
