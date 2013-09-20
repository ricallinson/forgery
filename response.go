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
}

/*
    Returns a new HTTP Response.
*/

func createResponse(res *stackr.Response) (*Response) {

    /*
        Return the finished forgery.Response.
    */

    return &Response{res}
}