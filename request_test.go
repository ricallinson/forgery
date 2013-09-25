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

        It("should return [true]", func() {
            req.Header.Set("Accept", "text/html")
            req.Accepts("text/html")
            AssertEqual(req.Accepts("text/html"), true)
        })

        It("should return [true]", func() {
            req.Header.Set("Accept", "text/html, text/plain")
            req.Accepts("text/html")
            AssertEqual(req.Accepts("text/plain"), true)
        })

        It("should return [true]", func() {
            req.Header.Set("Accept", "text/html, text/plain")
            req.Accepts("text/html")
            AssertEqual(req.Accepts("Text/Plain"), true)
        })

        It("should return [true]", func() {
            req.Header.Set("Accept", "text/html, text/plain")
            req.Accepts("text/html")
            AssertEqual(req.Accepts("text"), false)
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

        It("should return [true]", func() {
            req.Header.Set("Accept-Charset", "utf-8")
            AssertEqual(req.AcceptsCharset("utf-8"), true)
        })

        It("should return [true]", func() {
            req.Header.Set("Accept-Charset", "utf-8")
            AssertEqual(req.AcceptsCharset("UTF-8"), true)
        })

        It("should return [true]", func() {
            req.Header.Set("Accept-Charset", "utf-8")
            AssertEqual(req.AcceptsCharset("ISO-8859-1"), false)
        })
    })

    Describe("AcceptsLanguage()", func() {

        It("should return [true]", func() {
            req.Header.Set("Accept-Language", "zh, en-us; q=0.8, en; q=0.6")
            AssertEqual(req.AcceptsLanguage("en"), true)
        })

        It("should return [true]", func() {
            req.Header.Set("Accept-Language", "zh, en-us; q=0.8, en; q=0.6")
            AssertEqual(req.AcceptsLanguage("en-US"), true)
        })

        It("should return [true]", func() {
            req.Header.Set("Accept-Language", "zh, en-us; q=0.8, en; q=0.6")
            AssertEqual(req.AcceptsLanguage("fr-CA"), false)
        })
    })

    Describe("Fresh()", func() {

        It("should return [false]", func() {
            req.Method = "DELETE"
            a := req.Fresh(100)
            AssertEqual(a, false)
        })

        It("should return [false]", func() {
            req.Method = "GET"
            a := req.Fresh(100)
            AssertEqual(a, false)
        })

        It("should return [false]", func() {
            req.Method = "GET"
            a := req.Fresh(200)
            AssertEqual(a, false)
        })

        It("should return [false]", func() {
            req.Method = "HEAD"
            a := req.Fresh(200)
            AssertEqual(a, false)
        })
    })

    Describe("Stale()", func() {

        It("should return [true]", func() {
            a := req.Stale(100)
            AssertEqual(a, true)
        })
    })

    Report(t)
}