package f

import (
	// "fmt"
	. "github.com/ricallinson/simplebdd"
	"testing"
)

func TestRouter(t *testing.T) {

	Describe("&Router{}", func() {

		It("should add a route", func() {
			r := &Router{}
			r.AddRoute("GET", "/", func(req *Request, res *Response, next func()) {
				// ...
			})
			AssertEqual(len(r.Routes), 1)
		})

		It("should add 2 route", func() {
			r := &Router{}
			r.AddRoute("GET", "/", func(req *Request, res *Response, next func()) {
				// ...
			})
			r.AddRoute("GET", "/", func(req *Request, res *Response, next func()) {
				// ...
			})
			AssertEqual(len(r.Routes), 2)
		})
	})

	Report(t)
}
