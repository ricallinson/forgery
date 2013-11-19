package f

import (
	// "fmt"
	// "strconv"
	. "github.com/ricallinson/simplebdd"
	"github.com/ricallinson/stackr"
	"net/http"
	"net/url"
	"testing"
)

func TestRequest(t *testing.T) {

	var req *Request
	var mock *MockResponseWriter

	BeforeEach(func() {
		mock = NewMockResponseWriter(false)
		req = createRequest(
			&stackr.Request{
				Request: &http.Request{
					URL:    &url.URL{},
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
				URL:    &url.URL{},
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

		It("should return []", func() {
			f := req.Cookie("foo")
			AssertEqual(f, "")
		})

		It("should return [bar]", func() {
			req.Header.Set("Cookie", "foo=YmFy")
			f := req.Cookie("foo")
			AssertEqual(f, "bar")
		})

		It("should return [bar]", func() {
			req.Header.Set("Cookie", "foo=eyJmb28iOiJiYXIifQ%3D%3D")
			var f map[string]interface{}
			t := req.Cookie("foo", &f)
			AssertEqual(t, "{\"foo\":\"bar\"}")
			AssertEqual(f["foo"], "bar")
		})
	})

	Describe("SignedCookie()", func() {

		It("should return [bar]", func() {
			req.app.Set("secret", "wordwordwordword")
			req.Header.Set("Cookie", "foo=YmFyLld2WHdGQVBpaDNuQllfWUJhWWp3MmlONmN6VTFqam5MNjU1ZHZrcnFjbE09")
			f := req.SignedCookie("foo")
			AssertEqual(f, "bar")
		})

		It("should return [bar]", func() {
			req.app.Set("secret", "wordwordwordword")
			req.Header.Set("Cookie", "foo=eyJmb28iOiJiYXIifS5QU1hjUGdOS3NwZFR6Q3BmOW1qN2JFR2RTUUx3MU5nWTRkMkE2QXpFTktjPQ%3D%3D")
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

	Describe("Subdomains()", func() {

		It("should return [0]", func() {
			req.app.Set("subdomain offset", "2")
			a := req.Subdomains()
			AssertEqual(len(a), 0)
		})

		It("should return [a, b, c]", func() {
			req.app.Set("subdomain offset", "2")
			req.Host = "a.b.c.domain.com"
			a := req.Subdomains()
			AssertEqual(a[0], "c")
			AssertEqual(a[1], "b")
			AssertEqual(a[2], "a")
			AssertEqual(len(a), 3)
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

	Describe("Accepted()", func() {

		It("should return [application/xml]", func() {
			req.Header.Set("Accept", "text/html, application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
			a := req.Accepted()
			AssertEqual(a[2], "application/xml")
		})
	})

	Describe("AcceptedLanguages()", func() {

		It("should return [en-us]", func() {
			req.Header.Set("Accept-Language", "zh, en-us; q=0.8, en; q=0.6")
			a := req.AcceptedLanguages()
			AssertEqual(a[1], "en-us")
		})
	})

	Describe("AcceptedCharsets()", func() {

		It("should return [utf-8]", func() {
			req.Header.Set("Accept-Charset", "utf-8")
			a := req.AcceptedCharsets()
			AssertEqual(a[0], "utf-8")
		})
	})

	Describe("processAccepted()", func() {

		It("should return [text/plain]", func() {
			a := req.processAccepted("text/plain")
			AssertEqual(a[0], "text/plain")
		})

		It("should return [application/json]", func() {
			a := req.processAccepted("text/plain, text/html,application/json , image/png")
			AssertEqual(a[2], "application/json")
		})

		It("should return [application/json]", func() {
			a := req.processAccepted("application/xml,application/xhtml+xml,text/html;q=0.9,text/plain;q=0.8,image/png,*/*;q=0.5")
			AssertEqual(a[2], "image/png")
		})
	})

	Report(t)
}
