package forgery

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

    // This property is a slice containing properties mapped to the named route "parameters". 
    // For example if you have the route "/user/:name", then the "name" property is available 
    // to you as req.params["name"]. This object defaults to {}.
    Params map[string]string

    // This property is a map containing the parsed request body. 
    // This feature is provided by the bodyParser() middleware, though other body 
    // parsing middleware may follow this convention as well.
    // This property defaults to {} when bodyParser() is used.
    Body map[string]string

    // This property is a map containing the parsed query-string, defaulting to {}.
    Query map[string]string

    // This property is a map of the files uploaded. This feature is provided 
    // by the bodyParser() middleware, though other body parsing middleware may 
    // follow this convention as well. This property defaults to {} when bodyParser() is used.
    Files map[string]interface{}

    // The currently matched Route containing several properties such as the 
    // route's original path string, the regexp generated, and so on.
    Route interface{}

    // When the cookieParser() middleware is used this object defaults to {}, 
    // otherwise contains the cookies sent by the user-agent.
    Cookies map[string]string

    // When the cookieParser(secret) middleware is used this object defaults to {}, 
    // otherwise contains the signed cookies sent by the user-agent, unsigned and ready for use. 
    // Signed cookies reside in a different object to show developer intent, otherwise a 
    // malicious attack could be placed on `req.cookie` values which are easy to spoof. 
    // Note that signing a cookie does not mean it is "hidden" nor encrypted, this simply 
    // prevents tampering as the secret used to sign is private.
    SignedCookies map[string]string

    // Return an array of Accepted media types ordered from highest quality to lowest.
    Accepted map[string]string

    // Return the remote address, or when "trust proxy" is enabled - the upstream address.
    Ip string

    // When "trust proxy" is `true`, parse the "X-Forwarded-For" ip address list and return a slice, 
    // otherwise an empty array is returned. For example if the value were "client, proxy1, proxy2" 
    // you would receive the slice {"client", "proxy1", "proxy2"} where "proxy2" is the furthest down-stream.
    Ips []string

    // Returns the request URL pathname.
    Path string

    // Check if the request is fresh - aka Last-Modified and/or the ETag still match, 
    // indicating that the resource is "fresh".
    Fresh bool

    // Check if the request is stale - aka Last-Modified and/or the ETag do not match, 
    // indicating that the resource is "stale".
    Stale bool

    // Check if the request was issued with the "X-Requested-With" header field set to "XMLHttpRequest" (jQuery etc).
    Xhr bool

    // Return the protocol string "http" or "https" when requested with TLS. 
    // When the "trust proxy" setting is enabled the "X-Forwarded-Proto" header field will be trusted. 
    // If you're running behind a reverse proxy that supplies https for you this may be enabled.
    Protocol string

    // Check if a TLS connection is established. This is a short-hand for: "https" == req.Protocol
    Secure bool

    // Return an slice of Accepted languages ordered from highest quality to lowest.
    AcceptedLanguages []string

    // Return an slice of Accepted charsets ordered from highest quality to lowest.
    AcceptedCharsets []string

    // Holds custom values set by functions in the request flow.
    Map map[string]interface{}
}

/*
    Returns a new Request.
*/

func createRequest(req *stackr.Request) (*Request) {

    this := &Request{}

    this.Request = req

    if this.Params == nil {
        this.Params = map[string]string{}
    }

    if this.Body == nil {
        this.Body = map[string]string{}
    }

    if this.Query == nil {
        this.Query = map[string]string{}
    }

    if this.Files == nil {
        this.Files = map[string]interface{}{}
    }

    if this.Cookies == nil {
        this.Cookies = map[string]string{}
    }

    if this.SignedCookies == nil {
        this.SignedCookies = map[string]string{}
    }

    this.Ip = this.RemoteAddr
    // this.Ips
    this.Path = this.URL.Path
    // this.Fresh
    // this.Stale
    // this.Xhr
    this.Protocol = this.URL.Scheme
    this.Secure = this.Protocol == "https"
    // this.AcceptedLanguages
    // this.AcceptedCharsets

    if this.Map == nil {
        this.Map = map[string]interface{}{}
    }

    /*
        Return the finished forgery.Request.
    */

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
func (this *Request) Accepts(t string) {
    panic(halt)
}

/*
    Check if the incoming request contains the "Content-Type" header field, and it matches the give mime "type".
*/
func (this *Request) Is(t string) (bool) {
    h := this.Get("content-type")
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