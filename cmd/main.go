package main

import (
	"fmt"
	"log"

	"github.com/phn00dev/golang-news-app/internal/app"
)

func main() {
	fmt.Println("Golang news App")
	getDependencies, err := app.GetDependencies()
	if err != nil {
		log.Fatal(err)
		return
	}
	runServer := fmt.Sprintf("%s:%s",
		getDependencies.Config.HttpServer.Host,
		getDependencies.Config.HttpServer.Port)

	newApp := app.NewApp(getDependencies)
	if err = newApp.Listen(runServer); err != nil {
		log.Fatal(err)
		return
	}
}
