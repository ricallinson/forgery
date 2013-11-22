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

func (this *Router) AddParamFunc(name string, fn func(*Request, *Response, func())) {
	this.ParamFuncs[name] = fn
}

func (this *Router) AddRoute(verb string, path string, funcs ...func(*Request, *Response, func())) {

	route := &Route{
		Method:        strings.ToUpper(verb),
		Path:          path,
		Funcs:         funcs,
		CaseSensitive: this.CaseSensitive,
		Strict:        this.Strict,
	}

	this.Routes = append(this.Routes, route)
}

/*
	Calls the function defined for each key in "req.Params" if available.
*/
func (this *Router) executeParamFuncs(req *Request, res *Response, next func()) {
	// Call param functions if they are set.
	for name := range req.Params {
		if fn, ok := this.ParamFuncs[name]; ok {
			fn(req, res, next)
			// If the response has been "closed" by a "param" function then return.
			if res.Closed == true {
				return
			}
		}
	}
}

/*
	Calls each function defined in the given route.
*/
func (this *Router) executeRouteFuncs(req *Request, res *Response, next func(), route *Route) {
	for _, fn := range route.Funcs {
		fn(req, res, next)
		// If the response has been "closed" by a "route" function then return.
		if res.Closed == true {
			return
		}
	}
}

func (this *Router) Middleware(app *Server) func(req *stackr.Request, res *stackr.Response, next func()) {

	this.ParamFuncs = map[string]func(*Request, *Response, func()){}

	/*
		This func is in the serving path so has to be super fast!
	*/
	return func(req *stackr.Request, res *stackr.Response, next func()) {

		// Relative to path only so remove the query string.
		path := req.OriginalUrl
		if i := strings.Index(path, "?"); i > 0 {
			path = path[:i]
		}

		// Create the f.Request and f.Response from stackr.Request and stackr.Response
		freq := createRequest(req, app)
		fres := createResponse(res, next, app)
		freq.SetResponse(fres) // Add the Response to the Request
		fres.SetRequest(freq)  // Add the Request to the Response

		for _, route := range this.Routes {
			// If the route matches use it.
			if params, ok := route.Match(req.Method, path); ok {
				// Set the route "params" found in the path.
				freq.Params = params
				// Call "param" functions if they are defined.
				this.executeParamFuncs(freq, fres, next)
				// Call "route" functions if they are defined.
				this.executeRouteFuncs(freq, fres, next, route)
			}
		}
	}
}
