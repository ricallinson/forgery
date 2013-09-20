package forgery

import(
    "github.com/ricallinson/stackr"
)

type Server struct {
    *stackr.Server
}

/*
    Create a new stackr server.
*/
func CreateServer() (*Server) {
    return &Server{&stackr.Server{}}
}

/*
    Create forgery Request, Response types from stackr Request, Response types.
*/
func make(req *stackr.Request, res *stackr.Response) (*Request, *Response) {
    return createRequest(req), createResponse(res)
}