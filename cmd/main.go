package main

import (
	"fmt"
	"log"

	"github.com/phn00dev/golang-news-app/internal/app"
	"github.com/phn00dev/golang-news-app/internal/setup/constructor"
)

func main() {
	getDependencies, err := app.GetDependencies()
	if err != nil {
		log.Fatal(err)
		return
	}
	constructor.Build(getDependencies)
	runServer := fmt.Sprintf("%s:%s",
		getDependencies.Config.HttpServer.Host,
		getDependencies.Config.HttpServer.Port)

	newApp := app.NewApp(getDependencies)
	if err = newApp.Listen(runServer); err != nil {
		log.Fatal(err)
		return
	}
}
