package f

import (
    "github.com/ricallinson/stackr"
    "strings"
)

type Router struct {

	// Holds the route mappings this Router will use.
	Routes []*Route

	// Params that trigger functions
	Params map[string]func(*Request, *Response, func())
}

func (this *Router) Middleware(app *Server) func(req *stackr.Request, res *stackr.Response, next func()) {

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
		fres.SetRequest(freq)  // Add the Request to the Response

		for _, route := range this.Routes {
			// If the route matches use it.
			if params, ok := route.Match(req.Method, path); ok {
				// Set the route params found in the path.
				freq.Params = params
				// Call each function the route has.
				for _, fn := range route.Funcs {
					fn(freq, fres, next)
					// If the response has been "closed" by a route function then return.
					// This happens on res.End() res.Send(), res.Json(), res.Jsonp() or res.Render()
					if res.Closed == true {
						return
					}
				}
			}
		}
	}
}
