package forgery

import(
    "github.com/ricallinson/stackr"
)

/*
    A Request represents an HTTP request received by the server.
*/
type Request struct {

    // The stackr.Request type.
    *stackr.Request
}

/*
    Returns a new Request.
*/

func createRequest(req *stackr.Request) (*Request) {

    /*
        Return the finished forgery.Request.
    */

    return &Request{req}
}