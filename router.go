package f

import(
    "github.com/ricallinson/stackr"
)

type Router struct {

    // Holds the route mappings this Router will use.
    Routes map[string]string

    // Params that trigger functions
    params map[string]func(*Request, *Response, func())
}

func (this *Router) Middleware() (func(req *stackr.Request, res *stackr.Response, next func())) {
    
    this.Routes = map[string]string{}

    this.params = map[string]func(*Request, *Response, func()){}

    return func(req *stackr.Request, res *stackr.Response, next func()) {
        // request
    }
}

func (this *Router) Param(p string, fn func(*Request, *Response, func())) {
    this.params[p] = fn
}