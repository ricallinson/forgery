package f

import(
    // "fmt"
    "testing"
    "net/url"
    "net/http"
    "github.com/ricallinson/stackr"
    . "github.com/ricallinson/simplebdd"
)

func TestRequest(t *testing.T) {

    var req *Request

    BeforeEach(func() {
        req = createRequest(
            &stackr.Request{
                Request: &http.Request{
                    URL: &url.URL{},
                },
            },
            &Server{},
        )
        req.Header = map[string][]string{}
    })

    Describe("Param()", func() {

        It("should return [bar]", func() {
            req.Params["foo"] = "bar"
            r := req.Param("foo")
            AssertEqual(r, "bar")
        })

        It("should return [bar]", func() {
            req.Body["foo"] = "bar"
            r := req.Param("foo")
            AssertEqual(r, "bar")
        })

        It("should return [bar]", func() {
            req.Query["foo"] = "bar"
            r := req.Param("foo")
            AssertEqual(r, "bar")
        })

        It("should return [bar]", func() {
            req.Params["foo"] = "bar"
            req.Body["foo"] = "bar1"
            req.Query["foo"] = "bar2"
            r := req.Param("foo")
            AssertEqual(r, "bar")
        })

        It("should return [bar]", func() {
            req.Body["foo"] = "bar"
            req.Query["foo"] = "bar1"
            r := req.Param("foo")
            AssertEqual(r, "bar")
        })

        It("should return []", func() {
            req.Params["foo"] = "bar"
            req.Body["foo"] = "bar1"
            req.Query["foo"] = "bar2"
            r := req.Param("bar")
            AssertEqual(r, "")
        })
    })

    Describe("Accepts()", func() {

        It("should return [skipped]", func() {
            AssertEqual(true, true)
        })
    })

    Describe("Is()", func() {

        It("should return [true]", func() {
            req.Header.Set("Content-Type", "text/html")
            AssertEqual(req.Is("text/html"), true)
        })

        It("should return [true]", func() {
            req.Header.Set("content-type", "text/html")
            AssertEqual(req.Is("text/html"), true)
        })
    })

    Describe("AcceptsCharset()", func() {

        It("should return [skipped]", func() {
            AssertEqual(true, true)
        })
    })

    Describe("AcceptsLanguage()", func() {

        It("should return [skipped]", func() {
            AssertEqual(true, true)
        })
    })

    Report(t)
}