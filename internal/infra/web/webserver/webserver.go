package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type HandlerMethod struct {
	method  string
	path    string
	handler http.HandlerFunc
}

type WebServer struct {
	Router             chi.Router
	Handlers           map[string]http.HandlerFunc
	HandlersWithMethod []HandlerMethod
	WebServerPort      string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:             chi.NewRouter(),
		Handlers:           make(map[string]http.HandlerFunc),
		HandlersWithMethod: []HandlerMethod{},
		WebServerPort:      serverPort,
	}
}

func (s *WebServer) AddHandler(path string, handler http.HandlerFunc) {
	s.Handlers[path] = handler
}

func (s *WebServer) AddHandlerWithMethod(method string, path string, handler http.HandlerFunc) {
	s.HandlersWithMethod = append(s.HandlersWithMethod, HandlerMethod{
		method:  method,
		path:    path,
		handler: handler,
	})
}

// loop through the handlers and add them to the router
// register middleware logger
// start the server
func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for path, handler := range s.Handlers {
		s.Router.Handle(path, handler)
	}

	for _, route := range s.HandlersWithMethod {
		s.Router.Method(route.method, route.path, route.handler)
	}

	http.ListenAndServe(s.WebServerPort, s.Router)
}
