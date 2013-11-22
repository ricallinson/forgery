package f

import (
	"bytes"
	. "github.com/ricallinson/simplebdd"
	"github.com/ricallinson/stackr"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestResponse(t *testing.T) {

	var res *Response
	var mock *MockResponseWriter

	BeforeEach(func() {
		mock = NewMockResponseWriter(false)
		res = createResponse(
			&stackr.Response{
				Writer: mock,
			},
			func() {},
			&Server{
				settings: map[string]string{},
			},
		)
		res.req = &Request{
			Request: &stackr.Request{
				Request: &http.Request{
					URL:    &url.URL{},
					Header: map[string][]string{},
				},
				Query: map[string]string{},
			},
		}
		res.req.res = res // Crazy shit!
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

		It("should clear the header", func() {
			res.Set("foo", "bar")
			res.Set("foo", "")
			AssertEqual(res.Get("foo"), "")
		})
	})

	Describe("Cookie()", func() {

		It("should return [foo=bar;]", func() {
			res.Cookie("foo", "bar")
			h := res.Get("Set-Cookie")
			// AssertEqual(h, "")
			AssertEqual(strings.Index(h, "foo=YmFy;") > -1, true)
			AssertEqual(strings.Index(h, "Path=/;") > -1, true)
		})

		It("should return [Path=/foo]", func() {
			res.Cookie("foo", "bar", &http.Cookie{Path: "/foo"})
			h := res.Get("Set-Cookie")
			// AssertEqual(h, "")
			AssertEqual(strings.Index(h, "foo=YmFy;") > -1, true)
			AssertEqual(strings.Index(h, "Path=/foo;") > -1, true)
		})

		It("should return [{\"foo\":\"bar\"};]", func() {
			res.Cookie("foo", map[string]string{"foo": "bar"})
			h := res.Get("Set-Cookie")
			// AssertEqual(h, "")
			AssertEqual(strings.Index(h, "foo=eyJmb28iOiJiYXIifQ%3D%3D;") > -1, true)
			AssertEqual(strings.Index(h, "Path=/;") > -1, true)
		})
	})

	Describe("SignedCookie()", func() {

		It("should return [foo=bar;]", func() {
			res.app.Set("secret", "wordwordwordword")
			res.SignedCookie("foo", "bar")
			h := res.Get("Set-Cookie")
			// AssertEqual(h, "")
			AssertEqual(strings.Index(h, "foo=YmFyLld2WHdGQVBpaDNuQllfWUJhWWp3MmlONmN6VTFqam5MNjU1ZHZrcnFjbE09;") > -1, true)
			AssertEqual(strings.Index(h, "Path=/;") > -1, true)
		})

		It("should return [{\"foo\":\"bar\"};]", func() {
			res.app.Set("secret", "wordwordwordword")
			res.SignedCookie("foo", map[string]string{"foo": "bar"})
			h := res.Get("Set-Cookie")
			// AssertEqual(h, "")
			AssertEqual(strings.Index(h, "foo=eyJmb28iOiJiYXIifS5QU1hjUGdOS3NwZFR6Q3BmOW1qN2JFR2RTUUx3MU5nWTRkMkE2QXpFTktjPQ%3D%3D;") > -1, true)
			AssertEqual(strings.Index(h, "Path=/;") > -1, true)
		})
	})

	Describe("ClearCookie()", func() {

		It("should return [foo=]", func() {
			res.ClearCookie("foo")
			h := res.Get("Set-Cookie")
			AssertEqual(strings.Index(h, "Path=/;") > -1, true)
			AssertEqual(strings.Index(h, "Max-Age=0") > -1, true)
		})

		It("should return [Path=/foo]", func() {
			res.ClearCookie("foo", &http.Cookie{Path: "/foo"})
			h := res.Get("Set-Cookie")
			// AssertEqual(h, "")
			AssertEqual(strings.Index(h, "Path=/foo;") > -1, true)
			AssertEqual(strings.Index(h, "Max-Age=0") > -1, true)
		})

		It("should return [foo=bar]", func() {

			/*
			   If the cookie has already been set, it cannot be deleted.
			*/

			res.Cookie("foo", "bar")
			res.ClearCookie("foo")
			h := res.Get("Set-Cookie")
			AssertEqual(strings.Index(h, "foo=YmFy;") > -1, true)
			AssertEqual(strings.Index(h, "Path=/;") > -1, true)
			AssertEqual(strings.Index(h, "Max-Age=0") > -1, false)
		})
	})

	Describe("Redirect()", func() {

		It("should return []", func() {
			res.Redirect("http://www.foo.com/")
			w := bytes.NewBuffer(mock.Written).String()
			AssertEqual(w, "")
		})

		It("should return [Moved Temporarily. Redirecting to http://www.foo.com/]", func() {
			res.req.Header.Set("Accept", "text/plain")
			res.Redirect("http://www.foo.com/")
			w := bytes.NewBuffer(mock.Written).String()
			AssertEqual(w, "Moved Temporarily. Redirecting to http://www.foo.com/")
		})

		It("should return [HTML <p>...]", func() {
			res.req.Header.Set("Accept", "text/html")
			res.Redirect("http://www.foo.com/")
			w := bytes.NewBuffer(mock.Written).String()
			AssertEqual(w, "<p>Moved Temporarily. Redirecting to <a href=\"http://www.foo.com/\">http://www.foo.com/</a></p>")
		})

		It("should return [Not Found. Redirecting to http://www.foo.com/]", func() {
			res.req.Header.Set("Accept", "text/plain")
			res.Redirect("http://www.foo.com/", 404)
			w := bytes.NewBuffer(mock.Written).String()
			AssertEqual(w, "Not Found. Redirecting to http://www.foo.com/")
		})

		It("should return [Not Found. Redirecting to http://www.foo.com/]", func() {
			res.req.Method = "HEAD"
			res.req.Header.Set("Accept", "text/plain")
			res.Redirect("http://www.foo.com/")
			w := bytes.NewBuffer(mock.Written).String()
			AssertEqual(w, "")
		})
	})

	Describe("Location()", func() {

		It("should return [http://www.foo.com/]", func() {
			res.Location("http://www.foo.com/")
			AssertEqual(res.Get("location"), "http://www.foo.com/")
		})

		It("should return [/bar/baz]", func() {
			res.Location("/bar/baz")
			AssertEqual(res.Get("location"), "/bar/baz")
		})

		It("should return [/bar/baz]", func() {
			res.req.OriginalUrl = "http://www.foo.com/"
			res.app.Set("app path", "/")
			res.Location("bar/baz")
			AssertEqual(res.Get("location"), "/bar/baz")
		})

		It("should return [/foo/bar/baz]", func() {
			res.req.OriginalUrl = "http://www.foo.com"
			res.app.Set("app path", "/foo/")
			res.Location("bar/baz")
			AssertEqual(res.Get("location"), "/foo/bar/baz")
		})

		It("should return [http://www.foo.com/bar/baz]", func() {
			res.req.OriginalUrl = "http://www.foo.com/"
			res.Location("./bar/baz")
			AssertEqual(res.Get("location"), "http://www.foo.com/bar/baz")
		})

		It("should return [http://www.foo.com/bar/baz]", func() {
			res.req.OriginalUrl = "http://www.foo.com"
			res.Location("./bar/baz")
			AssertEqual(res.Get("location"), "http://www.foo.com/bar/baz")
		})

		It("should return [http://www.foo.com/bar/baz]", func() {
			res.req.OriginalUrl = "http://www.foo.com/?foo=bar"
			res.Location("./bar/baz")
			AssertEqual(res.Get("location"), "http://www.foo.com/bar/baz")
		})

		It("should return [http://www.foo.com/bar/baz]", func() {
			res.req.OriginalUrl = "http://www.foo.com?foo=bar"
			res.Location("./bar/baz")
			AssertEqual(res.Get("location"), "http://www.foo.com/bar/baz")
		})

		It("should return [http://www.foo.com/bar/baz]", func() {
			res.req.OriginalUrl = "http://www.foo.com/foo/bar"
			res.Location("../baz")
			AssertEqual(res.Get("location"), "http://www.foo.com/foo/baz")
		})

		It("should return [http://www.foo.com/]", func() {
			res.req.Header.Set("Referer", "http://www.foo.com/")
			res.Location("back")
			AssertEqual(res.Get("location"), "http://www.foo.com/")
		})

		It("should return [/]", func() {
			res.Location("back")
			AssertEqual(res.Get("location"), "/")
		})
	})

	Describe("Send()", func() {

		It("should return [404 & text/plain & 9 & utf-8 & Not Found]", func() {
			res.Send(404)
			w := bytes.NewBuffer(mock.Written).String()
			AssertEqual(res.StatusCode, 404)
			AssertEqual(res.Get("Content-Type"), "text/plain; charset=utf-8")
			AssertEqual(res.Get("Content-Length"), "9")
			AssertEqual(res.Charset, "utf-8")
			AssertEqual(w, "Not Found")
		})

		It("should return [200 & text/html & 3 & utf-8 & foo]", func() {
			res.Send("foo")
			res.StatusCode = 200
			w := bytes.NewBuffer(mock.Written).String()
			AssertEqual(res.StatusCode, 200)
			AssertEqual(res.Get("Content-Type"), "text/html; charset=utf-8")
			AssertEqual(res.Get("Content-Length"), "3")
			AssertEqual(res.Charset, "utf-8")
			AssertEqual(w, "foo")
		})

		It("should return [500 & text/html & 3 & utf-8 & foo]", func() {
			res.Send("foo", 500)
			w := bytes.NewBuffer(mock.Written).String()
			AssertEqual(res.StatusCode, 500)
			AssertEqual(res.Get("Content-Type"), "text/html; charset=utf-8")
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

		It("should return [200 & text/html & 3 & utf-8 & foo]", func() {
			res.Charset = ""
			res.Send(bytes.NewBufferString("foo").Bytes(), 200)
			w := bytes.NewBuffer(mock.Written).String()
			AssertEqual(res.Get("Content-Type"), "text/html; charset=utf-8")
			AssertEqual(res.Get("Content-Length"), "3")
			AssertEqual(res.Get("Transfer-Encoding"), "")
			AssertEqual(res.Charset, "utf-8")
			AssertEqual(w, "foo")
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
			AssertEqual(res.Get("Content-Type"), "text/html; charset=utf-8")
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
			res.req.Header.Set("if-none-match", "*")
			res.req.Method = "GET"
			res.Send("foo", 100)
			w := bytes.NewBuffer(mock.Written).String()
			AssertEqual(res.StatusCode, 304)
			AssertEqual(res.Get("Content-Type"), "")
			AssertEqual(res.Get("Content-Length"), "")
			AssertEqual(res.Get("Transfer-Encoding"), "")
			AssertEqual(w, "")
		})
	})

	Describe("Etag()", func() {

		It("should return [2356372769]", func() {
			t := res.Etag("foo")
			AssertEqual(t, "2356372769")
		})
	})

	Describe("json()", func() {

		It("should return [{\"foo\":\"bar\"}]", func() {
			res.Charset = ""
			j := res.json(map[string]string{"foo": "bar"})
			AssertEqual(j, "{\"foo\":\"bar\"}")
		})

		It("should return [{\"foo\":\"bar\"}]", func() {
			res.Charset = ""
			res.app.Set("env", "development")
			res.app.Set("json spaces", "  ")
			j := res.json(map[string]string{"foo": "bar"})
			AssertEqual(j, "{\n  \"foo\": \"bar\"\n}")
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

		It("should return [true]", func() {
			test := false
			formats := map[string]func(){
				"default": func() {
					test = true
				},
			}
			res.Format(formats)
			AssertEqual(test, true)
		})

		It("should return [406]", func() {
			test := false
			formats := map[string]func(){
				"none": func() {
					test = true
				},
			}
			res.Format(formats)
			w := bytes.NewBuffer(mock.Written).String()
			AssertEqual(res.StatusCode, 406)
			AssertEqual(w, "Not Acceptable")
		})

		It("should return [text/html]", func() {
			test := ""
			formats := map[string]func(){
				"default": func() {
					test = "default"
				},
				"text/html": func() {
					test = "text/html"
				},
			}
			res.req.Header.Set("Accept", "text/html")
			res.Format(formats)
			AssertEqual(test, "text/html")
		})
	})

	Describe("Attachment()", func() {

		It("should return [attachment]", func() {
			res.Attachment()
			AssertEqual(res.Get("Content-Type"), "")
			AssertEqual(res.Get("Content-Disposition"), "attachment")
		})

		It("should return [image/png & attachment; filename=\"logo.png\"]", func() {
			res.Attachment("path/to/logo.png")
			AssertEqual(res.Get("Content-Type"), "image/png")
			AssertEqual(res.Get("Content-Disposition"), "attachment; filename=\"logo.png\"")
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

		It("should return [<http://api.example.com/users?page=2>; rel=\"next\"]", func() {
			res.Links("http://api.example.com/users?page=2", "next")
			l := res.Get("Link")
			AssertEqual(l, "<http://api.example.com/users?page=2>; rel=\"next\"")
		})

		It("should return [rel=\"next\", rel=\"last\"]", func() {
			res.Links("http://api.example.com/users?page=2", "next")
			res.Links("http://api.example.com/users?page=5", "last")
			l := res.Get("Link")
			AssertEqual(l, "<http://api.example.com/users?page=2>; rel=\"next\", <http://api.example.com/users?page=5>; rel=\"last\"")
		})
	})

	Describe("Vary()", func() {

		It("should return [Accept]", func() {
			res.Vary("Accept")
			v := res.Get("Vary")
			AssertEqual(v, "Accept")
		})

		It("should return [Accept, *]", func() {
			res.Vary("Accept")
			res.Vary("*")
			v := res.Get("Vary")
			AssertEqual(v, "Accept, *")
		})

		It("should return [Accept, *]", func() {
			res.Vary("Accept")
			res.Vary("*")
			res.Vary("*")
			res.Vary("Accept")
			v := res.Get("Vary")
			AssertEqual(v, "Accept, *")
		})
	})

	Describe("Render()", func() {

		It("should return [skipped]", func() {
			AssertEqual(true, true)
		})
	})

	Describe("ContentType()", func() {

		It("should return [foo/bar]", func() {
			res.ContentType("foo/bar")
			AssertEqual(res.Get("content-type"), "foo/bar")
		})

		It("should return [foo/bar]", func() {
			res.ContentType("foo/bar")
			AssertEqual(res.Get("Content-Type"), "foo/bar")
		})

		It("should return [image/png]", func() {
			res.ContentType(".png")
			AssertEqual(res.Get("Content-Type"), "image/png")
		})

		It("should return [text/html; charset=utf-8]", func() {
			res.ContentType(".html")
			AssertEqual(res.Get("Content-Type"), "text/html; charset=utf-8")
		})

		It("should return [image/jpeg]", func() {
			res.ContentType("jpg")
			AssertEqual(res.Get("Content-Type"), "image/jpeg")
		})
	})

	Report(t)
}
