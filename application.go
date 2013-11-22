/*
   __WARNING: WORK IN PROGRESS__

   Forgery is a minimal and flexible golang web application framework, providing a robust set of
   features for building single and multi-page, web applications.

   Note: This project started out as a clone of the superb Node.js library [Express](http://expressjs.com/).
*/
package f

import (
	"errors"
	"github.com/ricallinson/httputils"
	"github.com/ricallinson/stackr"
	"os"
	"path/filepath"
)

const (
	TRUE  = "true"
	FALSE = "false"
	halt  = "HALT: Not implemented yet!"
)

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
	engines map[string]Renderer
}

/*
   Create a new stackr server.

   * "env" Environment mode, defaults to $GO_ENV or "development"
   * "trust proxy" Enables reverse proxy support, disabled by default
   * "jsonp callback name" Changes the default callback name of "?callback="
   * "json spaces" JSON response spaces for formatting, defaults to 2 in development, 0 in production
   * "case sensitive routing" Enable case sensitivity, disabled by default, treating "/Foo" and "/foo" as the same
   * "strict routing" Enable strict routing, by default "/foo" and "/foo/" are treated the same by the router
   * X "view cache" Enables view template compilation caching, enabled in production by default
   * "view engine" The default engine extension to use when omitted
   * "views" The view directory path, defaulting to "./views"
*/
func CreateServer() *Server {
	this := &Server{
		Server:   &stackr.Server{},
		Locals:   map[string]string{},
		Router:   &Router{},
		settings: map[string]string{},
		engines:  map[string]Renderer{},
	}
	this.defaultConfiguration()
	return this
}

/*
   Initialize application configuration.
*/
func (this *Server) defaultConfiguration() {

	cwd, err := os.Getwd()

	if err != nil {
		panic("Cannot get current working directory!")
	}

	// default settings
	this.Enable("x-powered-by")
	this.Enable("etag")
	this.Set("env", os.Getenv("GO_ENV"))
	if this.Get("env") == "" {
		this.Set("env", "development")
	}

	// default configuration
	this.Configure(func() {
		this.Set("subdomain offset", "2")
		this.Set("views", filepath.Join(cwd, "views"))
		this.Set("jsonp callback name", "callback")
		this.Set("app path", "/")
	})

	this.Configure("development", func() {
		this.Set("json spaces", "  ")
	})

	this.Configure("production", func() {
		this.Enable("view cache")
	})
}

/*
   Configure callback for zero or more envs,
   when no `env` is specified that callback will
   be invoked for all environments. Any combination
   can be used multiple times, in any order desired.

   Examples:

       app.Configure(func() {
         // executed for all envs
       })

       app.Configure("stage", func() {
         // executed staging env
       })

       app.Configure("stage", "production", func() {
         // executed for stage and production
       })

   Note:

   These callbacks are invoked immediately, and
   are effectively sugar for the following:

   var env = os.Getenv("GO_ENV")

   switch (env) {
   case 'development':
   ...
   case 'stage':
   ...
   case 'production':
   ...
   }
*/
func (this *Server) Configure(i ...interface{}) {

	var envs []string
	var fn func()

	// Look at the given inputs.
	for _, t := range i {
		switch t.(type) {
		case string:
			envs = append(envs, t.(string))
		case func():
			fn = t.(func())
		}
	}

	// If there are no envs call the func and return.
	if len(envs) == 0 {
		fn()
		return
	}

	// Loop over the envs until a match is found.
	// Then call the function.
	for _, e := range envs {
		if e == this.Get("env") {
			fn()
			return
		}
	}
}

/*
   Returns the root of this app.
*/
func (this *Server) Path() string {
	return this.Get("app path")
}

/*
   Assigns setting "name" to "value".
*/
func (this *Server) Set(n string, v ...string) string {
	if len(v) == 0 {
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
func (this *Server) Get(path string, fn ...func(*Request, *Response, func())) string {

	/*
	   If there is no function then this is really a call to .Set()
	*/

	if len(fn) == 0 {
		return this.Set(path)
	}

	/*
	   Otherwise it's a call to .Verb()
	*/

	this.Verb("GET", path, fn...)

	return ""
}

/*
   Set setting "name" to "true".
*/
func (this *Server) Enable(n string) {
	this.Set(n, TRUE)
}

/*
   Set setting "name" to "false".
*/
func (this *Server) Disable(n string) {
	this.Set(n, FALSE)
}

/*
   Check if setting "name" is enabled.
*/
func (this *Server) Enabled(n string) bool {
	return this.Get(n) == TRUE
}

/*
   Check if setting "name" is disabled.
*/
func (this *Server) Disabled(n string) bool {
	return this.Get(n) == FALSE
}

/*
   Register the given template engine callback as ext.
*/
func (this *Server) Engine(ext string, fn Renderer) {
	this.engines[ext] = fn
}

/*
   Render a "view" responding with the rendered string.
   This is the app-level variant of "res.render()", and otherwise behaves the same way.
*/
func (this *Server) Render(view string, i ...interface{}) (string, error) {

	ext := filepath.Ext(view)

	if _, ok := this.engines[ext]; ok == false {
		return "", errors.New("Engine not found.")
	}

	file := filepath.Join(this.Get("views"), view)

	if _, err := os.Stat(file); err != nil || os.IsNotExist(err) {
		return "", errors.New("Failed to lookup view '" + file + "'")
	}

	i = append(i, this.Locals)

	t, err := this.engines[ext].Render(file, i...)

	if err != nil {
		return "", errors.New("Problem rendering view.")
	}

	return t, nil
}

/*
   This method functions just like the app.Verb(verb, ...) method, however it matches all HTTP verbs.
*/
func (this *Server) All(path string, fn ...func(*Request, *Response, func())) {
	for _, verb := range httphelp.Methods {
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
		this.Router.CaseSensitive = this.Enabled("case sensitive routing")
		this.Router.Strict = this.Enabled("strict routing")
		this.Use(this.Router.Middleware(this))
		this.usedRouter = true
	}

	this.Router.AddRoute(verb, path, funcs...)
}

/*
   Map logic to route parameters. For example when ":user" is
   present in a route path you may map user loading logic to
   automatically provide req.Map["user"] to the route, or perform
   validations on the parameter input.
*/
func (this *Server) Param(p string, fn func(*Request, *Response, func())) {
	this.Router.AddParamFunc(p, fn)
}
