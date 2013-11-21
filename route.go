package f

import (
    // "fmt"
    "regexp"
    "strings"
)

type Route struct {

    // HTTP Method
    Method string

    // URL Path
    Path string

    // The regex used to match the route.
    Regex *regexp.Regexp

    // Slice of functions
    Funcs []func(*Request, *Response, func())

    // The slice of param names to return from the path.
    ParamNames []string
}

func (this *Route) Match(method string, path string) (map[string]string, bool) {

    // If there is no Regex, compile one.
    if this.Regex == nil {
        this.Regex = this.CompileRegex(this.Path)
    }

    // Prime the params map.
    params := map[string]string{}

    // Test if this route matches the given "method" and "path".
    if method != this.Method || this.Regex.MatchString(path) == false {
        return params, false
    }

    // Extract all params from the matched path.
    paramsList := this.Regex.FindAllStringSubmatch(path, -1)

    // For each found param, add it to the params map.
    for i, name := range this.ParamNames {
        params[name] = paramsList[0][i + 1]
    }

    // Debug this stuff.
    // fmt.Println(path)
    // fmt.Println(this.Path)
    // fmt.Println(paramsList)
    // fmt.Println(params)
    // fmt.Println("")

    // Return the complete params map.
    return params, true
}

/*
    Parse the given path to return a regex and slice of param names.

    Path: /foo/:param1/bar/:param2/baz
    => ^/foo/(.*)/bar/(.*)/baz$
    => [":param1", :param2]
*/
func (this *Route) CompileRegex(path string) (*regexp.Regexp) {

    // Compile the the param finder.
    finder := regexp.MustCompile(`:[a-zA-Z0-9]+`)

    // Find param names.
    this.ParamNames = finder.FindAllString(path, -1)

    // Clean the param names we got.
    for i, name := range this.ParamNames {
        this.ParamNames[i] = name[1:]
    }

    // For each param name found, replace it with a regex group.
    for _, param := range this.ParamNames {
        path = strings.Replace(path, ":" + param, "([^/]+)", 1)
    }

    // If the path does not end in with "*", then force a line end match.
    if path[len(path) - 1:] != "*" {
        path = path + "$"
    }

    // Build the regexp used for all future request to this route.
    return regexp.MustCompile("^" + path)
}