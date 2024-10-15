package misugo

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
)

type MisugoApp struct {
	server *http.ServeMux
}
type MisugoHandler struct {
	handler http.Handler
}

func (m *MisugoHandler) Next(ctx *ContextRequest) {
	m.handler.ServeHTTP(ctx.w, ctx.r)
}

func NewMisugoHandler(h http.Handler) *MisugoHandler {
	return &MisugoHandler{handler: h}
}

func NewApp() *MisugoApp {
	return &MisugoApp{
		server: http.NewServeMux(),
	}
}

var contextPool = sync.Pool{
	New: func() interface{} {
		return &ContextRequest{}
	},
}

func (m *MisugoApp) Serve(port string) error {
	server := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: m.server,
	}

	return server.ListenAndServe()
}

func (m *MisugoApp) Get(pattern string, handler func(*ContextRequest) error, middlewares ...func(*MisugoHandler) *MisugoHandler) {
	// Ensure the pattern starts with "/"
	if !strings.HasPrefix(pattern, "/") {
		pattern = "/" + pattern
	}

	finalHandler := NewMisugoHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := contextPool.Get().(*ContextRequest)
		w.Header().Set("Content-Type", "application/json")
		ctx.w = w
		ctx.r = r
		defer contextPool.Put(ctx)

		if err := handler(ctx); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}))

	if len(middlewares) != 0 {
		for i := len(middlewares) - 1; i >= 0; i-- {
			finalHandler = middlewares[i](finalHandler)
		}
	}

	m.server.HandleFunc(fmt.Sprintf("GET %s", pattern), finalHandler.handler.ServeHTTP)
}

func (m *MisugoApp) Post(pattern string, handler func(*ContextRequest) error, middlewares ...func(*MisugoHandler) *MisugoHandler) {
	// cannot have a pattern without /
	if !strings.HasPrefix(pattern, "/") {
		pattern = "/" + pattern
	}
	finalHandler := NewMisugoHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := contextPool.Get().(*ContextRequest)
		w.Header().Set("Content-Type", "application/json")
		ctx.w = w
		ctx.r = r
		defer contextPool.Put(ctx)

		if err := handler(ctx); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}))

	if len(middlewares) != 0 {
		for i := len(middlewares) - 1; i >= 0; i-- {
			finalHandler = middlewares[i](finalHandler)
		}
	}

	m.server.HandleFunc(fmt.Sprintf("POST %s", pattern), finalHandler.handler.ServeHTTP)
}

func (m *MisugoApp) Delete(pattern string, handler func(*ContextRequest) error, middlewares ...func(*MisugoHandler) *MisugoHandler) {
	if !strings.HasPrefix(pattern, "/") {
		pattern = "/" + pattern
	}
	finalHandler := NewMisugoHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := contextPool.Get().(*ContextRequest)
		w.Header().Set("Content-Type", "application/json")
		ctx.w = w
		ctx.r = r
		defer contextPool.Put(ctx)

		if err := handler(ctx); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}))

	if len(middlewares) != 0 {
		for i := len(middlewares) - 1; i >= 0; i-- {
			finalHandler = middlewares[i](finalHandler)
		}
	}

	m.server.HandleFunc(fmt.Sprintf("DELETE %s", pattern), finalHandler.handler.ServeHTTP)
}

func (m *MisugoApp) Put(pattern string, handler func(*ContextRequest) error, middlewares ...func(*MisugoHandler) *MisugoHandler) {
	if !strings.HasPrefix(pattern, "/") {
		pattern = "/" + pattern
	}
	finalHandler := NewMisugoHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := contextPool.Get().(*ContextRequest)
		w.Header().Set("Content-Type", "application/json")
		ctx.w = w
		ctx.r = r
		defer contextPool.Put(ctx)

		if err := handler(ctx); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}))

	if len(middlewares) != 0 {
		for i := len(middlewares) - 1; i >= 0; i-- {
			finalHandler = middlewares[i](finalHandler)
		}
	}

	m.server.HandleFunc(fmt.Sprintf("DELETE %s", pattern), finalHandler.handler.ServeHTTP)
}

func (m *MisugoApp) Patch(pattern string, handler func(*ContextRequest) error, middlewares ...func(*MisugoHandler) *MisugoHandler) {
	if !strings.HasPrefix(pattern, "/") {
		pattern = "/" + pattern
	}
	finalHandler := NewMisugoHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := contextPool.Get().(*ContextRequest)
		w.Header().Set("Content-Type", "application/json")
		ctx.w = w
		ctx.r = r
		defer contextPool.Put(ctx)

		if err := handler(ctx); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}))

	if len(middlewares) != 0 {
		for i := len(middlewares) - 1; i >= 0; i-- {
			finalHandler = middlewares[i](finalHandler)
		}
	}

	m.server.HandleFunc(fmt.Sprintf("DELETE %s", pattern), finalHandler.handler.ServeHTTP)
}
