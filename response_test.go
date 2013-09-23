package f

import(
    // "fmt"
    "testing"
    "github.com/ricallinson/stackr"
    . "github.com/ricallinson/simplebdd"
)

func TestResponse(t *testing.T) {

    var res *Response

    BeforeEach(func() {
        res = createResponse(
        	&Request{},
            &stackr.Response{},
            &Server{},
        )
    })

    Describe("Status()", func() {

        It("should return [true]", func() {
        	res.Status(404)
            AssertEqual(res.StatusCode, 404)
        })
    })

    Report(t)
}