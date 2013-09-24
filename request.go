package f

import(
    "strings"
    "github.com/ricallinson/stackr"
)

/*
    A Request represents an HTTP request received by the server.
*/
type Request struct {

    // The stackr.Request type.
    *stackr.Request

    // The application server.
    app *Server

    // This property is a slice containing properties mapped to the named route "parameters". 
    // For example if you have the route "/user/:name", then the "name" property is available 
    // to you as req.params["name"]. This object defaults to {}.
    Params map[string]string

    // The currently matched Route containing several properties such as the 
    // route's original path string, the regexp generated, and so on.
    Route interface{}
}

/*
    Returns a new Request.
*/

func createRequest(req *stackr.Request, app *Server) (*Request) {

    this := &Request{}

    this.Request = req

    this.app = app

    if this.Params == nil {
        this.Params = map[string]string{}
    }

    return this
}

/*
    Return the value of param "name" when present. Lookup is performed in the following order:

    * Params
    * Body
    * Query

    Direct access to req.body, req.params, and req.query should be favoured for clarity - 
    unless you truly accept input from each object.
*/
func (this *Request) Param(n string) (string) {
    var v string
    var ok bool
    v, ok = this.Params[n]
    if ok {
        return v
    }
    v, ok = this.Body[n]
    if ok {
        return v
    }
    v, ok = this.Query[n]
    if ok {
        return v
    }
    return ""
}

/*
    Get the case-insensitive request header field. The Referrer and Referer fields are interchangeable.
*/
func (this *Request) Get(f string) (string) {

    /*
        Possible future bug.
        http://www.w3.org/Protocols/rfc2616/rfc2616-sec4.html#sec4.2
        Message headers are case-insensitive.
    */

    return this.Header.Get(f)
}

/*
    Check if the given types are acceptable, returning the best match when true, 
    otherwise undefined - in which case you should respond with 406 "Not Acceptable".
*/
func (this *Request) Accepts(t string) (bool) {
    for _, v := range this.Accepted {
        if strings.ToLower(t) == v {
            return true
        }
    }
    return false
}

/*
    Check if the incoming request contains the "Content-Type" header field, and it matches the give mime "type".
*/
func (this *Request) Is(t string) (bool) {
    h := this.Get("Content-Type")
    return strings.ToLower(h) == strings.ToLower(t)
}

/*
    Check if the given "charset" is acceptable.
*/
func (this *Request) AcceptsCharset(c string) {
    panic(halt)
}

/*
    Check if the given "lang" is acceptable.
*/
func (this *Request) AcceptsLanguage(l string) {
    panic(halt)
}