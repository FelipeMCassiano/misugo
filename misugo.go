package misugo

import (
	"fmt"
	"net/http"
	"strings"
	"sync"

	ContextRequest "github.com/FelipeMCassiano/misugo/contextRequest"
)

type MisugoApp struct {
	server *http.ServeMux
}

func NewApp() *MisugoApp {
	return &MisugoApp{
		server: http.NewServeMux(),
	}
}

type MisugoHandler func(*ContextRequest.ContextRequest)

var contextPool = sync.Pool{
	New: func() interface{} {
		return &ContextRequest.ContextRequest{}
	},
}

func (m *MisugoApp) Get(pattern string, handler MisugoHandler) {
	// cannot have a pattern without /
	if !strings.HasPrefix(pattern, "/") {
		pattern = "/" + pattern
	}

	m.server.HandleFunc(fmt.Sprintf("GET %s", pattern), func(w http.ResponseWriter, r *http.Request) {
		ctx := contextPool.Get().(*ContextRequest.ContextRequest)
		ctx.W = w
		ctx.R = r
		defer contextPool.Put(ctx)

		handler(ctx)
	})
}

func (m *MisugoApp) Post(pattern string, handler MisugoHandler) {
	// cannot have a pattern without /
	if !strings.HasPrefix(pattern, "/") {
		pattern = "/" + pattern
	}

	m.server.HandleFunc(fmt.Sprintf("POST %s", pattern), func(w http.ResponseWriter, r *http.Request) {
		ctx := contextPool.Get().(*ContextRequest.ContextRequest)
		ctx.W = w
		ctx.R = r
		defer contextPool.Put(ctx)

		handler(ctx)
	})
}

func (m *MisugoApp) Delete(pattern string, handler MisugoHandler) {
	if !strings.HasPrefix(pattern, "/") {
		pattern = "/" + pattern
	}

	m.server.HandleFunc(fmt.Sprintf("DELETE %s", pattern), func(w http.ResponseWriter, r *http.Request) {
		ctx := contextPool.Get().(*ContextRequest.ContextRequest)
		ctx.W = w
		ctx.R = r
		defer contextPool.Put(ctx)

		handler(ctx)
	})
}

func (m *MisugoApp) Put(pattern string, handler MisugoHandler) {
	if !strings.HasPrefix(pattern, "/") {
		pattern = "/" + pattern
	}

	m.server.HandleFunc(fmt.Sprintf("PUT %s", pattern), func(w http.ResponseWriter, r *http.Request) {
		ctx := contextPool.Get().(*ContextRequest.ContextRequest)
		ctx.W = w
		ctx.R = r
		defer contextPool.Put(ctx)

		handler(ctx)
	})
}

func (m *MisugoApp) Patch(pattern string, handler MisugoHandler) {
	if !strings.HasPrefix(pattern, "/") {
		pattern = "/" + pattern
	}

	m.server.HandleFunc(fmt.Sprintf("PATCH %s", pattern), func(w http.ResponseWriter, r *http.Request) {
		ctx := contextPool.Get().(*ContextRequest.ContextRequest)
		ctx.W = w
		ctx.R = r
		defer contextPool.Put(ctx)

		handler(ctx)
	})
}
