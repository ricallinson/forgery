package f

import (
	"github.com/ricallinson/stackr"
	"strings"
)

type Router struct {

	// Holds the route mappings this Router will use.
	Routes []*Route

	// Params that trigger functions
	ParamFuncs map[string]func(*Request, *Response, func())

	// Enable case-sensitive routes
	CaseSensitive bool

	// Enable strict matching for trailing slashes
	Strict bool
}

/*
	Add a "param" function to the Router.
*/
func (this *Router) AddParamFunc(name string, fn func(*Request, *Response, func())) {
	this.ParamFuncs[name] = fn
}

/*
	Add a new route to the Router.
*/
func (this *Router) AddRoute(verb string, path string, funcs ...func(*Request, *Response, func())) {

	route := &Route{
		Method:        strings.ToUpper(verb),
		Path:          path,
		Callbacks:     funcs,
		CaseSensitive: this.CaseSensitive,
		Strict:        this.Strict,
	}

	this.Routes = append(this.Routes, route)
}

/*
	Calls the function defined for each key in "req.Params" if available.
*/
func (this *Router) executeParamFuncs(req *Request, res *Response, next func()) bool {
	// Call each "param" function if one is set.
	for name := range req.Params {
		if pfn, ok := this.ParamFuncs[name]; ok {
			pfn(req, res, next)
			// If the response has been "closed" by a "param" function then return.
			if res.Closed == true {
				return true
			}
		}
	}
	return false
}

/*
	Calls each function defined in the given route.
*/
func (this *Router) executeRouteFuncs(req *Request, res *Response, next func(), route *Route) bool {
	for _, rfn := range route.Callbacks {
		rfn(req, res, next)
		// If the response has been "closed" by a "route" function then return.
		if res.Closed == true {
			return true
		}
	}
	return false
}

func (this *Router) handle(req *Request, res *Response, next func()) {

	// Relative to path only so remove the query string.
	path := req.OriginalUrl
	if i := strings.Index(path, "?"); i > 0 {
		path = path[:i]
	}

	for _, route := range this.Routes {
		// If the route matches use it.
		if params, ok := route.Match(req.Method, path); ok {
			// Set the route "params" found in the path.
			req.Params = params
			// Set the matched route.
			req.Route = route
			// Call "param" functions if they are defined.
			// If it returns true then the connection has closed so stop processing.
			if this.executeParamFuncs(req, res, next) {
				return
			}
			// Call "route" functions if they are defined.
			// If it returns true then the connection has closed so stop processing.
			if this.executeRouteFuncs(req, res, next, route) {
				return
			}
		}
	}

	// If nothing closed the conection then call next().
	next()
}

func (this *Router) Middleware(app *Server) func(req *stackr.Request, res *stackr.Response, next func()) {

	this.ParamFuncs = map[string]func(*Request, *Response, func()){}

	return func(req *stackr.Request, res *stackr.Response, next func()) {

		// Create the f.Request and f.Response from stackr.Request and stackr.Response
		freq := createRequest(req, app)
		fres := createResponse(res, next, app)
		freq.SetResponse(fres) // Add the Response to the Request
		fres.SetRequest(freq)  // Add the Request to the Response

		// Handle the request.
		this.handle(freq, fres, next)
	}
}
