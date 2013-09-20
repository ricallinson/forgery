package forgery

import(
    "github.com/ricallinson/stackr"
)

/*
    Response represents the response from an HTTP request.
*/
type Response struct {

    // The stackr.Response.
    *stackr.Response

    // Assign the charset. Defaults to "utf-8".
    Charset string

    // Response local variables are scoped to the request, thus only available 
    // to the view(s) rendered during that request / response cycle, if any. 
    // Otherwise this API is identical to app.locals.
    Locals string
}

/*
    Returns a new HTTP Response.
*/

func createResponse(res *stackr.Response) (*Response) {

    /*
        Create the Response.
    */

    r := &Response{}

    /*
        Set the stackr.Response.
    */

    r.Response = res

    /*
        Return the finished forgery.Response.
    */

    return r
}

/*
    Chainable alias of stackr's "res.StatusCode=".
*/
func (this *Response) Status(c int) (*Response) {
    return this
}

/*
    Set header "field" to "value".
*/
func (this *Response) Set(f string, v string) {

}

/*
    Get the case-insensitive response header "field".
*/
func (this *Response) Get(f string) {

}

/*
    Set cookie "name" to "value", where "value" may be a string or interface 
    converted to JSON. The "path" option defaults to "/".
*/
func (this *Response) Cookie(n string, v string, opt ...interface{}) {

}

/*
    Clear cookie "name". The "path" option defaults to "/".
*/
func (this *Response) ClearCookie(n string, opt ...interface{}) {

}

/*
    Redirect to the given "url" with optional "status" code defaulting to 302 "Found".
*/
func (this *Response) Redirect(url string, s ...int) {

}

/*
    Set the location header.
*/
func (this *Response) Location(url string) {

}

/*
    Send a response. This method performs a myriad of useful tasks for simple non-streaming 
    responses such as automatically assigning the Content-Length unless previously defined 
    and providing automatic HEAD and HTTP cache freshness support.

    res.Send(new Buffer('whoop'));
    res.Send(interface{ "some": "json" });
    res.Send('some html');
    res.Send('Sorry, we cannot find that!', 404);
    res.Send(interface{ "error": 'something blew up' }, 500);
    res.Send(200);
*/
func (this *Response) Send(b interface{}, s ...int) {

}

/*
    Send a JSON response. This method is identical to res.Send() when an object or slice is passed, 
    however it may be used for explicit JSON conversion of non-objects 
    (null, undefined, etc), though these are technically not valid JSON.

    res.Json(null)
    res.Json({ "user": 'tobi' })
    res.Json(interface{ error: 'message' }, 500)
*/
func (this *Response) Json(b interface{}, s ...int) {

}

/*
    Send a JSON response with JSONP support.
    This method is identical to "res.Json()" however opts-in to JSONP callback support.
*/
func (this *Response) Jsonp(b interface{}, s ...int) {

}

/*
    Sets the Content-Type to the mime lookup of type, or when "/" is present the 
    Content-Type is simply set to this literal value.
*/
func (this *Response) Type(t string) {

}

/*
    Performs content-negotiation on the request Accept header field when present. 
    This method uses "req.accepted", an array of acceptable types ordered by their 
    quality values, otherwise the first callback is invoked. When no match is performed 
    the server responds with 406 "Not Acceptable", or invokes the "default" callback.
*/
func (this *Response) Format(i interface{}) {

}

/*
    Sets the Content-Disposition header field to "attachment". If a filename is given then 
    the Content-Type will be automatically set based on the extname via res.type(), 
    and the Content-Disposition's "filename=" parameter will be set.

    res.Attachment();
    // Content-Disposition: attachment

    res.Attachment('path/to/logo.png');
    // Content-Disposition: attachment; filename="logo.png"
    // Content-Type: image/png
*/
func (this *Response) Attachment(f ...string) {

}

/*
    Transfer the file at the given path. Automatically defaults the Content-Type response 
    header field based on the filename's extension.
*/
func (this *Response) Sendfile(p string, opt ...interface{}) {

}

/*
    Transfer the file at path as an "attachment", typically browsers will prompt the user for download. 
    The Content-Disposition "filename=" parameter, aka the one that will appear in the brower 
    dialog is set to path by default, however you may provide an override filename.
*/
func (this *Response) Download(p string, opt ...interface{}) {

}

/*
    Join the given "links" to populate the "Link" response header field.
*/
func (this *Response) Links(l []string) {

}

/*
    Render a "view". When an error occurs next(err) is invoked internally.
*/
func (this *Response) Render(v string, l ...interface{}) {
    
}