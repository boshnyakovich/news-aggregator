package fasthttpserver

import (
	"context"
	"github.com/fasthttp/router"
	"github.com/prometheus/common/log"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/pprofhandler"
	"net"
)

type FastHTTPServer struct {
	addr string

	listener net.Listener

	server *fasthttp.Server
	router *router.Router
}

func newFastHTTPServer(config *Config) (srv *FastHTTPServer) {
	srv = &FastHTTPServer{
		addr: config.Addr,

		listener: config.listener,

		server: &fasthttp.Server{
			ReadBufferSize:     config.ReadBufferSize,
			ReadTimeout:        config.ReadTimeout,
			WriteBufferSize:    config.WriteBufferSize,
			WriteTimeout:       config.WriteTimeout,
			ReduceMemoryUsage:  config.ReduceMemoryUsage,
			MaxRequestBodySize: config.MaxRequestBodySize,
		},
		router: router.New(),
	}

	if config.listener != nil {
		srv.addr = config.listener.Addr().String()
	}

	srv.server.Handler = srv.router.Handler

	srv.router.GET("/debug/pprof", pprofhandler.PprofHandler)

	return srv
}

func (v *FastHTTPServer) Router() *router.Router {
	return v.router
}

func (v FastHTTPServer) Address() string {
	return v.addr
}

func (v *FastHTTPServer) Start(ctx context.Context) (err error) {
	curCtx, cancel := context.WithCancel(ctx)

	go func() {
		<-curCtx.Done()
		if err := v.server.Shutdown(); err != nil {
			log.Warn("cannot shutdown server")
		}
		cancel()
	}()

	if v.listener != nil {
		return v.server.Serve(v.listener)
	}

	return v.server.ListenAndServe(v.addr)
}

func (v *FastHTTPServer) GracefulStop() error {
	if v.listener != nil {
		return v.listener.Close()
	}

	return v.server.Shutdown()
}
