package forgery

import(
    "github.com/ricallinson/stackr"
)

type Server struct {

    // A stackr.Server type
    *stackr.Server

    // Application local variables are provided to all templates rendered within the application. 
    // This is useful for providing helper functions to templates, as well as app-level data.
    Locals map[string]string

    // 
    Routes map[string]string
}

/*
    Create a new stackr server.
*/
func CreateServer() (*Server) {
    return &Server{
        Server: &stackr.Server{},
        Locals: map[string]string{},
        Routes: map[string]string{},
    }
}

/*
    Create forgery Request, Response types from stackr Request, Response types.
*/
func make(req *stackr.Request, res *stackr.Response) (*Request, *Response) {
    return createRequest(req), createResponse(res)
}

/*
    Assigns setting "name" to "value".
*/
func (this *Server) Set(n string, v string) {

}

/*
    Get setting "name" value.
*/
func (this *Server) Get(n string) {

}

/*
    Set setting "name" to "true".
*/
func (this *Server) Enable(n string) {

}

/*
    Set setting "name" to "false".
*/
func (this *Server) Disable(n string) {

}

/*
    Check if setting "name" is enabled.
*/
func (this *Server) Enabled(n string) {

}

/*
    Check if setting "name" is disabled.
*/
func (this *Server) Disabled(n string) {

}

/*
    Not implemented.
*/
func (this *Server) Configure(n string, fn func()) {

}

/*
    Register the given template engine callback as ext.
*/
func (this *Server) Engine(ext string, fn func()) {

}

/*
    Map logic to route parameters. For example when ":user" is 
    present in a route path you may map user loading logic to 
    automatically provide req.Map["user"] to the route, or perform 
    validations on the parameter input.
*/
func (this *Server) Param(p string, fn func()) {

}

/*
    Render a "view" with a callback responding with the rendered string. 
    This is the app-level variant of "res.render()", and otherwise behaves the same way.
*/
func (this *Server) Render(v string, opt interface{}, fn func()) {

}

/*
    This method functions just like the app.Verb(verb, ...) method, however it matches all HTTP verbs.
*/
func (this *Server) All(path string, fn ...func(*Request, *Response, func())) {

}

/*
    The method provides the routing functionality in Forgery, where "verb" is one of the HTTP verbs, 
    such as app.Verb("post", ...). Multiple callbacks may be given, all are treated equally, 
    and behave just like middleware, with the one exception that these callbacks may invoke 
    next('route') to bypass the remaining route callback(s). This mechanism can be used to perform 
    pre-conditions on a route then pass control to subsequent routes when there is no reason to 
    proceed with the route matched.
*/
func (this *Server) Verb(verb string, path string, fn ...func(*Request, *Response, func())) {

}