package f

import(
    "fmt"
    "mime"
    "html"
    "time"
    "strings"
    "net/url"
    "net/http"
    "path/filepath"
    "encoding/json"
    "github.com/ricallinson/stackr"
    "github.com/ricallinson/httphelp"
)

/*
    Response represents the response from an HTTP request.
*/
type Response struct {

    // The stackr.Response.
    *stackr.Response

    // The forgery.Request matched to this Response.
    req *Request

    // The next function.
    next func()

    // The application server.
    app *Server

    // Assign the charset. Defaults to "utf-8".
    Charset string

    // Response local variables are scoped to the request, thus only available 
    // to the view(s) rendered during that request / response cycle, if any. 
    // Otherwise this API is identical to app.Locals.
    Locals map[string]string
}

/*
    Returns a new HTTP Response.
*/

func createResponse(res *stackr.Response, next func(), app *Server) (*Response) {

    this := &Response{}

    this.Response = res

    this.app = app

    this.next = next

    this.Charset = "utf-8"

    this.Locals = map[string]string{}

    return this
}

/*
    Chainable alias of stackr's "res.StatusCode=".
*/
func (this *Response) Status(c int) {
    this.StatusCode = c
}

/*
    Set header "field" to "value".
*/
func (this *Response) Set(f string, v string) {
    this.SetHeader(f, v)
}

/*
    Get the case-insensitive response header "field".
*/
func (this *Response) Get(f string) (string) {
    return this.Writer.Header().Get(f)
}

/*
    Set cookie "name" to "value", where "value" may be a string or interface 
    converted to JSON. The "path" option defaults to "/".

    Set cookie `name` to `val`, with the given `options`.

    Options:

        - `maxAge`   max-age in milliseconds, converted to `expires`
        - `signed`   sign the cookie
        - `path`     defaults to "/"

    Examples:

        // "Remember Me" for 15 minutes
        res.Cookie("rememberme", "1", &httpCookie{ expires: new Date(Date.now() + 900000), httpOnly: true });

        // save as above
        res.Cookie("rememberme", "1", &httpCookie{ maxAge: 900000, httpOnly: true })
*/
func (this *Response) Cookie(n string, i interface{}, o ...*http.Cookie) {

    var cookie *http.Cookie

    /*
        If we were given cookie options use them.
    */

    if len(o) == 1 {
        cookie = o[0]
    } else {
        cookie = &http.Cookie{}
    }

    /*
        Convert the given interface to a string so it can be used as a cookie value.
    */

    var v string
    switch i.(type) {
    default:
        b, e := json.Marshal(i)
        v = string(b)
        if e != nil {
            v = e.Error()
        }
    case string:
        v = i.(string)
    }

    cookie.Name = n
    cookie.Value = url.QueryEscape(Encode(v))

    if cookie.Path == "" {
        cookie.Path = "/"
    }

    if cookie.MaxAge == 0 {
        // max-age in milliseconds, converted to `expires`
        // TODO: Check the timing here.
        cookie.Expires = time.Now().UTC().Add(time.Duration(cookie.MaxAge) * (time.Millisecond * time.Microsecond))
        cookie.MaxAge = cookie.MaxAge / 1000
    }

    // cookie.Domain = ""
    // cookie.Secure = false
    // cookie.HttpOnly = false

    /*
        Possible bug if headers are already sent.
    */

    http.SetCookie(this.Writer, cookie)
}

func (this *Response) SignedCookie(n string, i interface{}, o ...*http.Cookie) {

    secret := this.app.Get("secret")

    if secret == "" {
        panic("f.Get(\"secret\") required for signed cookies")
    }

    /*
        Convert the given interface to a string so it can be signed.
    */

    var v string
    switch i.(type) {
    default:
        b, _ := json.Marshal(i)
        v = string(b)
    case string:
        v = i.(string)
    }

    /*
        Encrypt the cookie string value.
    */

    v = Sign(v, secret)

    this.Cookie(n, v, o...)
}

/*
    Clear cookie "name". The "path" option defaults to "/".
*/
func (this *Response) ClearCookie(n string, o ...*http.Cookie) {

    var opt *http.Cookie

    if len(o) == 1 {
        opt = o[0]
    } else {
        opt = &http.Cookie{}
    }

    opt.MaxAge = -1

    if opt.Path == "" {
        opt.Path = "/"
    }

    this.Cookie(n, "", opt)
}

