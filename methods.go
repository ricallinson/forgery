package forgery

/*
    The method provides the routing functionality for POST requests to the given "path".
*/
func (this *Server) Post(path string, fn ...func(*Request, *Response, func())) {
    this.Verb("POST", path, fn...);
}

/*
    The method provides the routing functionality for PUT requests to the given "path".
*/
func (this *Server) Put(path string, fn ...func(*Request, *Response, func())) {
    this.Verb("PUT", path, fn...);
}

/*
    The method provides the routing functionality for HEAD requests to the given "path".
*/
func (this *Server) Head(path string, fn ...func(*Request, *Response, func())) {
    this.Verb("HEAD", path, fn...);
}

/*
    The method provides the routing functionality for DELETE requests to the given "path".
*/
func (this *Server) Delete(path string, fn ...func(*Request, *Response, func())) {
    this.Verb("DELETE", path, fn...);
}

/*
    The method provides the routing functionality for OPTIONS requests to the given "path".
*/
func (this *Server) Options(path string, fn ...func(*Request, *Response, func())) {
    this.Verb("OPTIONS", path, fn...);
}

/*
    The method provides the routing functionality for TRACE requests to the given "path".
*/
func (this *Server) Trace(path string, fn ...func(*Request, *Response, func())) {
    this.Verb("TRACE", path, fn...);
}

/*
    The method provides the routing functionality for COPY requests to the given "path".
*/
func (this *Server) Copy(path string, fn ...func(*Request, *Response, func())) {
    this.Verb("COPY", path, fn...);
}

/*
    The method provides the routing functionality for LOCK requests to the given "path".
*/
func (this *Server) Lock(path string, fn ...func(*Request, *Response, func())) {
    this.Verb("LOCK", path, fn...);
}

/*
    The method provides the routing functionality for MKCOL requests to the given "path".
*/
func (this *Server) Mkcol(path string, fn ...func(*Request, *Response, func())) {
    this.Verb("MKCOL", path, fn...);
}

/*
    The method provides the routing functionality for MOVE requests to the given "path".
*/
func (this *Server) Move(path string, fn ...func(*Request, *Response, func())) {
    this.Verb("MOVE", path, fn...);
}

/*
    The method provides the routing functionality for PROPFIND requests to the given "path".
*/
func (this *Server) Propfind(path string, fn ...func(*Request, *Response, func())) {
    this.Verb("PROPFIND", path, fn...);
}

/*
    The method provides the routing functionality for PROPPATCH requests to the given "path".
*/
func (this *Server) Proppatch(path string, fn ...func(*Request, *Response, func())) {
    this.Verb("PROPPATCH", path, fn...);
}

/*
    The method provides the routing functionality for UNLOCK requests to the given "path".
*/
func (this *Server) Unlock(path string, fn ...func(*Request, *Response, func())) {
    this.Verb("UNLOCK", path, fn...);
}

/*
    The method provides the routing functionality for REPORT requests to the given "path".
*/
func (this *Server) Report(path string, fn ...func(*Request, *Response, func())) {
    this.Verb("REPORT", path, fn...);
}

/*
    The method provides the routing functionality for MKACTIVITY requests to the given "path".
*/
func (this *Server) Mkactivity(path string, fn ...func(*Request, *Response, func())) {
    this.Verb("MKACTIVITY", path, fn...);
}

/*
    The method provides the routing functionality for CHECKOUT requests to the given "path".
*/
func (this *Server) Checkout(path string, fn ...func(*Request, *Response, func())) {
    this.Verb("CHECKOUT", path, fn...);
}

/*
    The method provides the routing functionality for MERGE requests to the given "path".
*/
func (this *Server) Merge(path string, fn ...func(*Request, *Response, func())) {
    this.Verb("MERGE", path, fn...);
}

/*
    The method provides the routing functionality for M-SEARCH requests to the given "path".
*/
func (this *Server) Msearch(path string, fn ...func(*Request, *Response, func())) {
    this.Verb("M-SEARCH", path, fn...);
}

/*
    The method provides the routing functionality for NOTIFY requests to the given "path".
*/
func (this *Server) Notify(path string, fn ...func(*Request, *Response, func())) {
    this.Verb("NOTIFY", path, fn...);
}

/*
    The method provides the routing functionality for SUBSCRIBE requests to the given "path".
*/
func (this *Server) Subscribe(path string, fn ...func(*Request, *Response, func())) {
    this.Verb("SUBSCRIBE", path, fn...);
}

/*
    The method provides the routing functionality for UNSUBSCRIBE requests to the given "path".
*/
func (this *Server) Unsubscribe(path string, fn ...func(*Request, *Response, func())) {
    this.Verb("UNSUBSCRIBE", path, fn...);
}

/*
    The method provides the routing functionality for PATCH requests to the given "path".
*/
func (this *Server) Patch(path string, fn ...func(*Request, *Response, func())) {
    this.Verb("PATCH", path, fn...);
}