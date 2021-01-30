# fasthttpserver

Wrapper that implements the HTTP server interface based on the [fasthttp](https://github.com/valyala/fasthttp) library.

# Table of Contents
- [Install](#install)
- [Usage](#usage)
- [Graceful Shutdown](#graceful-shutdown)
- [Liveness/Readiness for k8s](#livenessreadiness-for-k8s)

## Install

```bash
$ go get git.betfavorit.cf/pkg/fasthttpserver
```

## Usage

In the vary simple situation usage of this library looks like this:

```go
package main

import (
    "fmt"
    "os"

    "git.betfavorit.cf/pkg/fasthttpserver"
)

func main() {
    //	init server
    server, err := fasthttpserver.New().
        WithConfig(&fasthttpserver.Config{
            ReduceMemoryUsage: true,
            ReadBufferSize:    4096,
            WriteBufferSize:   4096,
            Addr:              "0.0.0.0:80",
        }).
        Build()
    if err != nil {
        _, _ = fmt.Fprintf(os.Stderr, "error while init server: %v \n", err)
        os.Exit(1)
    }

    //  run server
    if err = server.Start(); err != nil {
        _, _ = fmt.Fprintf(os.Stderr, "error while starting HTTP server: %v \n", err)
    }
}
```

## Graceful Shutdown

You can implement graceful shutdown, like this:

```go
package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"

    "golang.org/x/sync/errgroup"

    "git.betfavorit.cf/pkg/fasthttpserver"
)

func main() {
    //	init server
    server, err := fasthttpserver.New().
        WithConfig(&fasthttpserver.Config{
            ReduceMemoryUsage: true,
            ReadBufferSize:    4096,
            WriteBufferSize:   4096,
            Addr:              "0.0.0.0:80",
        }).
        Build()
    if err != nil {
        _, _ = fmt.Fprintf(os.Stderr, "error while init server: %v \n", err)
        os.Exit(1)
    }

    var (
        errGroup = errgroup.Group{}
        quit     = make(chan os.Signal, 1)
    )
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

    errGroup.Go(func() error {
        _, _ = fmt.Fprintf(os.Stdout, "starting server on %s \n", server.Address())

        return server.Start()
    })

    <-quit

    _, _ = fmt.Fprintln(os.Stderr, "stopping application")

    //	close server at the end of test
    if err := server.GracefulStop(); err != nil {
        _, _ = fmt.Fprintf(os.Stderr, "error while stopping server: %v \n", err)
    }

    if err = errGroup.Wait(); err != nil {
        _, _ = fmt.Fprintf(os.Stderr, "erorr while stopping application: %v \n", err)
    }

    _, _ = fmt.Fprintln(os.Stderr, "application is stopped")
}
```

## Liveness/Readiness for k8s

Cause we run our services in k8s, we should provide liveness/readiness handlers for healthcheck operations. You can use builtin handlers:

```go
package main

import (
    "fmt"
    "os"

    "git.betfavorit.cf/pkg/fasthttpserver"
)

func main() {
    //	init server
    server, err := fasthttpserver.New().
        WithConfig(&fasthttpserver.Config{
            ReduceMemoryUsage: true,
            ReadBufferSize:    4096,
            WriteBufferSize:   4096,
            Addr:              "0.0.0.0:80",
        }).
        WithDefaultLivenessHandler().
        WithDefaultReadinessHandler().
        Build()
    if err != nil {
        _, _ = fmt.Fprintf(os.Stderr, "error while init server: %v \n", err)
        os.Exit(1)
    }

    //  run server
    if err = server.Start(); err != nil {
        _, _ = fmt.Fprintf(os.Stderr, "error while starting HTTP server: %v \n", err)
    }
}
```

Or you can specify your own handlers:

```go
package main

import (
    "fmt"
    "os"

    "github.com/valyala/fasthttp"

    "git.betfavorit.cf/pkg/fasthttpserver"
)

func LivenessHandler(requestCtx *fasthttp.RequestCtx) {
	requestCtx.SetStatusCode(200)
	requestCtx.SetContentType("application/json")
	requestCtx.SetBodyString("Liveness")
}

func ReadinessHandler(requestCtx *fasthttp.RequestCtx) {
	requestCtx.SetStatusCode(200)
	requestCtx.SetContentType("application/json")
	requestCtx.SetBodyString("Readiness")
}

func main() {
    //	init server
    server, err := fasthttpserver.New().
        WithConfig(&fasthttpserver.Config{
            ReduceMemoryUsage: true,
            ReadBufferSize:    4096,
            WriteBufferSize:   4096,
            Addr:              "0.0.0.0:80",
        }).
        WithLivenessHandler("/customliveness", LivenessHandler).
        WithReadinessHandler("/customreadiness", ReadinessHandler).
        Build()
    if err != nil {
        _, _ = fmt.Fprintf(os.Stderr, "error while init server: %v \n", err)
        os.Exit(1)
    }

    //  run server
    if err = server.Start(); err != nil {
        _, _ = fmt.Fprintf(os.Stderr, "error while starting HTTP server: %v \n", err)
    }
}
```