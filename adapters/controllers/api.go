package controllers

import (
	"context"
	"fmt"
	"github.com/Go-routine-4595/myTimeServer/domain"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type ApiServer struct {
	svc        domain.IService
	port       string
	wg         *sync.WaitGroup
	infoString string
}

func NewApiServer(svc domain.IService, p string, wg *sync.WaitGroup, info string) *ApiServer {
	return &ApiServer{
		port:       p,
		svc:        svc,
		wg:         wg,
		infoString: info,
	}
}

func (a *ApiServer) Start() {

	router := gin.Default()

	server := &http.Server{
		Addr:    ":" + a.port,
		Handler: router,
	}

	router.StaticFile("favicon.ico", "./favicon.ico")
	router.GET("/", a.slash)
	router.GET("/gettime", a.getTime)

	// trap SIGINT / SIGTERM to exit cleanly
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Shutting down...")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		a.wg.Wait()
		server.Shutdown(ctx)
		os.Exit(1)
	}()

	err := server.ListenAndServe()
	//err := http.ListenAndServe(":"+a.port, router)
	if err != nil {
		fmt.Println(err)
	}

}

func (a *ApiServer) getTime(c *gin.Context) {
	c.JSONP(http.StatusOK, gin.H{
		"server":       a.infoString,
		"time RFC3999": a.svc.Time(),
	})
}

func (a *ApiServer) slash(c *gin.Context) {
	c.JSONP(http.StatusOK, gin.H{
		"server": a.infoString,
		"status": "running",
	})
}
