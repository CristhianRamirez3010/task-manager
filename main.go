package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/CristhianRamirez3010/task-manager-go/src/config/constants"
	"github.com/CristhianRamirez3010/task-manager-go/src/config/errorManagerDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/config/responseDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/routes"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/handler"
	"github.com/joho/godotenv"
)

func main() {
	_, err := run()
	if err != nil {
		log.Fatal(err)
	}
	constanst := constants.BuildConstants()
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", constanst.GetPort()),
		Handler: routes.LoadRoutes(),
	}

	err = srv.ListenAndServe()

	log.Fatal(err)
}

type TestModel struct {
	userHandler handler.IUserHandler
}

var (
	buildIUserHandler = handler.BuildIUserHandler
)

func (t *TestModel) RunTest() (*responseDto.ResponseDto, *errorManagerDto.ErrorManagerDto) {
	t.userHandler = buildIUserHandler(nil)
	d := t.userHandler.GetDocuments()
	return d, nil
}

func run() (*constants.Constants, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}
	constanst := constants.BuildConstants()

	return constanst, nil
}