/*
    Redirect to the given "url" with optional "status" code defaulting to 302 "Found".
*/
func (this *Response) Redirect(uri string, s ...int) {

    req := this.req
    isHead := req.Method == "HEAD"

    this.StatusCode = 302

    // If we were given a status, us it.
    if len(s) == 1 {
        this.StatusCode = s[0]
    }

    this.Location(uri)

    uri = this.Get("Location")

    body := ""

    if this.req.Accepts("text/plain") {
        body = httphelp.StatusCodes[this.StatusCode] + ". Redirecting to " + uri;
    } else if this.req.Accepts("text/html") {
        u := html.EscapeString(uri)
        body = "<p>" + httphelp.StatusCodes[this.StatusCode] + ". Redirecting to <a href=\"" + u + "\">" + u + "</a></p>";
    }

    this.Set("Content-Length", fmt.Sprint(len(body)))

    if isHead {
        body = ""
    }

    this.End(body)
}

/*
    Set the location header.
*/
func (this *Response) Location(uri string) {

    app := this.app
    req := this.req

    /*
        If the given uri == "back" then try and get the "Referrer"; or use "/"
    */

    if uri == "back" {
        uri = req.Get("Referrer")
        if uri == "" {
            uri = "/"
        }
    }

    /*
        If the uri is relative then build it.
    */

    if strings.Index(uri, "://") == -1 {
        if uri[:1] == "." {
            // relative to path
            path := req.OriginalUrl
            if i := strings.Index(path, "?"); i > 0 {
                path = path[:i]
            }
            i := strings.Index(path, "://") + 2
            uri = path[:i] + filepath.Join(path[i:], uri)
        } else if uri[:1] != "/" {
            // relative to mount-point
            uri = filepath.Join(app.Path(), uri)
        }
    }

    this.Set("Location", uri)
}

/*
    Send a response. This method performs a myriad of useful tasks for simple non-streaming 
    responses such as automatically assigning the Content-Length unless previously defined 
    and providing automatic HEAD and HTTP cache freshness support.

    res.Send([]byte{114, 105, 99}]);
    res.Send(map[string]string{"some": "json"});
    res.Send("some html");
    res.Send("Sorry, we cannot find that!", 404);
    res.Send(map[string]string{"error": "msg"}, 500);
    res.Send(200);
*/
func (this *Response) Send(b interface{}, s ...int) {

    var body string
    var bytes []byte

    req := this.req
    isHead := req.Method == "HEAD"

    // If we were given a status, us it.
    if len(s) == 1 {
        this.StatusCode = s[0]
    }

    switch b.(type) {
    default: // JSON
        this.Json(b, s...)
        return
    case int: // Status Code
        if len(this.Get("Content-Type")) == 0 {
            this.ContentType("text/plain")
        }
        this.StatusCode = b.(int)
        body = httphelp.StatusCodes[b.(int)]
    case string:
        if len(this.Get("Content-Type")) == 0 {
            this.ContentType("text/html")
        }
        if len(this.Charset) == 0 {
            this.Charset = "utf-8"
        }
        body = b.(string)
    case []byte:
        if len(this.Get("Content-Type")) == 0 {
            this.ContentType("text/html")
        }
        if len(this.Charset) == 0 {
            this.Charset = "utf-8"
        }
        bytes = b.([]byte)
    }

    // Populate Content-Length
    if len(this.Get("Content-Length")) == 0 {
        if len(bytes) > 0 {
            this.Set("Content-Length", fmt.Sprint(len(bytes)))
        } else {
            this.Set("Content-Length", fmt.Sprint(len(body)))
        }
    }

    // ETag support

    // Freshness
    if req.Fresh() {
        this.StatusCode = 304;
    }

    // Strip irrelevant headers
    if this.StatusCode == 204 || this.StatusCode == 304 {
        this.RemoveHeader("Content-Type");
        this.RemoveHeader("Content-Length");
        this.RemoveHeader("Transfer-Encoding");
        body = "";
    }

    if isHead {
        body = ""
    } else if len(bytes) > 0 {
        this.WriteBytes(bytes)
    }

    this.End(body)
}

