package apis

import (
	"net/http"

	"github.com/dlbarduzzi/sentinel/core"
)

type router struct {
	app core.App
}

func newRouter(app core.App) *router {
	r := &router{app: app}
	return r
}

func (r *router) buildMux() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/health", healthCheck)
	return mux
}
