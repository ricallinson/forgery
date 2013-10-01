package f

import(
    "github.com/ricallinson/stackr"
)

type Route struct {

    // HTTP Method
    Method string

    // URL Path
    Url string

    // Slice of functions
    Funcs []func(*Request, *Response, func())
}

type Router struct {

    // Holds the route mappings this Router will use.
    Routes []Route

    // Params that trigger functions
    Params map[string]func(*Request, *Response, func())
}

func (this *Router) Middleware(app *Server) (func(req *stackr.Request, res *stackr.Response, next func())) {
    
    this.Routes = []Route{}

    this.Params = map[string]func(*Request, *Response, func()){}

    return func(req *stackr.Request, res *stackr.Response, next func()) {

        for _, route := range this.Routes {

            if req.Method == route.Method && req.OriginalUrl == route.Url {

                for _, fn := range route.Funcs {
                    freq := createRequest(req, app)
                    fres := createResponse(res, next, app)
                    freq.res = fres // Add the Response to the Request
                    fres.req = freq // Add the Request to the Response
                    fn(freq, fres, next)
                }

                return
            }
        }
    }
}

func (this *Router) Param(p string, fn func(*Request, *Response, func())) {
    this.Params[p] = fn
}