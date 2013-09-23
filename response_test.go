package f

import(
    "bytes"
    "testing"
    "net/url"
    "net/http"
    "github.com/ricallinson/stackr"
    . "github.com/ricallinson/simplebdd"
)

func TestResponse(t *testing.T) {

    var res *Response
    var mock *MockResponseWriter

    BeforeEach(func() {
        mock = NewMockResponseWriter(false)
        res = createResponse(
            &Request{
                Request: &stackr.Request{
                    Request: &http.Request{
                        URL: &url.URL{},
                    },
                },
                Query: map[string]string{},
            },
            &stackr.Response{
                Writer: mock,
            },
            &Server{
                settings: map[string]string{},
            },
        )
    })

    Describe("Status()", func() {

        It("should return [true]", func() {
            res.Status(404)
            AssertEqual(res.StatusCode, 404)
        })
    })

    Describe("Get()", func() {

        It("should return [bar]", func() {
            res.Set("foo", "bar")
            AssertEqual(res.Get("foo"), "bar")
        })
    })

    Describe("Set()", func() {

        It("should return [bar]", func() {
            res.Set("foo", "bar")
            AssertEqual(res.Get("foo"), "bar")
        })
    })

    Describe("Cookie()", func() {

        It("should return [skipped]", func() {
            AssertEqual(true, true)
        })
    })

    Describe("ClearCookie()", func() {

        It("should return [skipped]", func() {
            AssertEqual(true, true)
        })
    })

    Describe("Redirect()", func() {

        It("should return [skipped]", func() {
            AssertEqual(true, true)
        })
    })

    Describe("Location()", func() {

        It("should return [skipped]", func() {
            AssertEqual(true, true)
        })
    })

    Describe("Send()", func() {

        It("should return [404 & text/plain & 9 & utf-8 & Not Found]", func() {
            res.Send(404)
            w := bytes.NewBuffer(mock.Written).String()
            AssertEqual(res.StatusCode, 404)
            AssertEqual(res.Get("Content-Type"), "text/plain")
            AssertEqual(res.Get("Content-Length"), "9")
            AssertEqual(res.Charset, "utf-8")
            AssertEqual(w, "Not Found")
        })

        It("should return [200 & text/html & 3 & utf-8 & foo]", func() {
            res.Send("foo")
            res.StatusCode = 200
            w := bytes.NewBuffer(mock.Written).String()
            AssertEqual(res.StatusCode, 200)
            AssertEqual(res.Get("Content-Type"), "text/html")
            AssertEqual(res.Get("Content-Length"), "3")
            AssertEqual(res.Charset, "utf-8")
            AssertEqual(w, "foo")
        })

        It("should return [500 & text/html & 3 & utf-8 & foo]", func() {
            res.Send("foo", 500)
            w := bytes.NewBuffer(mock.Written).String()
            AssertEqual(res.StatusCode, 500)
            AssertEqual(res.Get("Content-Type"), "text/html")
            AssertEqual(res.Get("Content-Length"), "3")
            AssertEqual(res.Charset, "utf-8")
            AssertEqual(w, "foo")
        })

        It("should return [200 & application/json & 3 & utf-8 & {\"foo\":\"bar\"}}]", func() {
            res.Send(map[string]string{"foo": "bar"})
            res.StatusCode = 200
            w := bytes.NewBuffer(mock.Written).String()
            AssertEqual(res.StatusCode, 200)
            AssertEqual(res.Get("Content-Type"), "application/json")
            AssertEqual(res.Get("Content-Length"), "13")
            AssertEqual(res.Charset, "utf-8")
            AssertEqual(w, "{\"foo\":\"bar\"}")
        })

        It("should return [500 & application/json & 3 & utf-8 & {\"foo\":\"bar\"}}]", func() {
            res.Send(map[string]string{"foo": "bar"}, 500)
            w := bytes.NewBuffer(mock.Written).String()
            AssertEqual(res.StatusCode, 500)
            AssertEqual(res.Get("Content-Type"), "application/json")
            AssertEqual(res.Get("Content-Length"), "13")
            AssertEqual(res.Charset, "utf-8")
            AssertEqual(w, "{\"foo\":\"bar\"}")
        })

        It("should return [204 & utf-8]", func() {
            res.Send(204)
            w := bytes.NewBuffer(mock.Written).String()
            AssertEqual(res.StatusCode, 204)
            AssertEqual(res.Get("Content-Type"), "")
            AssertEqual(res.Get("Content-Length"), "")
            AssertEqual(res.Get("Transfer-Encoding"), "")
            AssertEqual(res.Charset, "utf-8")
            AssertEqual(w, "")
        })

        It("should return [204 & utf-8]", func() {
            res.Send("foo", 204)
            w := bytes.NewBuffer(mock.Written).String()
            AssertEqual(res.StatusCode, 204)
            AssertEqual(res.Get("Content-Type"), "")
            AssertEqual(res.Get("Content-Length"), "")
            AssertEqual(res.Get("Transfer-Encoding"), "")
            AssertEqual(res.Charset, "utf-8")
            AssertEqual(w, "")
        })

        It("should return [304 & utf-8]", func() {
            res.Send(304)
            w := bytes.NewBuffer(mock.Written).String()
            AssertEqual(res.StatusCode, 304)
            AssertEqual(res.Get("Content-Type"), "")
            AssertEqual(res.Get("Content-Length"), "")
            AssertEqual(res.Get("Transfer-Encoding"), "")
            AssertEqual(res.Charset, "utf-8")
            AssertEqual(w, "")
        })

        It("should return [200 & utf-8]", func() {
            res.req.Method = "HEAD"
            res.Send("foo", 200)
            w := bytes.NewBuffer(mock.Written).String()
            AssertEqual(res.StatusCode, 200)
            AssertEqual(res.Get("Content-Type"), "text/html")
            AssertEqual(res.Get("Content-Length"), "3")
            AssertEqual(res.Get("Transfer-Encoding"), "")
            AssertEqual(res.Charset, "utf-8")
            AssertEqual(w, "")
        })

        It("should return [utf-8]", func() {
            res.Charset = ""
            res.Send("foo", 200)
            AssertEqual(res.Charset, "utf-8")
        })

        It("should return [304]", func() {
            res.req.Fresh = true
            res.Send("foo", 200)
            w := bytes.NewBuffer(mock.Written).String()
            AssertEqual(res.StatusCode, 304)
            AssertEqual(res.Get("Content-Type"), "")
            AssertEqual(res.Get("Content-Length"), "")
            AssertEqual(res.Get("Transfer-Encoding"), "")
            AssertEqual(w, "")
        })
    })
    
    Describe("json()", func() {

        It("should return [{\"foo\":\"bar\"}]", func() {
            res.Charset = ""
            j := res.json(map[string]string{"foo": "bar"})
            AssertEqual(j, "{\"foo\":\"bar\"}")
        })
    })

    Describe("Json()", func() {

        It("should return [{\"foo\":\"bar\"}]", func() {
            res.Json(map[string]string{"foo": "bar"}, 500)
            w := bytes.NewBuffer(mock.Written).String()
            AssertEqual(res.StatusCode, 500)
            AssertEqual(res.Get("Content-Type"), "application/json")
            AssertEqual(res.Get("Content-Length"), "13")
            AssertEqual(res.Charset, "utf-8")
            AssertEqual(w, "{\"foo\":\"bar\"}")
        })
    })

    Describe("Jsonp()", func() {

        It("should return [{\"foo\":\"bar\"}]", func() {
            res.Jsonp(map[string]string{"foo": "bar"}, 500)
            w := bytes.NewBuffer(mock.Written).String()
            AssertEqual(res.StatusCode, 500)
            AssertEqual(res.Get("Content-Type"), "application/json")
            AssertEqual(res.Get("Content-Length"), "13")
            AssertEqual(res.Charset, "utf-8")
            AssertEqual(w, "{\"foo\":\"bar\"}")
        })

        It("should return [{\"foo\":\"bar\"}]", func() {
            res.app.Set("jsonp callback name", "callback")
            res.req.Query["callback"] = "cb"
            res.Jsonp(map[string]string{"foo": "bar"}, 500)
            w := bytes.NewBuffer(mock.Written).String()
            AssertEqual(res.StatusCode, 500)
            AssertEqual(res.Get("Content-Type"), "text/javascript")
            AssertEqual(res.Get("Content-Length"), "24")
            AssertEqual(res.Charset, "utf-8")
            AssertEqual(w, "cb && cb({\"foo\":\"bar\"});")
        })
    })

    Describe("Format()", func() {

        It("should return [skipped]", func() {
            AssertEqual(true, true)
        })
    })

    Describe("Attachment()", func() {

        It("should return [skipped]", func() {
            AssertEqual(true, true)
        })
    })

    Describe("Sendfile()", func() {

        It("should return [skipped]", func() {
            AssertEqual(true, true)
        })
    })

    Describe("Download()", func() {

        It("should return [skipped]", func() {
            AssertEqual(true, true)
        })
    })

    Describe("Links()", func() {

        It("should return [skipped]", func() {
            AssertEqual(true, true)
        })
    })

    Describe("Render()", func() {

        It("should return [skipped]", func() {
            AssertEqual(true, true)
        })
    })

    Describe("ContentType()", func() {

        It("should return [foo]", func() {
            res.ContentType("foo/bar")
            AssertEqual(res.Get("content-type"), "foo/bar")
        })

        It("should return [foo]", func() {
            res.ContentType("foo/bar")
            AssertEqual(res.Get("Content-Type"), "foo/bar")
        })
    })

    Report(t)
}