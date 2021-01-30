package fasthttpserver

import (
	"bufio"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttputil"
)

func TestReadinessHandler(t *testing.T) {
	t.Parallel()

	//	create listener
	ln := fasthttputil.NewInmemoryListener()

	//	init server
	server, err := New().
		WithConfig(&Config{
			ReduceMemoryUsage: true,
			ReadBufferSize:    4096,
			ReadTimeout:       time.Second,
			WriteBufferSize:   4096,
		}).
		WithListener(ln).
		WithDefaultReadinessHandler().
		Build()
	if err != nil {
		t.Fatal(err)
	}

	var (
		done     = make(chan struct{}, 1)
		startErr error
	)

	go func(server *FastHTTPServer, err *error, done chan<- struct{}) {
		defer func(done chan<- struct{}) {
			done <- struct{}{}
		}(done)

		//	start server
		if *err = server.Start(); *err != nil {
			t.Fatal(*err)
		}
	}(server, &startErr, done)

	//	do request
	conn, err := ln.Dial()
	if err != nil {
		t.Fatal(err)
	}

	_, err = conn.Write([]byte("GET /readiness HTTP/1.1\r\n\r\n"))
	if err != nil {
		t.Fatal(err)
	}

	var (
		reader   = bufio.NewReader(conn)
		response fasthttp.Response
	)

	if err = response.Read(reader); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, response.StatusCode())
	assert.Equal(t, "application/json", string(response.Header.Peek("Content-Type")))
	assert.Equal(t, append([]byte{}, "OK"...), response.Body())

	//	close connection
	if err = conn.Close(); err != nil {
		t.Fatal(err)
	}

	//	close server at the end of test
	if err := server.GracefulStop(); err != nil {
		t.Fatal(err)
	}

	<-done

	assert.Nil(t, startErr)
}

func TestCustomReadinessHandler(t *testing.T) {
	t.Parallel()

	//	create listener
	ln := fasthttputil.NewInmemoryListener()

	//	init server
	server, err := New().
		WithConfig(&Config{
			ReduceMemoryUsage: true,
			ReadBufferSize:    4096,
			ReadTimeout:       time.Second,
			WriteBufferSize:   4096,
		}).
		WithListener(ln).
		WithReadinessHandler("/readiness", func(ctx *fasthttp.RequestCtx) {
			ctx.SetStatusCode(200)
			ctx.SetContentType("application/json")
			ctx.SetBodyString("Readiness")
		}).
		Build()
	if err != nil {
		t.Fatal(err)
	}

	var (
		done     = make(chan struct{}, 1)
		startErr error
	)

	go func(server *FastHTTPServer, err *error, done chan<- struct{}) {
		defer func(done chan<- struct{}) {
			done <- struct{}{}
		}(done)

		//	start server
		if *err = server.Start(); *err != nil {
			t.Fatal(*err)
		}
	}(server, &startErr, done)

	//	do request
	conn, err := ln.Dial()
	if err != nil {
		t.Fatal(err)
	}

	_, err = conn.Write([]byte("GET /readiness HTTP/1.1\r\n\r\n"))
	if err != nil {
		t.Fatal(err)
	}

	var (
		reader   = bufio.NewReader(conn)
		response fasthttp.Response
	)

	if err = response.Read(reader); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, response.StatusCode())
	assert.Equal(t, "application/json", string(response.Header.Peek("Content-Type")))
	assert.Equal(t, append([]byte{}, "Readiness"...), response.Body())

	//	close connection
	if err = conn.Close(); err != nil {
		t.Fatal(err)
	}

	//	close server at the end of test
	if err := server.GracefulStop(); err != nil {
		t.Fatal(err)
	}

	<-done

	assert.Nil(t, startErr)
}
