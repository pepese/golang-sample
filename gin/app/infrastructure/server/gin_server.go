package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

var (
	gEn     *gin.Engine
	gEnOnce sync.Once
)

type ginServer struct{}

func NewGinServer() *gin.Engine {
	gEnOnce.Do(func() {
		e := gin.New()
		e.Use(requestScope())
		e.Use(accessLog())
		e.Use(gin.Recovery())
		gEn = e
	})
	return gEn
}

func requestScope() gin.HandlerFunc {
	return func(c *gin.Context) {
		// set traceID to context.Context as request scope
		traceID := xid.New().String()
		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, "traceID", traceID)
		c.Request = c.Request.WithContext(ctx)
	}
}

func accessLog() gin.HandlerFunc {
	// Logging Skip Logic
	notlogged := []string{"healthcheck"}
	var skip map[string]struct{}
	if length := len(notlogged); length > 0 {
		skip = make(map[string]struct{}, length)
		for _, path := range notlogged {
			skip[path] = struct{}{}
		}
	}
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		if _, ok := skip[path]; !ok {
			ctx := c.Request.Context()
			fmt.Printf("Request with TraceID-%s.\n", ctx.Value("traceID"))        // Request Log
			defer fmt.Printf("Response with TraceID-%s.\n", ctx.Value("traceID")) // Response Log
			c.Next()
		}
	}
}

func (hs *ginServer) Run() {
	NewGinServer()
	RunWithGracefulStop(gEn)
}

func RunWithGracefulStop(router interface{}) {
	var srv *http.Server
	switch value := router.(type) {
	case *http.ServeMux:
		srv = &http.Server{Addr: ":8080", Handler: value}
	case *gin.Engine:
		srv = &http.Server{Addr: ":8080", Handler: value}
	default:
		log.Fatalln("fatal error")
		return
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
