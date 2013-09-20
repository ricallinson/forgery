/*
    __WARNING: WORK IN PROGRESS__

    Forgery is a minimal and flexible golang web application framework, providing a robust set of 
    features for building single and multi-page, web applications.

    Note: This project started out as a clone of the superb Node.js library [Express](http://expressjs.com/).
*/
package forgery

import(
    "strings"
    "github.com/ricallinson/stackr"
)

var halt = "HALT: Not implemented yet!"

type Server struct {

    // A stackr.Server type
    *stackr.Server

    // Application local variables are provided to all templates rendered within the application. 
    // This is useful for providing helper functions to templates, as well as app-level data.
    Locals map[string]string

    // The Router middleware function.
    Router *Router

    // Has the Router been added to stackr.
    usedRouter bool

    // Stores the applications settings.
    settings map[string]string

    // The rendering engines assigned.
    engines map[string]func()
}

/*
    Create a new stackr server.
*/
func CreateServer() (*Server) {
    return &Server{
        Server: &stackr.Server{},
        Locals: map[string]string{},
        Router: &Router{},
        settings: map[string]string{},
        engines: map[string]func(){},
    }
}

/*
    Assigns setting "name" to "value".
*/
func (this *Server) Set(n string, v ...string) (string) {
    if len(v) == 1 {
        return this.settings[n]
    }
    this.settings[n] = v[0]
    return v[0]
}

/*
    Get setting "name" value.
    or;
    Provides the routing functionality for GET requests to the given "path".
*/
func (this *Server) Get(path string, fn ...func(*Request, *Response, func())) (string) {

    /*
        If there is no function then this is really a call to .Set()
    */

    if len(fn) == 0 {
        return this.Set(path)
    }

    /*
        Otherwise it's a call to .Verb()
    */

    this.Verb("GET", path, fn...);

    return ""
}

/*
    Set setting "name" to "true".
*/
func (this *Server) Enable(n string) {
    this.Set(n, "TRUE")
}

/*
    Set setting "name" to "false".
*/
func (this *Server) Disable(n string) {
    this.Set(n, "FALSE")
}

/*
    Check if setting "name" is enabled.
*/
func (this *Server) Enabled(n string) (bool) {
    return this.Get(n) == "TRUE"
}

/*
    Check if setting "name" is disabled.
*/
func (this *Server) Disabled(n string) (bool) {
    return this.Get(n) == "FALSE"
}

/*
    Not supported.
*/
func (this *Server) Configure(n string, fn func()) {
    panic("ERROR: app.Configure() is not supported")
}

/*
    Register the given template engine callback as ext.
*/
func (this *Server) Engine(ext string, fn func()) {
    this.engines[ext] = fn
}

/*
    Map logic to route parameters. For example when ":user" is 
    present in a route path you may map user loading logic to 
    automatically provide req.Map["user"] to the route, or perform 
    validations on the parameter input.
*/
func (this *Server) Param(p string, fn func(*Request, *Response, func())) {
    this.Router.Param(p, fn)
}

/*
    Render a "view" with a callback responding with the rendered string. 
    This is the app-level variant of "res.render()", and otherwise behaves the same way.
*/
func (this *Server) Render(v string, opt interface{}, fn func()) {
    panic(halt)
}

/*
    This method functions just like the app.Verb(verb, ...) method, however it matches all HTTP verbs.
*/
func (this *Server) All(path string, fn ...func(*Request, *Response, func())) {
    for _, verb := range methods {
        this.Verb(verb, path, fn...)
    }
}

/*
    The method provides the routing functionality in Forgery, where "verb" is one of the HTTP verbs, 
    such as app.Verb("post", ...). Multiple callbacks may be given, all are treated equally, 
    and behave just like middleware, with the one exception that these callbacks may invoke 
    next('route') to bypass the remaining route callback(s). This mechanism can be used to perform 
    pre-conditions on a route then pass control to subsequent routes when there is no reason to 
    proceed with the route matched.
*/
func (this *Server) Verb(verb string, path string, funcs ...func(*Request, *Response, func())) {

    if this.usedRouter == false {
        this.Use("/", this.Router.Middleware())
        this.usedRouter = true
    }

    verb = strings.ToUpper(verb)

    /*
        This is temporary code in place of a real URL router.
    */

    this.Use(path, func(req *stackr.Request, res *stackr.Response, next func()) {

        if strings.ToUpper(req.Method) != verb {
            return
        }

        for _, fn := range funcs {
            fn(createRequest(req), createResponse(res), next)
        }
    })
}