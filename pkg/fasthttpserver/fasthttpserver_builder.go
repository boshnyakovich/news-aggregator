package fasthttpserver

import (
	"net"
	"time"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
)

type Config struct {
	ReduceMemoryUsage bool
	ReadBufferSize    int
	WriteBufferSize   int
	MaxRequestBodySize int
	Addr              string
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration

	listener net.Listener
}

type Builder struct {
	config *Config

	listener net.Listener

	livenessPath    string
	livenessHandler fasthttp.RequestHandler

	readinessPath    string
	readinessHandler fasthttp.RequestHandler
}

func New() *Builder {
	return new(Builder)
}

func (v *Builder) WithConfig(config *Config) *Builder {
	v.config = config

	return v
}

func (v *Builder) WithLivenessHandler(path string, handler fasthttp.RequestHandler) *Builder {
	v.livenessPath = path
	v.livenessHandler = handler

	return v
}

func (v *Builder) WithDefaultLivenessHandler() *Builder {
	v.livenessPath = "/liveness"
	v.livenessHandler = LivenessHandler

	return v
}

func (v *Builder) WithReadinessHandler(path string, handler fasthttp.RequestHandler) *Builder {
	v.readinessPath = path
	v.readinessHandler = handler

	return v
}

func (v *Builder) WithDefaultReadinessHandler() *Builder {
	v.readinessPath = "/readiness"
	v.readinessHandler = ReadinessHandler

	return v
}

func (v *Builder) WithListener(listener net.Listener) *Builder {
	v.listener = listener

	return v
}

func (v *Builder) Build() (srv *FastHTTPServer, err error) {
	if v.config == nil {
		return nil, errors.New("config is empty")
	}

	v.config.listener = v.listener

	srv = newFastHTTPServer(v.config)

	if v.livenessPath != "" {
		srv.Router().GET(v.livenessPath, v.livenessHandler)
	}

	if v.readinessPath != "" {
		srv.Router().GET(v.readinessPath, v.readinessHandler)
	}

	return srv, nil
}
