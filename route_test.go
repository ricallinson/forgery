package f

import (
    // "fmt"
    . "github.com/ricallinson/simplebdd"
    "testing"
)

func TestRoute(t *testing.T) {

    Describe("&Route{}", func() {

        It("should match the route", func() {
            r := &Route{
                Method: "GET",
                Path: "/",
            }
            _, ok := r.Match("GET", "/")
            AssertEqual(ok, true)
        })

        It("should match the route", func() {
            r := &Route{
                Method: "GET",
                Path: "/foo",
            }
            _, ok := r.Match("GET", "/foo")
            AssertEqual(ok, true)
        })

        It("should match the route", func() {
            r := &Route{
                Method: "GET",
                Path: "*",
            }
            _, ok := r.Match("GET", "/foo")
            AssertEqual(ok, true)
        })

        It("should match the route with a wildcard ending", func() {
            r := &Route{
                Method: "GET",
                Path: "/*",
            }
            _, ok := r.Match("GET", "/foo")
            AssertEqual(ok, true)
        })

        It("should match the route with a wildcard ending", func() {
            r := &Route{
                Method: "GET",
                Path: "/foo*",
            }
            _, ok := r.Match("GET", "/foo/bar")
            AssertEqual(ok, true)
        })

        It("should match the route with a wildcard ending", func() {
            r := &Route{
                Method: "GET",
                Path: "/foo/*",
            }
            _, ok := r.Match("GET", "/foo/bar")
            AssertEqual(ok, true)
        })

        It("should match the route with a param", func() {
            r := &Route{
                Method: "GET",
                Path: "/:param",
            }
            p, ok := r.Match("GET", "/foo")
            AssertEqual(ok, true)
            AssertEqual(p["param"], "foo")
        })

        It("should match the route with a param", func() {
            r := &Route{
                Method: "GET",
                Path: "/:param/bar",
            }
            p, ok := r.Match("GET", "/foo/bar")
            AssertEqual(ok, true)
            AssertEqual(p["param"], "foo")
        })

        It("should match the route with 2 params", func() {
            r := &Route{
                Method: "GET",
                Path: "/:param1/bar/:param2",
            }
            p, ok := r.Match("GET", "/foo/bar/baz")
            AssertEqual(ok, true)
            AssertEqual(p["param1"], "foo")
            AssertEqual(p["param2"], "baz")
        })

        It("should match the route with 2 params", func() {
            r := &Route{
                Method: "GET",
                Path: "/:param1/bar/:param2/qux",
            }
            p, ok := r.Match("GET", "/foo/bar/baz/qux")
            AssertEqual(ok, true)
            AssertEqual(p["param1"], "foo")
            AssertEqual(p["param2"], "baz")
        })

        It("should match the route with 2 params and wildcard", func() {
            r := &Route{
                Method: "GET",
                Path: "/:param1/bar/:param2/*",
            }
            p, ok := r.Match("GET", "/foo/bar/baz/qux/xuq")
            AssertEqual(ok, true)
            AssertEqual(p["param1"], "foo")
            AssertEqual(p["param2"], "baz")
        })

        It("should match the route with 2 params and no slashes", func() {
            r := &Route{
                Method: "GET",
                Path: "/:param1-bar-:param2-xuq",
            }
            p, ok := r.Match("GET", "/foo-bar-baz-qux-xuq")
            AssertEqual(ok, true)
            AssertEqual(p["param1"], "foo")
            AssertEqual(p["param2"], "baz-qux")
        })

        It("should NOT match the route", func() {
            r := &Route{
                Method: "GET",
                Path: "foo",
            }
            _, ok := r.Match("GET", "/foo")
            AssertEqual(ok, false)
        })

        It("should NOT match the route", func() {
            r := &Route{
                Method: "GET",
                Path: "foo*",
            }
            _, ok := r.Match("GET", "/foo")
            AssertEqual(ok, false)
        })

        It("should NOT match the route", func() {
            r := &Route{
                Method: "GET",
                Path: ":param1",
            }
            _, ok := r.Match("GET", "/foo")
            AssertEqual(ok, false)
        })
    })

    Report(t)
}