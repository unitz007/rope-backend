package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

type MuxMultiplexer struct {
	router *mux.Router
}

func (mux *MuxMultiplexer) Name() string {
	return "Mux Router"
}

func NewMuxMultiplexer(mux *mux.Router) Router {
	return &MuxMultiplexer{router: mux}
}

func (mux *MuxMultiplexer) Get(path string, handlerFunc func(w http.ResponseWriter, r *http.Request)) {
	mux.router.HandleFunc(path, handlerFunc).Methods(http.MethodGet)
}

func (mux *MuxMultiplexer) Serve(addr string) error {
	return http.ListenAndServe(addr, mux.router)
}

func (mux *MuxMultiplexer) Post(path string, handlerFunc func(w http.ResponseWriter, r *http.Request)) {
	mux.router.HandleFunc(path, handlerFunc).Methods(http.MethodPost)
}
