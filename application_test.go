package f

import(
    // "fmt"
    "testing"
    . "github.com/ricallinson/simplebdd"
)

func TestApplication(t *testing.T) {

    Describe("CreateServer()", func() {

        It("should return [development]", func() {
            s := CreateServer()
            AssertEqual(s.Get("env"), "development")
        })
    })

    Describe("defaultConfiguration()", func() {

        It("should return default values", func() {
            s := &Server{
                settings: map[string]string{},
            }
            s.defaultConfiguration()
            AssertEqual(s.Enabled("x-powered-by"), true)
            AssertEqual(s.Enabled("etag"), true)
            AssertEqual(s.Get("env"), "development")
            // AssertEqual(s.Get("subdomain offset"), "2")
            AssertEqual(s.Get("jsonp callback name"), "callback")
        })
    })

    Describe("Set()", func() {

        It("should return [bar]", func() {
            s := CreateServer()
            s.Set("foo", "bar")
            AssertEqual(s.Get("foo"), "bar")
        })
    })

    Describe("Get()", func() {

        It("should return [bar]", func() {
            s := CreateServer()
            s.Set("foo", "bar")
            AssertEqual(s.Get("foo"), "bar")
        })
    })

    Describe("Enable()", func() {

        It("should return string [true]", func() {
            s := CreateServer()
            s.Enable("foo")
            AssertEqual(s.Get("foo"), TRUE)
        })
    })

    Describe("Disable()", func() {

        It("should return string [false]", func() {
            s := CreateServer()
            s.Disable("foo")
            AssertEqual(s.Get("foo"), FALSE)
        })
    })

    Describe("Enabled()", func() {

        It("should return [true & false]", func() {
            s := CreateServer()
            s.Enable("foo")
            AssertEqual(s.Enabled("foo"), true)
            AssertEqual(s.Enabled("bar"), false)
        })
    })

    Describe("Disabled()", func() {

        It("should return [true & false]", func() {
            s := CreateServer()
            s.Disable("foo")
            AssertEqual(s.Disabled("foo"), true)
            AssertEqual(s.Disabled("bar"), false)
        })
    })

    Describe("Engine()", func() {

        It("should return [true]", func() {
            s := CreateServer()
            s.Engine(".html", &MockRenderer{})
            _, ok := s.engines[".html"]
            AssertEqual(ok, true)
        })
    })

    Describe("Param()", func() {

        It("should return [true]", func() {
            s := CreateServer()
            s.Router.Middleware()
            s.Param(":user", func(req *Request, res *Response, next func()){})
            _, ok := s.Router.Params[":user"]
            AssertEqual(ok, true)
        })
    })

    Describe("Render()", func() {

        It("should return [skipped]", func() {
            AssertEqual(true, true)
        })
    })

    Describe("All()", func() {

        It("should return [true]", func() {
            s := CreateServer()
            s.All("/path", func(req *Request, res *Response, next func()){})
            AssertEqual(true, true)
        })
    })

    Describe("All Verbs Functions", func() {

        It("should return [true]", func() {
            s := CreateServer()
            s.Get("/path", func(req *Request, res *Response, next func()){})
            AssertEqual(true, true)
        })

        It("should return [true]", func() {
            s := CreateServer()
            s.Proppatch("/path", func(req *Request, res *Response, next func()){})
            AssertEqual(true, true)
        })

        It("should return [true]", func() {
            s := CreateServer()
            s.Post("/path", func(req *Request, res *Response, next func()){})
            AssertEqual(true, true)
        })

        It("should return [true]", func() {
            s := CreateServer()
            s.Put("/path", func(req *Request, res *Response, next func()){})
            AssertEqual(true, true)
        })

        It("should return [true]", func() {
            s := CreateServer()
            s.Head("/path", func(req *Request, res *Response, next func()){})
            AssertEqual(true, true)
        })

        It("should return [true]", func() {
            s := CreateServer()
            s.Delete("/path", func(req *Request, res *Response, next func()){})
            AssertEqual(true, true)
        })

        It("should return [true]", func() {
            s := CreateServer()
            s.Options("/path", func(req *Request, res *Response, next func()){})
            AssertEqual(true, true)
        })

        It("should return [true]", func() {
            s := CreateServer()
            s.Trace("/path", func(req *Request, res *Response, next func()){})
            AssertEqual(true, true)
        })

        It("should return [true]", func() {
            s := CreateServer()
            s.Copy("/path", func(req *Request, res *Response, next func()){})
            AssertEqual(true, true)
        })

        It("should return [true]", func() {
            s := CreateServer()
            s.Lock("/path", func(req *Request, res *Response, next func()){})
            AssertEqual(true, true)
        })

        It("should return [true]", func() {
            s := CreateServer()
            s.Move("/path", func(req *Request, res *Response, next func()){})
            AssertEqual(true, true)
        })

        It("should return [true]", func() {
            s := CreateServer()
            s.Propfind("/path", func(req *Request, res *Response, next func()){})
            AssertEqual(true, true)
        })

        It("should return [true]", func() {
            s := CreateServer()
            s.Unlock("/path", func(req *Request, res *Response, next func()){})
            AssertEqual(true, true)
        })

        It("should return [true]", func() {
            s := CreateServer()
            s.Report("/path", func(req *Request, res *Response, next func()){})
            AssertEqual(true, true)
        })

        It("should return [true]", func() {
            s := CreateServer()
            s.Mkactivity("/path", func(req *Request, res *Response, next func()){})
            AssertEqual(true, true)
        })

        It("should return [true]", func() {
            s := CreateServer()
            s.Mkcol("/path", func(req *Request, res *Response, next func()){})
            AssertEqual(true, true)
        })

        It("should return [true]", func() {
            s := CreateServer()
            s.Checkout("/path", func(req *Request, res *Response, next func()){})
            AssertEqual(true, true)
        })

        It("should return [true]", func() {
            s := CreateServer()
            s.Merge("/path", func(req *Request, res *Response, next func()){})
            AssertEqual(true, true)
        })

        It("should return [true]", func() {
            s := CreateServer()
            s.Msearch("/path", func(req *Request, res *Response, next func()){})
            AssertEqual(true, true)
        })

        It("should return [true]", func() {
            s := CreateServer()
            s.Notify("/path", func(req *Request, res *Response, next func()){})
            AssertEqual(true, true)
        })

        It("should return [true]", func() {
            s := CreateServer()
            s.Unsubscribe("/path", func(req *Request, res *Response, next func()){})
            AssertEqual(true, true)
        })

        It("should return [true]", func() {
            s := CreateServer()
            s.Patch("/path", func(req *Request, res *Response, next func()){})
            AssertEqual(true, true)
        })

        It("should return [true]", func() {
            s := CreateServer()
            s.Subscribe("/path", func(req *Request, res *Response, next func()){})
            AssertEqual(true, true)
        })
    })

    Report(t)
}