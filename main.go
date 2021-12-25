package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/CristhianRamirez3010/task-manager-go/src/config/constants"
	"github.com/CristhianRamirez3010/task-manager-go/src/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
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
