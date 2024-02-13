package main

import (
	"fmt"
	"github.com/Go-routine-4595/myTimeServer/adapters/controllers"
	"github.com/Go-routine-4595/myTimeServer/domain"
	"github.com/Go-routine-4595/myTimeServer/service"
	"sync"
)

var CompileDate string

const version = 0.1

func main() {
	var (
		wg        *sync.WaitGroup
		svc       domain.IService
		apiServer *controllers.ApiServer
	)

	fmt.Println(CompileDate)
	infoString := fmt.Sprintf("version: %.2f build date/time: %s", version, CompileDate)

	wg = &sync.WaitGroup{}
	svc = service.NewService()

	apiServer = controllers.NewApiServer(svc, "3000", wg, infoString)
	apiServer.Start()

}
