package router

import (
	"amusingx.fit/amusingx/services/amusingapiserv/api/pong"
	"net/http"
)

func Register(mux *http.ServeMux) {
	mux.HandleFunc("/v1/pong", pong.Pong)
}
