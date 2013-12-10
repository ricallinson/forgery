# Forgery

[![Build Status](https://secure.travis-ci.org/ricallinson/forgery.png?branch=master)](http://travis-ci.org/ricallinson/forgery)

__WARNING: WORK IN PROGRESS__

Forgery is a minimal and flexible golang web application framework, providing a robust set of features for building single and multi-page, web applications.

    package main

    import(
        "github.com/ricallinson/forgery"
    )

    func init() {
        app := f.CreateServer()
        app.Get("/", func(req *f.Request, res *f.Response, next func()) {
            res.Send("Hello world.")
        })
        app.Listen(3000)
    }

* Built on [Stackr](https://github.com/ricallinson/stackr)
* Robust routing (almost there)
* HTTP helpers (redirection, caching, etc)
* View system supporting 1 template engine (hopefully more will come)
* Content negotiation
* Focus on high performance
* Environment based configuration
* High test coverage

## Testing

The following should all be executed from the `forgery` directory _$GOPATH/src/github.com/ricallinson/forgery/_.

#### Install

    go get github.com/ricallinson/simplebdd

#### Run

    go test

### Code Coverage

#### Install

    go get github.com/axw/gocov/gocov
    go get -u github.com/matm/gocov-html

#### Generate

    gocov test | gocov-html > ./reports/coverage.html

## Notes

This project started out as a clone of the superb Node.js library [Express](http://expressjs.com/).
