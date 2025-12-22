package apis

import (
	"fmt"
	"net/http"

	"github.com/dlbarduzzi/sentinel/core"
	"github.com/dlbarduzzi/sentinel/tools/event"
)

type route struct {
	pattern string
	handler func(*core.EventRequest)
}

type router struct {
	app    core.App
	routes []route
}

func newRouter(app core.App) *router {
	r := &router{app: app}
	bindHealthApi(r)
	return r
}

func (r *router) add(pattern string, handler func(*core.EventRequest)) {
	r.routes = append(r.routes, route{
		pattern: pattern,
		handler: handler,
	})
}

func (r *router) get(pattern string, handler func(*core.EventRequest)) {
	r.add(fmt.Sprintf("GET %s", pattern), handler)
}

func (r *router) buildMux() http.Handler {
	mux := http.NewServeMux()

	for _, route := range r.routes {
		mux.HandleFunc(route.pattern, func(res http.ResponseWriter, req *http.Request) {
			route.handler(&core.EventRequest{
				App: r.app,
				Event: event.Event{
					Request:  req,
					Response: res,
				},
			})
		})
	}

	return mux
}
