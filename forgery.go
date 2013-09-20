package forgery

import(
    "github.com/ricallinson/stackr"
)

/*
    Create forgery Request, Response types.
*/
func Make(req *stackr.Request, res *stackr.Response) (*Request, *Response) {
    return createRequest(req), createResponse(res)
}