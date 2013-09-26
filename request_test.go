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
    var mock *MockResponseWriter

    BeforeEach(func() {
        mock = NewMockResponseWriter(false)
        req = createRequest(
            &stackr.Request{
                Request: &http.Request{
                    URL: &url.URL{},
                    Header: map[string][]string{},
                },
            },
            &Server{
                settings: map[string]string{},
            },
        )
        req.res = &Response{
            Response: &stackr.Response{
                Writer: mock,
            },
        }
    })

    Describe("createRequest()", func() {

        It("should return [true]", func() {
            app := &Server{
                settings: map[string]string{},
            }
            httpReq := &http.Request{
                URL: &url.URL{},
                Header: map[string][]string{},
            }

            app.Enable("trust proxy")
            httpReq.Header.Set("X-Forwarded-For", "129.78.138.66, 129.78.64.103")

            req := createRequest(
                &stackr.Request{
                    Request: httpReq,
                },
                app,
            )
            AssertEqual(req.Ips[1], "129.78.64.103")
        })
    })

    Describe("Cookie()", func() {

        It("should return [bar]", func() {
            req.Header.Set("Cookie", "foo=bar;")
            f := req.Cookie("foo")
            AssertEqual(f, "bar")
        })

        It("should return [bar]", func() {
            req.Header.Set("Cookie", "foo=" + url.QueryEscape("{\"foo\":\"bar\"}"))
            var f map[string]interface{}
            t := req.Cookie("foo", &f)
            AssertEqual(t, "{\"foo\":\"bar\"}")
            AssertEqual(f["foo"], "bar")
        })
    })

    Describe("SignedCookie()", func() {

        It("should return [bar]", func() {
            req.app.Set("secret", "word")
            req.Header.Set("Cookie", "foo=bar;")
            f := req.SignedCookie("foo")
            AssertEqual(f, "bar")
        })

        It("should return [bar]", func() {
            req.app.Set("secret", "word")
            req.Header.Set("Cookie", "foo=" + url.QueryEscape("{\"foo\":\"bar\"}"))
            var f map[string]interface{}
            t := req.SignedCookie("foo", &f)
            AssertEqual(t, "{\"foo\":\"bar\"}")
            AssertEqual(f["foo"], "bar")
        })
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
            req.res.StatusCode = 100
            a := req.Fresh()
            AssertEqual(a, false)
        })

        It("should return [false]", func() {
            req.Method = "GET"
            req.res.StatusCode = 100
            a := req.Fresh()
            AssertEqual(a, false)
        })

        It("should return [false]", func() {
            req.Method = "GET"
            req.res.StatusCode = 200
            a := req.Fresh()
            AssertEqual(a, false)
        })

        It("should return [false]", func() {
            req.Method = "HEAD"
            req.res.StatusCode = 200
            a := req.Fresh()
            AssertEqual(a, false)
        })
    })

    Describe("Stale()", func() {

        It("should return [true]", func() {
            req.res.StatusCode = 100
            a := req.Stale()
            AssertEqual(a, true)
        })
    })

    Report(t)
}