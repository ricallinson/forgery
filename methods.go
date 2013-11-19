package f

/*
   The method provides the routing functionality for POST requests to the given "path".
*/
func (this *Server) Post(path string, fn ...func(*Request, *Response, func())) {
	this.Verb("post", path, fn...)
}

/*
   The method provides the routing functionality for PUT requests to the given "path".
*/
func (this *Server) Put(path string, fn ...func(*Request, *Response, func())) {
	this.Verb("put", path, fn...)
}

/*
   The method provides the routing functionality for HEAD requests to the given "path".
*/
func (this *Server) Head(path string, fn ...func(*Request, *Response, func())) {
	this.Verb("head", path, fn...)
}

/*
   The method provides the routing functionality for DELETE requests to the given "path".
*/
func (this *Server) Delete(path string, fn ...func(*Request, *Response, func())) {
	this.Verb("delete", path, fn...)
}

/*
   The method provides the routing functionality for OPTIONS requests to the given "path".
*/
func (this *Server) Options(path string, fn ...func(*Request, *Response, func())) {
	this.Verb("options", path, fn...)
}

/*
   The method provides the routing functionality for TRACE requests to the given "path".
*/
func (this *Server) Trace(path string, fn ...func(*Request, *Response, func())) {
	this.Verb("trace", path, fn...)
}

/*
   The method provides the routing functionality for COPY requests to the given "path".
*/
func (this *Server) Copy(path string, fn ...func(*Request, *Response, func())) {
	this.Verb("copy", path, fn...)
}

/*
   The method provides the routing functionality for LOCK requests to the given "path".
*/
func (this *Server) Lock(path string, fn ...func(*Request, *Response, func())) {
	this.Verb("lock", path, fn...)
}

/*
   The method provides the routing functionality for MKCOL requests to the given "path".
*/
func (this *Server) Mkcol(path string, fn ...func(*Request, *Response, func())) {
	this.Verb("mkcol", path, fn...)
}

/*
   The method provides the routing functionality for MOVE requests to the given "path".
*/
func (this *Server) Move(path string, fn ...func(*Request, *Response, func())) {
	this.Verb("move", path, fn...)
}

/*
   The method provides the routing functionality for PROPFIND requests to the given "path".
*/
func (this *Server) Propfind(path string, fn ...func(*Request, *Response, func())) {
	this.Verb("propfind", path, fn...)
}

/*
   The method provides the routing functionality for PROPPATCH requests to the given "path".
*/
func (this *Server) Proppatch(path string, fn ...func(*Request, *Response, func())) {
	this.Verb("proppatch", path, fn...)
}

/*
   The method provides the routing functionality for UNLOCK requests to the given "path".
*/
func (this *Server) Unlock(path string, fn ...func(*Request, *Response, func())) {
	this.Verb("unlock", path, fn...)
}

/*
   The method provides the routing functionality for REPORT requests to the given "path".
*/
func (this *Server) Report(path string, fn ...func(*Request, *Response, func())) {
	this.Verb("report", path, fn...)
}

/*
   The method provides the routing functionality for MKACTIVITY requests to the given "path".
*/
func (this *Server) Mkactivity(path string, fn ...func(*Request, *Response, func())) {
	this.Verb("mkactivity", path, fn...)
}

/*
   The method provides the routing functionality for CHECKOUT requests to the given "path".
*/
func (this *Server) Checkout(path string, fn ...func(*Request, *Response, func())) {
	this.Verb("checkout", path, fn...)
}

/*
   The method provides the routing functionality for MERGE requests to the given "path".
*/
func (this *Server) Merge(path string, fn ...func(*Request, *Response, func())) {
	this.Verb("merge", path, fn...)
}

/*
   The method provides the routing functionality for M-SEARCH requests to the given "path".
*/
func (this *Server) Msearch(path string, fn ...func(*Request, *Response, func())) {
	this.Verb("m-search", path, fn...)
}

/*
   The method provides the routing functionality for NOTIFY requests to the given "path".
*/
func (this *Server) Notify(path string, fn ...func(*Request, *Response, func())) {
	this.Verb("notify", path, fn...)
}

/*
   The method provides the routing functionality for SUBSCRIBE requests to the given "path".
*/
func (this *Server) Subscribe(path string, fn ...func(*Request, *Response, func())) {
	this.Verb("subscribe", path, fn...)
}

/*
   The method provides the routing functionality for UNSUBSCRIBE requests to the given "path".
*/
func (this *Server) Unsubscribe(path string, fn ...func(*Request, *Response, func())) {
	this.Verb("unsubscribe", path, fn...)
}

/*
   The method provides the routing functionality for PATCH requests to the given "path".
*/
func (this *Server) Patch(path string, fn ...func(*Request, *Response, func())) {
	this.Verb("patch", path, fn...)
}
