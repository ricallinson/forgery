package forgery

import(
    "github.com/ricallinson/stackr"
)

type Router struct {

    // Holds the route mappings this Router will use.
    Routes map[string]string
}

func (this *Router) Middleware() (func(req *stackr.Request, res *stackr.Response, next func())) {
    
    this.Routes = map[string]string{}

    return func(req *stackr.Request, res *stackr.Response, next func()) {
        // request
    }
}