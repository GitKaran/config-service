package main

import (
	"github.com/hellofreshdevtests/GitKaran-devops-test/database"
	"github.com/hellofreshdevtests/GitKaran-devops-test/routers"
	"log"
	"os"
)

func main() {

	port, exists := os.LookupEnv("SERVE_PORT")
	if !exists {
		log.Fatal("Env variable SERVE_PORT not defined")
	}

	database.Setup()
	routers.NewApp().Start(port)

}
