package http

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

type Params map[string]string
type Body map[string]string
type Controller func(params Params, body Body) Response

type Route struct {
	method     string
	pattern    *regexp.Regexp
	controller Controller
	paramKeys  []string
}

type Router struct {
	Routes []Route
}

func (router *Router) Add(method string, endpoint string, controller Controller) {
	paramKeys := []string{}
	pattern := regexp.MustCompile(":([a-zA-Z]+)")
	matches := pattern.FindAllStringSubmatch(endpoint, -1)

	if len(matches) > 0 {
		endpoint = pattern.ReplaceAllLiteralString(endpoint, "([^/]+)")
		for i := 0; i < len(matches); i++ {
			paramKeys = append(paramKeys, matches[i][1])
		}
	}

	newRoute := Route{method, regexp.MustCompile("^" + endpoint + "$"), controller, paramKeys}
	router.Routes = append(router.Routes, newRoute)
}

func (router *Router) GET(endpoint string, controller Controller) {
	router.Add(MethodGet, endpoint, controller)
}

func (router *Router) POST(endpoint string, controller Controller) {
	router.Add(MethodPost, endpoint, controller)
}

func (router *Router) PUT(endpoint string, controller Controller) {
	router.Add(MethodGet, endpoint, controller)
}

func (router *Router) PATCH(endpoint string, controller Controller) {
	router.Add(MethodPatch, endpoint, controller)
}

func (router *Router) DELETE(endpoint string, controller Controller) {
	router.Add(MethodDelete, endpoint, controller)
}

func (router *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	allow := []string{}

	for i := 0; i < len(router.Routes); i++ {
		var route Route = router.Routes[i]

		matches := route.pattern.FindStringSubmatch(req.URL.Path)
		if len(matches) == 0 {
			continue
		}
		if req.Method != route.method {
			allow = append(allow, route.method)
			continue
		}

		var params Params = ParseRequestParams(route.paramKeys, matches[1:])
		var body Body = ParseRequestBody(req.Body)
		var response Response = route.controller(params, body)

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(response.StatusCode)
		w.Write(response.Body)

		return
	}

	if len(allow) > 0 {
		w.Header().Set("Allow", strings.Join(allow, ", "))
		w.WriteHeader(StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(StatusNotFound)
	fmt.Fprintln(w, "404 page not found")
}
