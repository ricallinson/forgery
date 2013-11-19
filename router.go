package f

import(
    "strings"
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

        // Relative to path only so remove the query string.
        path := req.OriginalUrl
        if i := strings.Index(path, "?"); i > 0 {
            path = path[:i]
        }

        // Create the f.Request and f.Response
        freq := createRequest(req, app)
        fres := createResponse(res, next, app)
        freq.SetResponse(fres) // Add the Response to the Request
        fres.SetRequest(freq) // Add the Request to the Response

        for _, route := range this.Routes {
            // If the route matches use it.
            if req.Method == route.Method && (path == route.Url || "*" == route.Url) {
                // Call each function the route has.
                for _, fn := range route.Funcs {
                    fn(freq, fres, next)
                    // If the response has been "closed" by the route function then return.
                    if res.Closed == true {
                        return
                    }
                }
            }
        }
    }
}