/*
    Given an interface returns a JSON string.
*/
func (this *Response) json(i interface{}) (string) {
    if len(this.Get("Content-Type")) == 0 {
        this.ContentType("application/json")
    }
    if len(this.Charset) == 0 {
        this.Charset = "utf-8"
    }
    b, err := json.Marshal(i)
    if err != nil {
        return ""
    }
    return string(b)
}

/*
    Send a JSON response. This method is identical to res.Send() when an object or slice is passed, 
    however it may be used for explicit JSON conversion of non-objects 
    (null, undefined, etc), though these are technically not valid JSON.

    res.Json(null)
    res.Json(map[string]string{"user": "ric"})
    res.Json(map[string]string{"error": "msg"}, 500)
*/
func (this *Response) Json(i interface{}, s ...int) {

    // If we were given a status, us it.
    if len(s) == 1 {
        this.StatusCode = s[0]
    }

    body := this.json(i)

    this.Send(body)
}

/*
    Send a JSON response with JSONP support.
    This method is identical to "res.Json()" however opts-in to JSONP callback support.
*/
func (this *Response) Jsonp(i interface{}, s ...int) {

    req := this.req
    app := this.app

    // If we were given a status, us it.
    if len(s) == 1 {
        this.StatusCode = s[0]
    }

    body := this.json(i)

    if cb, ok := req.Query[app.Get("jsonp callback name")]; ok {
        this.ContentType("text/javascript");
        body = cb + " && " + cb + "(" + body + ");";
    }

    this.Send(body)
}

/*
    Performs content-negotiation on the request Accept header field when present. 
    This method uses "req.Accepted", a slice of acceptable types ordered by their 
    quality values, otherwise the first callback is invoked. When no match is performed 
    the server responds with 406 "Not Acceptable", or invokes the "default" callback.
*/
func (this *Response) Format(i interface{}) {
    panic(halt)
}

/*
    Add `field` to Vary. If already present in the Vary set, then this call is simply ignored.
*/
func (this *Response) Vary() {
    panic(halt)
}

/*
    Sets the Content-Disposition header field to "attachment". If a filename is given then 
    the Content-Type will be automatically set based on the extname via res.Type(), 
    and the Content-Disposition's "filename=" parameter will be set.

    res.Attachment();
    // Content-Disposition: attachment

    res.Attachment('path/to/logo.png');
    // Content-Disposition: attachment; filename="logo.png"
    // Content-Type: image/png
*/
func (this *Response) Attachment(f ...string) {
    if len(f) > 0 {
        this.ContentType(filepath.Ext(f[0]))
        this.Set("Content-Disposition", "attachment; filename=\"" + filepath.Base(f[0]) + "\"")
    } else {
        this.Set("Content-Disposition", "attachment")
    }
}

/*
    Transfer the file at the given path. Automatically defaults the Content-Type response 
    header field based on the filename's extension.
*/
func (this *Response) Sendfile(p string, opt ...interface{}) {
    panic(halt)
}

/*
    Transfer the file at path as an "attachment", typically browsers will prompt the user for download. 
    The Content-Disposition "filename=" parameter, aka the one that will appear in the browser 
    dialog is set to path by default, however you may provide an override filename.
*/
func (this *Response) Download(p string, opt ...interface{}) {
    panic(halt)
}

/*
    Join the given "links" to populate the "Link" response header field.
*/
func (this *Response) Links(l []string) {
    panic(halt)
}

/*
    Render a "view". When an error occurs next(err) is invoked internally.
*/
func (this *Response) Render(view string, i ...interface{}) {
    i = append(i, this.Locals)
    s, err := this.app.Render(view, i...)
    if err != nil {
        this.next()
        return
    }
    this.Send(s)
}

/*
    Sets the Content-Type to the mime lookup of type, or when "/" is present the 
    Content-Type is simply set to this literal value.

    Examples:

         res.type('.html');
         res.type('html');
         res.type('json');
         res.type('application/json');
         res.type('png');
*/
func (this *Response) ContentType(t string) {
    if strings.Index(t, "/") == -1 {
        if t[:1] != "." {
            t = "." + t
        }
        t = mime.TypeByExtension(t)
    }
    this.Set("Content-Type", t)
}