package main

import (
	"context"
	"net/http"
)

func main() {

}

type App struct {
	cbs []ShutdownCallBacks
}

type ServerMux struct {
	*http.ServeMux
	reject bool
}

func (s *ServerMux) ServerHttp(w http.ResponseWriter, r *http.Request) {
	if s.reject {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("服务已经关闭"))
		return
	}

	s.ServeMux.ServeHTTP(w, r)
}

type ShutdownCallBacks func(ctx context.Context)

type Option func(app *App)

func WithShutdownCallBacks(cbs ...ShutdownCallBacks) Option {
	return func(app *App) {
		app.cbs = cbs
	}
}
